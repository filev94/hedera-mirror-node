/*-
 * ‌
 * Hedera Mirror Node
 * ​
 * Copyright (C) 2019 - 2022 Hedera Hashgraph, LLC
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

package types

import "github.com/coinbase/rosetta-sdk-go/types"

const (
	OperationTypeCryptoCreateAccount = "CRYPTOCREATEACCOUNT"
	OperationTypeCryptoTransfer      = "CRYPTOTRANSFER"
	OperationTypeTokenAssociate      = "TOKENASSOCIATE"
	OperationTypeTokenBurn           = "TOKENBURN"
	OperationTypeTokenCreate         = "TOKENCREATION"
	OperationTypeTokenDelete         = "TOKENDELETION"
	OperationTypeTokenDissociate     = "TOKENDISSOCIATE" // #nosec
	OperationTypeTokenFreeze         = "TOKENFREEZE"
	OperationTypeTokenGrantKyc       = "TOKENGRANTKYC"
	OperationTypeTokenMint           = "TOKENMINT"
	OperationTypeTokenRevokeKyc      = "TOKENREVOKEKYC"
	OperationTypeTokenUnfreeze       = "TOKENUNFREEZE"
	OperationTypeTokenUpdate         = "TOKENUPDATE"
	OperationTypeTokenWipe           = "TOKENWIPE"
)

const (
	Blockchain = "Hedera"

	currencySymbol   = "HBAR"
	currencyDecimals = 8
)

var TransactionResults = map[int32]string{
	0:   "OK",
	1:   "INVALID_TRANSACTION",
	2:   "PAYER_ACCOUNT_NOT_FOUND",
	3:   "INVALID_NODE_ACCOUNT",
	4:   "TRANSACTION_EXPIRED",
	5:   "INVALID_TRANSACTION_START",
	6:   "INVALID_TRANSACTION_DURATION",
	7:   "INVALID_SIGNATURE",
	8:   "MEMO_TOO_LONG",
	9:   "INSUFFICIENT_TX_FEE",
	10:  "INSUFFICIENT_PAYER_BALANCE",
	11:  "DUPLICATE_TRANSACTION",
	12:  "BUSY",
	13:  "NOT_SUPPORTED",
	14:  "INVALID_FILE_ID",
	15:  "INVALID_ACCOUNT_ID",
	16:  "INVALID_CONTRACT_ID",
	17:  "INVALID_TRANSACTION_ID",
	18:  "RECEIPT_NOT_FOUND",
	19:  "RECORD_NOT_FOUND",
	20:  "INVALID_SOLIDITY_ID",
	21:  "UNKNOWN",
	22:  "SUCCESS",
	23:  "FAIL_INVALID",
	24:  "FAIL_FEE",
	25:  "FAIL_BALANCE",
	26:  "KEY_REQUIRED",
	27:  "BAD_ENCODING",
	28:  "INSUFFICIENT_ACCOUNT_BALANCE",
	29:  "INVALID_SOLIDITY_ADDRESS",
	30:  "INSUFFICIENT_GAS",
	31:  "CONTRACT_SIZE_LIMIT_EXCEEDED",
	32:  "LOCAL_CALL_MODIFICATION_EXCEPTION",
	33:  "CONTRACT_REVERT_EXECUTED",
	34:  "CONTRACT_EXECUTION_EXCEPTION",
	35:  "INVALID_RECEIVING_NODE_ACCOUNT",
	36:  "MISSING_QUERY_HEADER",
	37:  "ACCOUNT_UPDATE_FAILED",
	38:  "INVALID_KEY_ENCODING",
	39:  "NULL_SOLIDITY_ADDRESS",
	40:  "CONTRACT_UPDATE_FAILED",
	41:  "INVALID_QUERY_HEADER",
	42:  "INVALID_FEE_SUBMITTED",
	43:  "INVALID_PAYER_SIGNATURE",
	44:  "KEY_NOT_PROVIDED",
	45:  "INVALID_EXPIRATION_TIME",
	46:  "NO_WACL_KEY",
	47:  "FILE_CONTENT_EMPTY",
	48:  "INVALID_ACCOUNT_AMOUNTS",
	49:  "EMPTY_TRANSACTION_BODY",
	50:  "INVALID_TRANSACTION_BODY",
	51:  "INVALID_SIGNATURE_TYPE_MISMATCHING_KEY",
	52:  "INVALID_SIGNATURE_COUNT_MISMATCHING_KEY",
	53:  "EMPTY_LIVE_HASH_BODY",
	54:  "EMPTY_LIVE_HASH",
	55:  "EMPTY_LIVE_HASH_KEYS",
	56:  "INVALID_LIVE_HASH_SIZE",
	57:  "EMPTY_QUERY_BODY",
	58:  "EMPTY_LIVE_HASH_QUERY",
	59:  "LIVE_HASH_NOT_FOUND",
	60:  "ACCOUNT_ID_DOES_NOT_EXIST",
	61:  "LIVE_HASH_ALREADY_EXISTS",
	62:  "INVALID_FILE_WACL",
	63:  "SERIALIZATION_FAILED",
	64:  "TRANSACTION_OVERSIZE",
	65:  "TRANSACTION_TOO_MANY_LAYERS",
	66:  "CONTRACT_DELETED",
	67:  "PLATFORM_NOT_ACTIVE",
	68:  "KEY_PREFIX_MISMATCH",
	69:  "PLATFORM_TRANSACTION_NOT_CREATED",
	70:  "INVALID_RENEWAL_PERIOD",
	71:  "INVALID_PAYER_ACCOUNT_ID",
	72:  "ACCOUNT_DELETED",
	73:  "FILE_DELETED",
	74:  "ACCOUNT_REPEATED_IN_ACCOUNT_AMOUNTS",
	75:  "SETTING_NEGATIVE_ACCOUNT_BALANCE",
	76:  "OBTAINER_REQUIRED",
	77:  "OBTAINER_SAME_CONTRACT_ID",
	78:  "OBTAINER_DOES_NOT_EXIST",
	79:  "MODIFYING_IMMUTABLE_CONTRACT",
	80:  "FILE_SYSTEM_EXCEPTION",
	81:  "AUTORENEW_DURATION_NOT_IN_RANGE",
	82:  "ERROR_DECODING_BYTESTRING",
	83:  "CONTRACT_FILE_EMPTY",
	84:  "CONTRACT_BYTECODE_EMPTY",
	85:  "INVALID_INITIAL_BALANCE",
	86:  "INVALID_RECEIVE_RECORD_THRESHOLD",
	87:  "INVALID_SEND_RECORD_THRESHOLD",
	88:  "ACCOUNT_IS_NOT_GENESIS_ACCOUNT",
	89:  "PAYER_ACCOUNT_UNAUTHORIZED",
	90:  "INVALID_FREEZE_TRANSACTION_BODY",
	91:  "FREEZE_TRANSACTION_BODY_NOT_FOUND",
	92:  "TRANSFER_LIST_SIZE_LIMIT_EXCEEDED",
	93:  "RESULT_SIZE_LIMIT_EXCEEDED",
	94:  "NOT_SPECIAL_ACCOUNT",
	95:  "CONTRACT_NEGATIVE_GAS",
	96:  "CONTRACT_NEGATIVE_VALUE",
	97:  "INVALID_FEE_FILE",
	98:  "INVALID_EXCHANGE_RATE_FILE",
	99:  "INSUFFICIENT_LOCAL_CALL_GAS",
	100: "ENTITY_NOT_ALLOWED_TO_DELETE",
	101: "AUTHORIZATION_FAILED",
	102: "FILE_UPLOADED_PROTO_INVALID",
	103: "FILE_UPLOADED_PROTO_NOT_SAVED_TO_DISK",
	104: "FEE_SCHEDULE_FILE_PART_UPLOADED",
	105: "EXCHANGE_RATE_CHANGE_LIMIT_EXCEEDED",
	106: "MAX_CONTRACT_STORAGE_EXCEEDED",
	107: "TRANSFER_ACCOUNT_SAME_AS_DELETE_ACCOUNT",
	108: "TOTAL_LEDGER_BALANCE_INVALID",
	110: "EXPIRATION_REDUCTION_NOT_ALLOWED",
	111: "MAX_GAS_LIMIT_EXCEEDED",
	112: "MAX_FILE_SIZE_EXCEEDED",
	113: "RECEIVER_SIG_REQUIRED",
	150: "INVALID_TOPIC_ID",
	155: "INVALID_ADMIN_KEY",
	156: "INVALID_SUBMIT_KEY",
	157: "UNAUTHORIZED",
	158: "INVALID_TOPIC_MESSAGE",
	159: "INVALID_AUTORENEW_ACCOUNT",
	160: "AUTORENEW_ACCOUNT_NOT_ALLOWED",
	162: "TOPIC_EXPIRED",
	163: "INVALID_CHUNK_NUMBER",
	164: "INVALID_CHUNK_TRANSACTION_ID",
	165: "ACCOUNT_FROZEN_FOR_TOKEN",
	166: "TOKENS_PER_ACCOUNT_LIMIT_EXCEEDED",
	167: "INVALID_TOKEN_ID",
	168: "INVALID_TOKEN_DECIMALS",
	169: "INVALID_TOKEN_INITIAL_SUPPLY",
	170: "INVALID_TREASURY_ACCOUNT_FOR_TOKEN",
	171: "INVALID_TOKEN_SYMBOL",
	172: "TOKEN_HAS_NO_FREEZE_KEY",
	173: "TRANSFERS_NOT_ZERO_SUM_FOR_TOKEN",
	174: "MISSING_TOKEN_SYMBOL",
	175: "TOKEN_SYMBOL_TOO_LONG",
	176: "ACCOUNT_KYC_NOT_GRANTED_FOR_TOKEN",
	177: "TOKEN_HAS_NO_KYC_KEY",
	178: "INSUFFICIENT_TOKEN_BALANCE",
	179: "TOKEN_WAS_DELETED",
	180: "TOKEN_HAS_NO_SUPPLY_KEY",
	181: "TOKEN_HAS_NO_WIPE_KEY",
	182: "INVALID_TOKEN_MINT_AMOUNT",
	183: "INVALID_TOKEN_BURN_AMOUNT",
	184: "TOKEN_NOT_ASSOCIATED_TO_ACCOUNT",
	185: "CANNOT_WIPE_TOKEN_TREASURY_ACCOUNT",
	186: "INVALID_KYC_KEY",
	187: "INVALID_WIPE_KEY",
	188: "INVALID_FREEZE_KEY",
	189: "INVALID_SUPPLY_KEY",
	190: "MISSING_TOKEN_NAME",
	191: "TOKEN_NAME_TOO_LONG",
	192: "INVALID_WIPING_AMOUNT",
	193: "TOKEN_IS_IMMUTABLE",
	194: "TOKEN_ALREADY_ASSOCIATED_TO_ACCOUNT",
	195: "TRANSACTION_REQUIRES_ZERO_TOKEN_BALANCES",
	196: "ACCOUNT_IS_TREASURY",
	197: "TOKEN_ID_REPEATED_IN_TOKEN_LIST",
	198: "TOKEN_TRANSFER_LIST_SIZE_LIMIT_EXCEEDED",
	199: "EMPTY_TOKEN_TRANSFER_BODY",
	200: "EMPTY_TOKEN_TRANSFER_ACCOUNT_AMOUNTS",
	201: "INVALID_SCHEDULE_ID",
	202: "SCHEDULE_IS_IMMUTABLE",
	205: "NO_NEW_VALID_SIGNATURES",
	204: "INVALID_SCHEDULE_ACCOUNT_ID",
	203: "INVALID_SCHEDULE_PAYER_ID",
	206: "UNRESOLVABLE_REQUIRED_SIGNERS",
	207: "SCHEDULED_TRANSACTION_NOT_IN_WHITELIST",
	208: "SOME_SIGNATURES_WERE_INVALID",
	209: "TRANSACTION_ID_FIELD_NOT_ALLOWED",
	210: "IDENTICAL_SCHEDULE_ALREADY_CREATED",
	211: "INVALID_ZERO_BYTE_IN_STRING",
	212: "SCHEDULE_ALREADY_DELETED",
	213: "SCHEDULE_ALREADY_EXECUTED",
	214: "MESSAGE_SIZE_TOO_LARGE",
	215: "OPERATION_REPEATED_IN_BUCKET_GROUPS",
	216: "BUCKET_CAPACITY_OVERFLOW",
	217: "NODE_CAPACITY_NOT_SUFFICIENT_FOR_OPERATION",
	218: "BUCKET_HAS_NO_THROTTLE_GROUPS",
	219: "THROTTLE_GROUP_HAS_ZERO_OPS_PER_SEC",
	220: "SUCCESS_BUT_MISSING_EXPECTED_OPERATION",
	221: "UNPARSEABLE_THROTTLE_DEFINITIONS",
	222: "INVALID_THROTTLE_DEFINITIONS",
	223: "ACCOUNT_EXPIRED_AND_PENDING_REMOVAL",
	224: "INVALID_TOKEN_MAX_SUPPLY",
	225: "INVALID_TOKEN_NFT_SERIAL_NUMBER",
	226: "INVALID_NFT_ID",
	227: "METADATA_TOO_LONG",
	228: "BATCH_SIZE_LIMIT_EXCEEDED",
	229: "INVALID_QUERY_RANGE",
	230: "FRACTION_DIVIDES_BY_ZERO",
	231: "INSUFFICIENT_PAYER_BALANCE_FOR_CUSTOM_FEE",
	232: "CUSTOM_FEES_LIST_TOO_LONG",
	233: "INVALID_CUSTOM_FEE_COLLECTOR",
	234: "INVALID_TOKEN_ID_IN_CUSTOM_FEES",
	235: "TOKEN_NOT_ASSOCIATED_TO_FEE_COLLECTOR",
	236: "TOKEN_MAX_SUPPLY_REACHED",
	237: "SENDER_DOES_NOT_OWN_NFT_SERIAL_NO",
	238: "CUSTOM_FEE_NOT_FULLY_SPECIFIED",
	239: "CUSTOM_FEE_MUST_BE_POSITIVE",
	240: "TOKEN_HAS_NO_FEE_SCHEDULE_KEY",
	241: "CUSTOM_FEE_OUTSIDE_NUMERIC_RANGE",
	242: "ROYALTY_FRACTION_CANNOT_EXCEED_ONE",
	243: "FRACTIONAL_FEE_MAX_AMOUNT_LESS_THAN_MIN_AMOUNT",
	244: "CUSTOM_SCHEDULE_ALREADY_HAS_NO_FEES",
	245: "CUSTOM_FEE_DENOMINATION_MUST_BE_FUNGIBLE_COMMON",
	246: "CUSTOM_FRACTIONAL_FEE_ONLY_ALLOWED_FOR_FUNGIBLE_COMMON",
	247: "INVALID_CUSTOM_FEE_SCHEDULE_KEY",
	248: "INVALID_TOKEN_MINT_METADATA",
	249: "INVALID_TOKEN_BURN_METADATA",
	250: "CURRENT_TREASURY_STILL_OWNS_NFTS",
	251: "ACCOUNT_STILL_OWNS_NFTS",
	252: "TREASURY_MUST_OWN_BURNED_NFT",
	253: "ACCOUNT_DOES_NOT_OWN_WIPED_NFT",
	254: "ACCOUNT_AMOUNT_TRANSFERS_ONLY_ALLOWED_FOR_FUNGIBLE_COMMON",
	255: "MAX_NFTS_IN_PRICE_REGIME_HAVE_BEEN_MINTED",
	256: "PAYER_ACCOUNT_DELETED",
	257: "CUSTOM_FEE_CHARGING_EXCEEDED_MAX_RECURSION_DEPTH",
	258: "CUSTOM_FEE_CHARGING_EXCEEDED_MAX_ACCOUNT_AMOUNTS",
	259: "INSUFFICIENT_SENDER_ACCOUNT_BALANCE_FOR_CUSTOM_FEE",
	260: "SERIAL_NUMBER_LIMIT_REACHED",
	261: "CUSTOM_ROYALTY_FEE_ONLY_ALLOWED_FOR_NON_FUNGIBLE_UNIQUE",
	262: "NO_REMAINING_AUTOMATIC_ASSOCIATIONS",
	263: "EXISTING_AUTOMATIC_ASSOCIATIONS_EXCEED_GIVEN_LIMIT",
	264: "REQUESTED_NUM_AUTOMATIC_ASSOCIATIONS_EXCEEDS_ASSOCIATION_LIMIT",
	265: "TOKEN_IS_PAUSED",
	266: "TOKEN_HAS_NO_PAUSE_KEY",
	267: "INVALID_PAUSE_KEY",
	268: "FREEZE_UPDATE_FILE_DOES_NOT_EXIST",
	269: "FREEZE_UPDATE_FILE_HASH_DOES_NOT_MATCH",
	270: "NO_UPGRADE_HAS_BEEN_PREPARED",
	271: "NO_FREEZE_IS_SCHEDULED",
	272: "UPDATE_FILE_HASH_CHANGED_SINCE_PREPARE_UPGRADE",
	273: "FREEZE_START_TIME_MUST_BE_FUTURE",
	274: "PREPARED_UPDATE_FILE_IS_IMMUTABLE",
	275: "FREEZE_ALREADY_SCHEDULED",
	276: "FREEZE_UPGRADE_IN_PROGRESS",
	277: "UPDATE_FILE_ID_DOES_NOT_MATCH_PREPARED",
	278: "UPDATE_FILE_HASH_DOES_NOT_MATCH_PREPARED",
	279: "CONSENSUS_GAS_EXHAUSTED",
	280: "REVERTED_SUCCESS",
	281: "MAX_STORAGE_IN_PRICE_REGIME_HAS_BEEN_USED",
	282: "INVALID_ALIAS_KEY",
	283: "UNEXPECTED_TOKEN_DECIMALS",
	284: "INVALID_PROXY_ACCOUNT_ID",
	285: "INVALID_TRANSFER_ACCOUNT_ID",
	286: "INVALID_FEE_COLLECTOR_ACCOUNT_ID",
	287: "ALIAS_IS_IMMUTABLE",
	288: "SPENDER_ACCOUNT_SAME_AS_OWNER",
	289: "AMOUNT_EXCEEDS_TOKEN_MAX_SUPPLY",
	290: "NEGATIVE_ALLOWANCE_AMOUNT",
	291: "CANNOT_APPROVE_FOR_ALL_FUNGIBLE_COMMON",
	292: "SPENDER_DOES_NOT_HAVE_ALLOWANCE",
	293: "AMOUNT_EXCEEDS_ALLOWANCE",
	294: "MAX_ALLOWANCES_EXCEEDED",
	295: "EMPTY_ALLOWANCES",
	296: "SPENDER_ACCOUNT_REPEATED_IN_ALLOWANCES",
	297: "REPEATED_SERIAL_NUMS_IN_NFT_ALLOWANCES",
	298: "FUNGIBLE_TOKEN_IN_NFT_ALLOWANCES",
	299: "NFT_IN_FUNGIBLE_TOKEN_ALLOWANCES",
	300: "INVALID_ALLOWANCE_OWNER_ID",
	301: "INVALID_ALLOWANCE_SPENDER_ID",
	302: "REPEATED_ALLOWANCES_TO_DELETE",
}

var TransactionTypes = map[int32]string{
	7:  "CONTRACTCALL",
	8:  "CONTRACTCREATEINSTANCE",
	9:  "CONTRACTUPDATEINSTANCE",
	10: "CRYPTOADDLIVEHASH",
	11: OperationTypeCryptoCreateAccount,
	12: "CRYPTODELETE",
	13: "CRYPTODELETELIVEHASH",
	14: OperationTypeCryptoTransfer,
	15: "CRYPTOUPDATEACCOUNT",
	16: "FILEAPPEND",
	17: "FILECREATE",
	18: "FILEDELETE",
	19: "FILEUPDATE",
	20: "SYSTEMDELETE",
	21: "SYSTEMUNDELETE",
	22: "CONTRACTDELETEINSTANCE",
	23: "FREEZE",
	24: "CONSENSUSCREATETOPIC",
	25: "CONSENSUSUPDATETOPIC",
	26: "CONSENSUSDELETETOPIC",
	27: "CONSENSUSSUBMITMESSAGE",
	28: "UNCHECKEDSUBMIT",
	29: OperationTypeTokenCreate,
	31: OperationTypeTokenFreeze,
	32: OperationTypeTokenUnfreeze,
	33: OperationTypeTokenGrantKyc,
	34: OperationTypeTokenRevokeKyc,
	35: OperationTypeTokenDelete,
	36: OperationTypeTokenUpdate,
	37: OperationTypeTokenMint,
	38: OperationTypeTokenBurn,
	39: OperationTypeTokenWipe,
	40: OperationTypeTokenAssociate,
	41: OperationTypeTokenDissociate,
	42: "SCHEDULECREATE",
	43: "SCHEDULEDELETE",
	44: "SCHEDULESIGN",
	45: "TOKENFEESCHEDULEUPDATE",
	46: "TOKENPAUSE",
	47: "TOKENUNPAUSE",
	48: "CRYPTOADJUSTALLOWANCE",
	49: "CRYPTOAPPROVEALLOWANCE",
}

var (
	CurrencyHbar = &types.Currency{
		Symbol:   currencySymbol,
		Decimals: currencyDecimals,
		Metadata: map[string]interface{}{
			"issuer": Blockchain,
		},
	}

	SupportedOperationTypes = []string{
		OperationTypeCryptoCreateAccount,
		OperationTypeCryptoTransfer,
		OperationTypeTokenAssociate,
		OperationTypeTokenBurn,
		OperationTypeTokenCreate,
		OperationTypeTokenDelete,
		OperationTypeTokenDissociate,
		OperationTypeTokenFreeze,
		OperationTypeTokenGrantKyc,
		OperationTypeTokenMint,
		OperationTypeTokenRevokeKyc,
		OperationTypeTokenUnfreeze,
		OperationTypeTokenUpdate,
		OperationTypeTokenWipe,
	}
)
