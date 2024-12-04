/*
 * Copyright (C) 2020-2024 Hedera Hashgraph, LLC
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

package com.hedera.mirror.test.e2e.acceptance.steps;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import com.hedera.hashgraph.sdk.FileId;
import com.hedera.hashgraph.sdk.FileInfo;
import com.hedera.mirror.rest.model.TransactionByIdResponse;
import com.hedera.mirror.rest.model.TransactionDetail;
import com.hedera.mirror.test.e2e.acceptance.client.FileClient;
import com.hedera.mirror.test.e2e.acceptance.client.MirrorNodeClient;
import com.hedera.mirror.test.e2e.acceptance.response.NetworkTransactionResponse;
import io.cucumber.java.en.Given;
import io.cucumber.java.en.Then;
import java.nio.charset.StandardCharsets;
import java.util.List;
import lombok.CustomLog;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;

@CustomLog
@RequiredArgsConstructor
public class FileFeature {
    private static final String ORIGINAL_FILE_CONTENTS = "Mirror Node v1";
    private static final String UPDATE_BASE_FILE_CONTENTS = "Mirror Node v2,";
    private static final String APPEND_FILE_CONTENTS = " new and improved";
    private static final String UPDATED_FILE_CONTENTS = UPDATE_BASE_FILE_CONTENTS + APPEND_FILE_CONTENTS;

    private final FileClient fileClient;
    private final MirrorNodeClient mirrorClient;

    private NetworkTransactionResponse networkTransactionResponse;
    private FileId fileId;
    private FileInfo fileInfo;

    @Given("I successfully create a file")
    public void createNewFile() {
        networkTransactionResponse = fileClient.createFile(ORIGINAL_FILE_CONTENTS.getBytes(StandardCharsets.UTF_8));

        assertNotNull(networkTransactionResponse.getTransactionId());
        assertNotNull(networkTransactionResponse.getReceipt());

        fileId = networkTransactionResponse.getReceipt().fileId;
        assertNotNull(fileId);
    }

    @Given("I successfully update the file")
    public void updateFile() {
        networkTransactionResponse =
                fileClient.updateFile(fileId, UPDATE_BASE_FILE_CONTENTS.getBytes(StandardCharsets.UTF_8));

        assertNotNull(networkTransactionResponse.getTransactionId());
        assertNotNull(networkTransactionResponse.getReceipt());
    }

    @Given("I successfully append to the file")
    public void appendFile() {
        networkTransactionResponse =
                fileClient.appendFile(fileId, APPEND_FILE_CONTENTS.getBytes(StandardCharsets.UTF_8));

        assertNotNull(networkTransactionResponse.getTransactionId());
        assertNotNull(networkTransactionResponse.getReceipt());
    }

    @Given("I successfully delete the file")
    public void deleteFile() {
        networkTransactionResponse = fileClient.deleteFile(fileId);

        assertNotNull(networkTransactionResponse.getTransactionId());
        assertNotNull(networkTransactionResponse.getReceipt());
    }

    @Then("the mirror node REST API should return status {int} for the file transaction")
    public void verifyMirrorAPIResponses(int status) {
        String transactionId = networkTransactionResponse.getTransactionIdStringNoCheckSum();
        TransactionByIdResponse mirrorTransactionsResponse = mirrorClient.getTransactions(transactionId);

        TransactionDetail mirrorTransaction = verifyMirrorTransactionsResponse(mirrorTransactionsResponse, status);

        assertThat(mirrorTransaction.getValidStartTimestamp())
                .isEqualTo(networkTransactionResponse.getValidStartString());
        assertThat(mirrorTransaction.getTransactionId()).isEqualTo(transactionId);
    }

    private TransactionDetail verifyMirrorTransactionsResponse(
            TransactionByIdResponse mirrorTransactionsResponse, int status) {
        List<TransactionDetail> transactions = mirrorTransactionsResponse.getTransactions();
        assertNotNull(transactions);
        assertThat(transactions).isNotEmpty();
        TransactionDetail mirrorTransaction = transactions.get(0);

        if (status == HttpStatus.OK.value()) {
            assertThat(mirrorTransaction.getResult()).isEqualTo("SUCCESS");
        }

        assertThat(mirrorTransaction.getValidStartTimestamp()).isNotNull();
        assertThat(mirrorTransaction.getName()).isNotNull();
        assertThat(mirrorTransaction.getResult()).isNotNull();
        assertThat(mirrorTransaction.getConsensusTimestamp()).isNotNull();
        assertThat(mirrorTransaction.getEntityId()).isEqualTo(fileId.toString());

        return mirrorTransaction;
    }
}
