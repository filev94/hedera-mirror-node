/*
 * Copyright (C) 2021-2025 Hedera Hashgraph, LLC
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

package com.hedera.services.fees.usage.token.meta;

import com.google.common.base.MoreObjects;
import com.hederahashgraph.api.proto.java.SubType;

/**
 *  Exact copy from hedera-services
 */
public abstract class TokenOpMetaBase {
    private final int bpt;
    private final SubType subType;
    private final long transferRecordRb;

    protected TokenOpMetaBase(final int bpt, final SubType subType, final long transferRecordRb) {
        this.bpt = bpt;
        this.subType = subType;
        this.transferRecordRb = transferRecordRb;
    }

    public SubType getSubType() {
        return subType;
    }

    public int getBpt() {
        return bpt;
    }

    public long getTransferRecordDb() {
        return transferRecordRb;
    }

    @Override
    public String toString() {
        return toStringHelper().toString();
    }

    protected MoreObjects.ToStringHelper toStringHelper() {
        return MoreObjects.toStringHelper(this)
                .add("bpt", bpt)
                .add("transferRecordDb", transferRecordRb)
                .add("subType", subType);
    }
}
