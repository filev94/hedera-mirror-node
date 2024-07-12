/*
 * Copyright (C) 2019-2024 Hedera Hashgraph, LLC
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

package com.hedera.mirror.common.domain.transaction;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.hedera.mirror.common.domain.entity.EntityId;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;
import org.springframework.data.domain.Persistable;

@AllArgsConstructor(access = AccessLevel.PRIVATE) // For builder
@Builder
@Data
@Entity
@NoArgsConstructor
public class EthereumTransaction implements Persistable<Long> {

    @ToString.Exclude
    private byte[] accessList;

    @ToString.Exclude
    private byte[] callData;

    private EntityId callDataId;

    @ToString.Exclude
    private byte[] chainId;

    @Id
    private long consensusTimestamp;

    @ToString.Exclude
    private byte[] data;

    // persisted in tinybar
    private Long gasLimit;

    // persisted in tinybar
    private byte[] gasPrice;

    @ToString.Exclude
    private byte[] hash;

    // persisted in tinybar
    private byte[] maxFeePerGas;

    // persisted in tinybar
    private Long maxGasAllowance;

    // persisted in tinybar
    private byte[] maxPriorityFeePerGas;

    private Long nonce;

    private EntityId payerAccountId;

    private Integer recoveryId;

    @Column(name = "signature_r")
    @ToString.Exclude
    private byte[] signatureR;

    @Column(name = "signature_s")
    @ToString.Exclude
    private byte[] signatureS;

    @Column(name = "signature_v")
    @ToString.Exclude
    private byte[] signatureV;

    @ToString.Exclude
    private byte[] toAddress;

    private Integer type;

    @ToString.Exclude
    private byte[] value;

    @JsonIgnore
    @Override
    public Long getId() {
        return consensusTimestamp;
    }

    @JsonIgnore
    @Override
    public boolean isNew() {
        return true; // Since we never update and use a natural ID, avoid Hibernate querying before insert
    }

    public TransactionHash toTransactionHash() {
        return TransactionHash.builder()
                .consensusTimestamp(consensusTimestamp)
                .hash(hash)
                .payerAccountId(payerAccountId.getId())
                .build();
    }
}
