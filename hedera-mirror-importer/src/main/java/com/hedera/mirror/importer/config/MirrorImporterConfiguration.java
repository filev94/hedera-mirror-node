package com.hedera.mirror.importer.config;

/*-
 * ‌
 * Hedera Mirror Node
 * ​
 * Copyright (C) 2019 Hedera Hashgraph, LLC
 * ​
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ‍
 */

import java.net.URI;
import java.time.Duration;
import lombok.RequiredArgsConstructor;
import lombok.extern.log4j.Log4j2;
import org.apache.commons.lang3.StringUtils;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.scheduling.annotation.EnableScheduling;
import software.amazon.awssdk.auth.credentials.AnonymousCredentialsProvider;
import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.AwsCredentialsProvider;
import software.amazon.awssdk.auth.credentials.StaticCredentialsProvider;
import software.amazon.awssdk.http.async.SdkAsyncHttpClient;
import software.amazon.awssdk.http.nio.netty.NettyNioAsyncHttpClient;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3AsyncClient;
import software.amazon.awssdk.services.s3.S3AsyncClientBuilder;

import com.hedera.mirror.importer.downloader.CommonDownloaderProperties;

@Configuration
@EnableAsync
@Log4j2
@RequiredArgsConstructor
public class MirrorImporterConfiguration {

    private final CommonDownloaderProperties downloaderProperties;

    @Bean
    public S3AsyncClient s3AsyncClient() {
        SdkAsyncHttpClient httpClient = NettyNioAsyncHttpClient.builder()
                .maxConcurrency(downloaderProperties.getMaxConcurrency())
                .connectionMaxIdleTime(Duration.ofSeconds(5))  // https://github.com/aws/aws-sdk-java-v2/issues/1122
                .build();

        log.info("Configured to download from {} in region {} with bucket name '{}'", downloaderProperties
                .getCloudProvider(), downloaderProperties.getRegion(), downloaderProperties.getBucketName());

        return S3AsyncClient.builder()
                .credentialsProvider(awsCredentialsProvider(downloaderProperties))
                .region(Region.of(downloaderProperties.getRegion()))
                .httpClient(httpClient)
                .applyMutation(this::overrideEndpoint)
                .build();
    }

    private void overrideEndpoint(S3AsyncClientBuilder builder) {
        // If Amazon S3, let the client generate the endpoint automatically based upon the region and bucket name
        if (downloaderProperties.getCloudProvider() != CommonDownloaderProperties.CloudProvider.S3) {
            URI endpoint = URI.create(downloaderProperties.getCloudProvider().getEndpoint());
            builder.endpointOverride(endpoint);
        }
    }

    @Bean
    AwsCredentialsProvider awsCredentialsProvider(CommonDownloaderProperties downloaderProperties) {
        AwsCredentialsProvider awsCredentials;
        String accessKey = downloaderProperties.getAccessKey();
        String secretKey = downloaderProperties.getSecretKey();

        if (StringUtils.isNotBlank(accessKey) && StringUtils.isNotBlank(secretKey)) {
            log.info("Setting up S3 async client using provided access/secret key");
            awsCredentials = StaticCredentialsProvider.create(AwsBasicCredentials.create(accessKey, secretKey));
        } else {
            log.info("Setting up S3 async client using anonymous credentials");
            awsCredentials = AnonymousCredentialsProvider.create();
        }

        return awsCredentials;
    }

    @Configuration
    @ConditionalOnProperty(prefix = "spring.task.scheduling", name = "enabled", havingValue = "true", matchIfMissing
            = true)
    @EnableScheduling
    protected static class SchedulingConfiguration {
    }
}
