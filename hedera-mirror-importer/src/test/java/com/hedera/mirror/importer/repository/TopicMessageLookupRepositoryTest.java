/*
 * Copyright (C) 2019-2025 Hedera Hashgraph, LLC
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

package com.hedera.mirror.importer.repository;

import static org.assertj.core.api.Assertions.assertThat;

import com.hedera.mirror.importer.ImporterIntegrationTest;
import lombok.RequiredArgsConstructor;
import org.junit.jupiter.api.Test;

@RequiredArgsConstructor
class TopicMessageLookupRepositoryTest extends ImporterIntegrationTest {

    private final TopicMessageLookupRepository repository;

    @Test
    void prune() {
        domainBuilder.topicMessageLookup().persist();
        var topicMessageLookup2 = domainBuilder.topicMessageLookup().persist();
        var topicMessageLookup3 = domainBuilder.topicMessageLookup().persist();
        repository.prune(topicMessageLookup2.getTimestampRange().upperEndpoint());
        assertThat(repository.findAll()).containsOnly(topicMessageLookup3);
    }

    @Test
    void save() {
        var topicMessageLookup = domainBuilder.topicMessageLookup().get();
        repository.save(topicMessageLookup);
        assertThat(repository.findById(topicMessageLookup.getId())).get().isEqualTo(topicMessageLookup);
    }
}
