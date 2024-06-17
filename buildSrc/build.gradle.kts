/*
 * Copyright (C) 2022-2024 Hedera Hashgraph, LLC
 *
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
 */

plugins { `kotlin-dsl` }

repositories {
    gradlePluginPortal()
    mavenCentral()
}

dependencies {
    val dockerJavaVersion = "3.3.6"
    val flywayVersion = "10.15.0"
    val jooqVersion = "3.19.10"

    // Add docker-java dependencies before gradle-docker-plugin to avoid the docker-java jars
    // embedded in the plugin being used by testcontainers-postgresql
    implementation("com.github.docker-java:docker-java-api:$dockerJavaVersion")
    implementation("com.github.docker-java:docker-java-core:$dockerJavaVersion")

    implementation("com.bmuschko:gradle-docker-plugin:9.4.0")
    implementation("com.diffplug.spotless:spotless-plugin-gradle:6.25.0")
    implementation("com.github.johnrengelman:shadow:8.1.1")
    implementation("com.github.node-gradle:gradle-node-plugin:7.0.2")
    implementation("com.google.protobuf:protobuf-gradle-plugin:0.9.4")
    implementation("com.gorylenko.gradle-git-properties:gradle-git-properties:2.4.2")
    implementation("com.graphql-java-generator:graphql-gradle-plugin3:2.6")
    implementation("gradle.plugin.io.snyk.gradle.plugin:snyk:0.6.1")
    implementation("gradle.plugin.org.flywaydb:gradle-plugin-publishing:$flywayVersion")
    implementation("io.freefair.gradle:lombok-plugin:8.6")
    implementation("io.spring.gradle:dependency-management-plugin:1.1.5")
    implementation("org.apache.commons:commons-compress:1.26.1")
    implementation("org.flywaydb:flyway-database-postgresql:$flywayVersion")
    implementation("org.gradle:test-retry-gradle-plugin:1.5.9")
    implementation("org.jooq:jooq-codegen-gradle:$jooqVersion")
    implementation("org.jooq:jooq-meta:$jooqVersion")
    implementation("org.openapitools:openapi-generator-gradle-plugin:7.6.0")
    implementation("org.owasp:dependency-check-gradle:9.2.0")
    implementation("org.sonarsource.scanner.gradle:sonarqube-gradle-plugin:5.0.0.4638")
    implementation("org.springframework.boot:spring-boot-gradle-plugin:3.3.0")
    implementation("org.testcontainers:postgresql:1.19.8")
}

val gitHook =
    tasks.register<Exec>("gitHook") {
        commandLine("git", "config", "core.hookspath", "buildSrc/src/main/resources/hooks")
    }

project.tasks.build { dependsOn(gitHook) }
