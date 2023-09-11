/*-
 * ‌
 * Hedera Mirror Node
 * ​
 * Copyright (C) 2019 - 2023 Hedera Hashgraph, LLC
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

plugins {
    `kotlin-dsl`
}

repositories {
    gradlePluginPortal()
    mavenCentral()
}

dependencies {
    implementation("com.bmuschko:gradle-docker-plugin:9.3.2")
    implementation("com.diffplug.spotless:spotless-plugin-gradle:6.21.0")
    implementation("com.github.johnrengelman:shadow:8.1.1")
    implementation("com.github.node-gradle:gradle-node-plugin:7.0.0")
    implementation("com.google.protobuf:protobuf-gradle-plugin:0.9.4")
    implementation("com.gorylenko.gradle-git-properties:gradle-git-properties:2.4.1")
    implementation("com.graphql-java-generator:graphql-gradle-plugin3:2.2")
    implementation("gradle.plugin.io.snyk.gradle.plugin:snyk:0.5")
    implementation("io.freefair.gradle:lombok-plugin:8.3")
    implementation("io.spring.gradle:dependency-management-plugin:1.1.3")
    implementation("org.apache.commons:commons-compress:1.24.0")
    implementation("org.gradle:test-retry-gradle-plugin:1.5.4")
    implementation("org.openapitools:openapi-generator-gradle-plugin:7.0.0")
    implementation("org.owasp:dependency-check-gradle:8.4.0")
    implementation("org.sonarsource.scanner.gradle:sonarqube-gradle-plugin:4.3.1.3277")
    implementation("org.springframework.boot:spring-boot-gradle-plugin:3.1.2")
}

val gitHook = tasks.register<Exec>("gitHook") {
    commandLine("git", "config", "core.hookspath", "buildSrc/src/main/resources/hooks")
}

project.tasks.build {
    dependsOn(gitHook)
}
