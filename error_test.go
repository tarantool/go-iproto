// Code generated by generate.sh; DO NOT EDIT.

package iproto_test

import (
	"testing"

	"github.com/tarantool/go-iproto"
)

func TestError(t *testing.T) {
	cases := []struct {
		Err iproto.Error
		Str string
	}{
		{iproto.ER_UNKNOWN, "ER_UNKNOWN"},
		{iproto.ER_ILLEGAL_PARAMS, "ER_ILLEGAL_PARAMS"},
		{iproto.ER_MEMORY_ISSUE, "ER_MEMORY_ISSUE"},
		{iproto.ER_TUPLE_FOUND, "ER_TUPLE_FOUND"},
		{iproto.ER_TUPLE_NOT_FOUND, "ER_TUPLE_NOT_FOUND"},
		{iproto.ER_UNSUPPORTED, "ER_UNSUPPORTED"},
		{iproto.ER_NONMASTER, "ER_NONMASTER"},
		{iproto.ER_READONLY, "ER_READONLY"},
		{iproto.ER_INJECTION, "ER_INJECTION"},
		{iproto.ER_CREATE_SPACE, "ER_CREATE_SPACE"},
		{iproto.ER_SPACE_EXISTS, "ER_SPACE_EXISTS"},
		{iproto.ER_DROP_SPACE, "ER_DROP_SPACE"},
		{iproto.ER_ALTER_SPACE, "ER_ALTER_SPACE"},
		{iproto.ER_INDEX_TYPE, "ER_INDEX_TYPE"},
		{iproto.ER_MODIFY_INDEX, "ER_MODIFY_INDEX"},
		{iproto.ER_LAST_DROP, "ER_LAST_DROP"},
		{iproto.ER_TUPLE_FORMAT_LIMIT, "ER_TUPLE_FORMAT_LIMIT"},
		{iproto.ER_DROP_PRIMARY_KEY, "ER_DROP_PRIMARY_KEY"},
		{iproto.ER_KEY_PART_TYPE, "ER_KEY_PART_TYPE"},
		{iproto.ER_EXACT_MATCH, "ER_EXACT_MATCH"},
		{iproto.ER_INVALID_MSGPACK, "ER_INVALID_MSGPACK"},
		{iproto.ER_PROC_RET, "ER_PROC_RET"},
		{iproto.ER_TUPLE_NOT_ARRAY, "ER_TUPLE_NOT_ARRAY"},
		{iproto.ER_FIELD_TYPE, "ER_FIELD_TYPE"},
		{iproto.ER_INDEX_PART_TYPE_MISMATCH, "ER_INDEX_PART_TYPE_MISMATCH"},
		{iproto.ER_UPDATE_SPLICE, "ER_UPDATE_SPLICE"},
		{iproto.ER_UPDATE_ARG_TYPE, "ER_UPDATE_ARG_TYPE"},
		{iproto.ER_FORMAT_MISMATCH_INDEX_PART, "ER_FORMAT_MISMATCH_INDEX_PART"},
		{iproto.ER_UNKNOWN_UPDATE_OP, "ER_UNKNOWN_UPDATE_OP"},
		{iproto.ER_UPDATE_FIELD, "ER_UPDATE_FIELD"},
		{iproto.ER_FUNCTION_TX_ACTIVE, "ER_FUNCTION_TX_ACTIVE"},
		{iproto.ER_KEY_PART_COUNT, "ER_KEY_PART_COUNT"},
		{iproto.ER_PROC_LUA, "ER_PROC_LUA"},
		{iproto.ER_NO_SUCH_PROC, "ER_NO_SUCH_PROC"},
		{iproto.ER_NO_SUCH_TRIGGER, "ER_NO_SUCH_TRIGGER"},
		{iproto.ER_NO_SUCH_INDEX_ID, "ER_NO_SUCH_INDEX_ID"},
		{iproto.ER_NO_SUCH_SPACE, "ER_NO_SUCH_SPACE"},
		{iproto.ER_NO_SUCH_FIELD_NO, "ER_NO_SUCH_FIELD_NO"},
		{iproto.ER_EXACT_FIELD_COUNT, "ER_EXACT_FIELD_COUNT"},
		{iproto.ER_FIELD_MISSING, "ER_FIELD_MISSING"},
		{iproto.ER_WAL_IO, "ER_WAL_IO"},
		{iproto.ER_MORE_THAN_ONE_TUPLE, "ER_MORE_THAN_ONE_TUPLE"},
		{iproto.ER_ACCESS_DENIED, "ER_ACCESS_DENIED"},
		{iproto.ER_CREATE_USER, "ER_CREATE_USER"},
		{iproto.ER_DROP_USER, "ER_DROP_USER"},
		{iproto.ER_NO_SUCH_USER, "ER_NO_SUCH_USER"},
		{iproto.ER_USER_EXISTS, "ER_USER_EXISTS"},
		{iproto.ER_CREDS_MISMATCH, "ER_CREDS_MISMATCH"},
		{iproto.ER_UNKNOWN_REQUEST_TYPE, "ER_UNKNOWN_REQUEST_TYPE"},
		{iproto.ER_UNKNOWN_SCHEMA_OBJECT, "ER_UNKNOWN_SCHEMA_OBJECT"},
		{iproto.ER_CREATE_FUNCTION, "ER_CREATE_FUNCTION"},
		{iproto.ER_NO_SUCH_FUNCTION, "ER_NO_SUCH_FUNCTION"},
		{iproto.ER_FUNCTION_EXISTS, "ER_FUNCTION_EXISTS"},
		{iproto.ER_BEFORE_REPLACE_RET, "ER_BEFORE_REPLACE_RET"},
		{iproto.ER_MULTISTATEMENT_TRANSACTION, "ER_MULTISTATEMENT_TRANSACTION"},
		{iproto.ER_TRIGGER_EXISTS, "ER_TRIGGER_EXISTS"},
		{iproto.ER_USER_MAX, "ER_USER_MAX"},
		{iproto.ER_NO_SUCH_ENGINE, "ER_NO_SUCH_ENGINE"},
		{iproto.ER_RELOAD_CFG, "ER_RELOAD_CFG"},
		{iproto.ER_CFG, "ER_CFG"},
		{iproto.ER_SAVEPOINT_EMPTY_TX, "ER_SAVEPOINT_EMPTY_TX"},
		{iproto.ER_NO_SUCH_SAVEPOINT, "ER_NO_SUCH_SAVEPOINT"},
		{iproto.ER_UNKNOWN_REPLICA, "ER_UNKNOWN_REPLICA"},
		{iproto.ER_REPLICASET_UUID_MISMATCH, "ER_REPLICASET_UUID_MISMATCH"},
		{iproto.ER_INVALID_UUID, "ER_INVALID_UUID"},
		{iproto.ER_REPLICASET_UUID_IS_RO, "ER_REPLICASET_UUID_IS_RO"},
		{iproto.ER_INSTANCE_UUID_MISMATCH, "ER_INSTANCE_UUID_MISMATCH"},
		{iproto.ER_REPLICA_ID_IS_RESERVED, "ER_REPLICA_ID_IS_RESERVED"},
		{iproto.ER_INVALID_ORDER, "ER_INVALID_ORDER"},
		{iproto.ER_MISSING_REQUEST_FIELD, "ER_MISSING_REQUEST_FIELD"},
		{iproto.ER_IDENTIFIER, "ER_IDENTIFIER"},
		{iproto.ER_DROP_FUNCTION, "ER_DROP_FUNCTION"},
		{iproto.ER_ITERATOR_TYPE, "ER_ITERATOR_TYPE"},
		{iproto.ER_REPLICA_MAX, "ER_REPLICA_MAX"},
		{iproto.ER_INVALID_XLOG, "ER_INVALID_XLOG"},
		{iproto.ER_INVALID_XLOG_NAME, "ER_INVALID_XLOG_NAME"},
		{iproto.ER_INVALID_XLOG_ORDER, "ER_INVALID_XLOG_ORDER"},
		{iproto.ER_NO_CONNECTION, "ER_NO_CONNECTION"},
		{iproto.ER_TIMEOUT, "ER_TIMEOUT"},
		{iproto.ER_ACTIVE_TRANSACTION, "ER_ACTIVE_TRANSACTION"},
		{iproto.ER_CURSOR_NO_TRANSACTION, "ER_CURSOR_NO_TRANSACTION"},
		{iproto.ER_CROSS_ENGINE_TRANSACTION, "ER_CROSS_ENGINE_TRANSACTION"},
		{iproto.ER_NO_SUCH_ROLE, "ER_NO_SUCH_ROLE"},
		{iproto.ER_ROLE_EXISTS, "ER_ROLE_EXISTS"},
		{iproto.ER_CREATE_ROLE, "ER_CREATE_ROLE"},
		{iproto.ER_INDEX_EXISTS, "ER_INDEX_EXISTS"},
		{iproto.ER_SESSION_CLOSED, "ER_SESSION_CLOSED"},
		{iproto.ER_ROLE_LOOP, "ER_ROLE_LOOP"},
		{iproto.ER_GRANT, "ER_GRANT"},
		{iproto.ER_PRIV_GRANTED, "ER_PRIV_GRANTED"},
		{iproto.ER_ROLE_GRANTED, "ER_ROLE_GRANTED"},
		{iproto.ER_PRIV_NOT_GRANTED, "ER_PRIV_NOT_GRANTED"},
		{iproto.ER_ROLE_NOT_GRANTED, "ER_ROLE_NOT_GRANTED"},
		{iproto.ER_MISSING_SNAPSHOT, "ER_MISSING_SNAPSHOT"},
		{iproto.ER_CANT_UPDATE_PRIMARY_KEY, "ER_CANT_UPDATE_PRIMARY_KEY"},
		{iproto.ER_UPDATE_INTEGER_OVERFLOW, "ER_UPDATE_INTEGER_OVERFLOW"},
		{iproto.ER_GUEST_USER_PASSWORD, "ER_GUEST_USER_PASSWORD"},
		{iproto.ER_TRANSACTION_CONFLICT, "ER_TRANSACTION_CONFLICT"},
		{iproto.ER_UNSUPPORTED_PRIV, "ER_UNSUPPORTED_PRIV"},
		{iproto.ER_LOAD_FUNCTION, "ER_LOAD_FUNCTION"},
		{iproto.ER_FUNCTION_LANGUAGE, "ER_FUNCTION_LANGUAGE"},
		{iproto.ER_RTREE_RECT, "ER_RTREE_RECT"},
		{iproto.ER_PROC_C, "ER_PROC_C"},
		{iproto.ER_UNKNOWN_RTREE_INDEX_DISTANCE_TYPE, "ER_UNKNOWN_RTREE_INDEX_DISTANCE_TYPE"},
		{iproto.ER_PROTOCOL, "ER_PROTOCOL"},
		{iproto.ER_UPSERT_UNIQUE_SECONDARY_KEY, "ER_UPSERT_UNIQUE_SECONDARY_KEY"},
		{iproto.ER_WRONG_INDEX_RECORD, "ER_WRONG_INDEX_RECORD"},
		{iproto.ER_WRONG_INDEX_PARTS, "ER_WRONG_INDEX_PARTS"},
		{iproto.ER_WRONG_INDEX_OPTIONS, "ER_WRONG_INDEX_OPTIONS"},
		{iproto.ER_WRONG_SCHEMA_VERSION, "ER_WRONG_SCHEMA_VERSION"},
		{iproto.ER_MEMTX_MAX_TUPLE_SIZE, "ER_MEMTX_MAX_TUPLE_SIZE"},
		{iproto.ER_WRONG_SPACE_OPTIONS, "ER_WRONG_SPACE_OPTIONS"},
		{iproto.ER_UNSUPPORTED_INDEX_FEATURE, "ER_UNSUPPORTED_INDEX_FEATURE"},
		{iproto.ER_VIEW_IS_RO, "ER_VIEW_IS_RO"},
		{iproto.ER_NO_TRANSACTION, "ER_NO_TRANSACTION"},
		{iproto.ER_SYSTEM, "ER_SYSTEM"},
		{iproto.ER_LOADING, "ER_LOADING"},
		{iproto.ER_CONNECTION_TO_SELF, "ER_CONNECTION_TO_SELF"},
		{iproto.ER_KEY_PART_IS_TOO_LONG, "ER_KEY_PART_IS_TOO_LONG"},
		{iproto.ER_COMPRESSION, "ER_COMPRESSION"},
		{iproto.ER_CHECKPOINT_IN_PROGRESS, "ER_CHECKPOINT_IN_PROGRESS"},
		{iproto.ER_SUB_STMT_MAX, "ER_SUB_STMT_MAX"},
		{iproto.ER_COMMIT_IN_SUB_STMT, "ER_COMMIT_IN_SUB_STMT"},
		{iproto.ER_ROLLBACK_IN_SUB_STMT, "ER_ROLLBACK_IN_SUB_STMT"},
		{iproto.ER_DECOMPRESSION, "ER_DECOMPRESSION"},
		{iproto.ER_INVALID_XLOG_TYPE, "ER_INVALID_XLOG_TYPE"},
		{iproto.ER_ALREADY_RUNNING, "ER_ALREADY_RUNNING"},
		{iproto.ER_INDEX_FIELD_COUNT_LIMIT, "ER_INDEX_FIELD_COUNT_LIMIT"},
		{iproto.ER_LOCAL_INSTANCE_ID_IS_READ_ONLY, "ER_LOCAL_INSTANCE_ID_IS_READ_ONLY"},
		{iproto.ER_BACKUP_IN_PROGRESS, "ER_BACKUP_IN_PROGRESS"},
		{iproto.ER_READ_VIEW_ABORTED, "ER_READ_VIEW_ABORTED"},
		{iproto.ER_INVALID_INDEX_FILE, "ER_INVALID_INDEX_FILE"},
		{iproto.ER_INVALID_RUN_FILE, "ER_INVALID_RUN_FILE"},
		{iproto.ER_INVALID_VYLOG_FILE, "ER_INVALID_VYLOG_FILE"},
		{iproto.ER_CASCADE_ROLLBACK, "ER_CASCADE_ROLLBACK"},
		{iproto.ER_VY_QUOTA_TIMEOUT, "ER_VY_QUOTA_TIMEOUT"},
		{iproto.ER_PARTIAL_KEY, "ER_PARTIAL_KEY"},
		{iproto.ER_TRUNCATE_SYSTEM_SPACE, "ER_TRUNCATE_SYSTEM_SPACE"},
		{iproto.ER_LOAD_MODULE, "ER_LOAD_MODULE"},
		{iproto.ER_VINYL_MAX_TUPLE_SIZE, "ER_VINYL_MAX_TUPLE_SIZE"},
		{iproto.ER_WRONG_DD_VERSION, "ER_WRONG_DD_VERSION"},
		{iproto.ER_WRONG_SPACE_FORMAT, "ER_WRONG_SPACE_FORMAT"},
		{iproto.ER_CREATE_SEQUENCE, "ER_CREATE_SEQUENCE"},
		{iproto.ER_ALTER_SEQUENCE, "ER_ALTER_SEQUENCE"},
		{iproto.ER_DROP_SEQUENCE, "ER_DROP_SEQUENCE"},
		{iproto.ER_NO_SUCH_SEQUENCE, "ER_NO_SUCH_SEQUENCE"},
		{iproto.ER_SEQUENCE_EXISTS, "ER_SEQUENCE_EXISTS"},
		{iproto.ER_SEQUENCE_OVERFLOW, "ER_SEQUENCE_OVERFLOW"},
		{iproto.ER_NO_SUCH_INDEX_NAME, "ER_NO_SUCH_INDEX_NAME"},
		{iproto.ER_SPACE_FIELD_IS_DUPLICATE, "ER_SPACE_FIELD_IS_DUPLICATE"},
		{iproto.ER_CANT_CREATE_COLLATION, "ER_CANT_CREATE_COLLATION"},
		{iproto.ER_WRONG_COLLATION_OPTIONS, "ER_WRONG_COLLATION_OPTIONS"},
		{iproto.ER_NULLABLE_PRIMARY, "ER_NULLABLE_PRIMARY"},
		{iproto.ER_NO_SUCH_FIELD_NAME_IN_SPACE, "ER_NO_SUCH_FIELD_NAME_IN_SPACE"},
		{iproto.ER_TRANSACTION_YIELD, "ER_TRANSACTION_YIELD"},
		{iproto.ER_NO_SUCH_GROUP, "ER_NO_SUCH_GROUP"},
		{iproto.ER_SQL_BIND_VALUE, "ER_SQL_BIND_VALUE"},
		{iproto.ER_SQL_BIND_TYPE, "ER_SQL_BIND_TYPE"},
		{iproto.ER_SQL_BIND_PARAMETER_MAX, "ER_SQL_BIND_PARAMETER_MAX"},
		{iproto.ER_SQL_EXECUTE, "ER_SQL_EXECUTE"},
		{iproto.ER_UPDATE_DECIMAL_OVERFLOW, "ER_UPDATE_DECIMAL_OVERFLOW"},
		{iproto.ER_SQL_BIND_NOT_FOUND, "ER_SQL_BIND_NOT_FOUND"},
		{iproto.ER_ACTION_MISMATCH, "ER_ACTION_MISMATCH"},
		{iproto.ER_VIEW_MISSING_SQL, "ER_VIEW_MISSING_SQL"},
		{iproto.ER_FOREIGN_KEY_CONSTRAINT, "ER_FOREIGN_KEY_CONSTRAINT"},
		{iproto.ER_NO_SUCH_MODULE, "ER_NO_SUCH_MODULE"},
		{iproto.ER_NO_SUCH_COLLATION, "ER_NO_SUCH_COLLATION"},
		{iproto.ER_CREATE_FK_CONSTRAINT, "ER_CREATE_FK_CONSTRAINT"},
		{iproto.ER_DROP_FK_CONSTRAINT, "ER_DROP_FK_CONSTRAINT"},
		{iproto.ER_NO_SUCH_CONSTRAINT, "ER_NO_SUCH_CONSTRAINT"},
		{iproto.ER_CONSTRAINT_EXISTS, "ER_CONSTRAINT_EXISTS"},
		{iproto.ER_SQL_TYPE_MISMATCH, "ER_SQL_TYPE_MISMATCH"},
		{iproto.ER_ROWID_OVERFLOW, "ER_ROWID_OVERFLOW"},
		{iproto.ER_DROP_COLLATION, "ER_DROP_COLLATION"},
		{iproto.ER_ILLEGAL_COLLATION_MIX, "ER_ILLEGAL_COLLATION_MIX"},
		{iproto.ER_SQL_NO_SUCH_PRAGMA, "ER_SQL_NO_SUCH_PRAGMA"},
		{iproto.ER_SQL_CANT_RESOLVE_FIELD, "ER_SQL_CANT_RESOLVE_FIELD"},
		{iproto.ER_INDEX_EXISTS_IN_SPACE, "ER_INDEX_EXISTS_IN_SPACE"},
		{iproto.ER_INCONSISTENT_TYPES, "ER_INCONSISTENT_TYPES"},
		{iproto.ER_SQL_SYNTAX_WITH_POS, "ER_SQL_SYNTAX_WITH_POS"},
		{iproto.ER_SQL_STACK_OVERFLOW, "ER_SQL_STACK_OVERFLOW"},
		{iproto.ER_SQL_SELECT_WILDCARD, "ER_SQL_SELECT_WILDCARD"},
		{iproto.ER_SQL_STATEMENT_EMPTY, "ER_SQL_STATEMENT_EMPTY"},
		{iproto.ER_SQL_KEYWORD_IS_RESERVED, "ER_SQL_KEYWORD_IS_RESERVED"},
		{iproto.ER_SQL_SYNTAX_NEAR_TOKEN, "ER_SQL_SYNTAX_NEAR_TOKEN"},
		{iproto.ER_SQL_UNKNOWN_TOKEN, "ER_SQL_UNKNOWN_TOKEN"},
		{iproto.ER_SQL_PARSER_GENERIC, "ER_SQL_PARSER_GENERIC"},
		{iproto.ER_SQL_ANALYZE_ARGUMENT, "ER_SQL_ANALYZE_ARGUMENT"},
		{iproto.ER_SQL_COLUMN_COUNT_MAX, "ER_SQL_COLUMN_COUNT_MAX"},
		{iproto.ER_HEX_LITERAL_MAX, "ER_HEX_LITERAL_MAX"},
		{iproto.ER_INT_LITERAL_MAX, "ER_INT_LITERAL_MAX"},
		{iproto.ER_SQL_PARSER_LIMIT, "ER_SQL_PARSER_LIMIT"},
		{iproto.ER_INDEX_DEF_UNSUPPORTED, "ER_INDEX_DEF_UNSUPPORTED"},
		{iproto.ER_CK_DEF_UNSUPPORTED, "ER_CK_DEF_UNSUPPORTED"},
		{iproto.ER_MULTIKEY_INDEX_MISMATCH, "ER_MULTIKEY_INDEX_MISMATCH"},
		{iproto.ER_CREATE_CK_CONSTRAINT, "ER_CREATE_CK_CONSTRAINT"},
		{iproto.ER_CK_CONSTRAINT_FAILED, "ER_CK_CONSTRAINT_FAILED"},
		{iproto.ER_SQL_COLUMN_COUNT, "ER_SQL_COLUMN_COUNT"},
		{iproto.ER_FUNC_INDEX_FUNC, "ER_FUNC_INDEX_FUNC"},
		{iproto.ER_FUNC_INDEX_FORMAT, "ER_FUNC_INDEX_FORMAT"},
		{iproto.ER_FUNC_INDEX_PARTS, "ER_FUNC_INDEX_PARTS"},
		{iproto.ER_NO_SUCH_FIELD_NAME, "ER_NO_SUCH_FIELD_NAME"},
		{iproto.ER_FUNC_WRONG_ARG_COUNT, "ER_FUNC_WRONG_ARG_COUNT"},
		{iproto.ER_BOOTSTRAP_READONLY, "ER_BOOTSTRAP_READONLY"},
		{iproto.ER_SQL_FUNC_WRONG_RET_COUNT, "ER_SQL_FUNC_WRONG_RET_COUNT"},
		{iproto.ER_FUNC_INVALID_RETURN_TYPE, "ER_FUNC_INVALID_RETURN_TYPE"},
		{iproto.ER_SQL_PARSER_GENERIC_WITH_POS, "ER_SQL_PARSER_GENERIC_WITH_POS"},
		{iproto.ER_REPLICA_NOT_ANON, "ER_REPLICA_NOT_ANON"},
		{iproto.ER_CANNOT_REGISTER, "ER_CANNOT_REGISTER"},
		{iproto.ER_SESSION_SETTING_INVALID_VALUE, "ER_SESSION_SETTING_INVALID_VALUE"},
		{iproto.ER_SQL_PREPARE, "ER_SQL_PREPARE"},
		{iproto.ER_WRONG_QUERY_ID, "ER_WRONG_QUERY_ID"},
		{iproto.ER_SEQUENCE_NOT_STARTED, "ER_SEQUENCE_NOT_STARTED"},
		{iproto.ER_NO_SUCH_SESSION_SETTING, "ER_NO_SUCH_SESSION_SETTING"},
		{iproto.ER_UNCOMMITTED_FOREIGN_SYNC_TXNS, "ER_UNCOMMITTED_FOREIGN_SYNC_TXNS"},
		{iproto.ER_SYNC_MASTER_MISMATCH, "ER_SYNC_MASTER_MISMATCH"},
		{iproto.ER_SYNC_QUORUM_TIMEOUT, "ER_SYNC_QUORUM_TIMEOUT"},
		{iproto.ER_SYNC_ROLLBACK, "ER_SYNC_ROLLBACK"},
		{iproto.ER_TUPLE_METADATA_IS_TOO_BIG, "ER_TUPLE_METADATA_IS_TOO_BIG"},
		{iproto.ER_XLOG_GAP, "ER_XLOG_GAP"},
		{iproto.ER_TOO_EARLY_SUBSCRIBE, "ER_TOO_EARLY_SUBSCRIBE"},
		{iproto.ER_SQL_CANT_ADD_AUTOINC, "ER_SQL_CANT_ADD_AUTOINC"},
		{iproto.ER_QUORUM_WAIT, "ER_QUORUM_WAIT"},
		{iproto.ER_INTERFERING_PROMOTE, "ER_INTERFERING_PROMOTE"},
		{iproto.ER_ELECTION_DISABLED, "ER_ELECTION_DISABLED"},
		{iproto.ER_TXN_ROLLBACK, "ER_TXN_ROLLBACK"},
		{iproto.ER_NOT_LEADER, "ER_NOT_LEADER"},
		{iproto.ER_SYNC_QUEUE_UNCLAIMED, "ER_SYNC_QUEUE_UNCLAIMED"},
		{iproto.ER_SYNC_QUEUE_FOREIGN, "ER_SYNC_QUEUE_FOREIGN"},
		{iproto.ER_UNABLE_TO_PROCESS_IN_STREAM, "ER_UNABLE_TO_PROCESS_IN_STREAM"},
		{iproto.ER_UNABLE_TO_PROCESS_OUT_OF_STREAM, "ER_UNABLE_TO_PROCESS_OUT_OF_STREAM"},
		{iproto.ER_TRANSACTION_TIMEOUT, "ER_TRANSACTION_TIMEOUT"},
		{iproto.ER_ACTIVE_TIMER, "ER_ACTIVE_TIMER"},
		{iproto.ER_TUPLE_FIELD_COUNT_LIMIT, "ER_TUPLE_FIELD_COUNT_LIMIT"},
		{iproto.ER_CREATE_CONSTRAINT, "ER_CREATE_CONSTRAINT"},
		{iproto.ER_FIELD_CONSTRAINT_FAILED, "ER_FIELD_CONSTRAINT_FAILED"},
		{iproto.ER_TUPLE_CONSTRAINT_FAILED, "ER_TUPLE_CONSTRAINT_FAILED"},
		{iproto.ER_CREATE_FOREIGN_KEY, "ER_CREATE_FOREIGN_KEY"},
		{iproto.ER_FOREIGN_KEY_INTEGRITY, "ER_FOREIGN_KEY_INTEGRITY"},
		{iproto.ER_FIELD_FOREIGN_KEY_FAILED, "ER_FIELD_FOREIGN_KEY_FAILED"},
		{iproto.ER_COMPLEX_FOREIGN_KEY_FAILED, "ER_COMPLEX_FOREIGN_KEY_FAILED"},
		{iproto.ER_WRONG_SPACE_UPGRADE_OPTIONS, "ER_WRONG_SPACE_UPGRADE_OPTIONS"},
		{iproto.ER_NO_ELECTION_QUORUM, "ER_NO_ELECTION_QUORUM"},
		{iproto.ER_SSL, "ER_SSL"},
		{iproto.ER_SPLIT_BRAIN, "ER_SPLIT_BRAIN"},
		{iproto.ER_OLD_TERM, "ER_OLD_TERM"},
		{iproto.ER_INTERFERING_ELECTIONS, "ER_INTERFERING_ELECTIONS"},
		{iproto.ER_ITERATOR_POSITION, "ER_ITERATOR_POSITION"},
		{iproto.ER_DEFAULT_VALUE_TYPE, "ER_DEFAULT_VALUE_TYPE"},
		{iproto.ER_UNKNOWN_AUTH_METHOD, "ER_UNKNOWN_AUTH_METHOD"},
		{iproto.ER_INVALID_AUTH_DATA, "ER_INVALID_AUTH_DATA"},
		{iproto.ER_INVALID_AUTH_REQUEST, "ER_INVALID_AUTH_REQUEST"},
		{iproto.ER_WEAK_PASSWORD, "ER_WEAK_PASSWORD"},
		{iproto.ER_OLD_PASSWORD, "ER_OLD_PASSWORD"},
		{iproto.ER_NO_SUCH_SESSION, "ER_NO_SUCH_SESSION"},
		{iproto.ER_WRONG_SESSION_TYPE, "ER_WRONG_SESSION_TYPE"},
		{iproto.ER_PASSWORD_EXPIRED, "ER_PASSWORD_EXPIRED"},
		{iproto.ER_AUTH_DELAY, "ER_AUTH_DELAY"},
		{iproto.ER_AUTH_REQUIRED, "ER_AUTH_REQUIRED"},
		{iproto.ER_SQL_SEQ_SCAN, "ER_SQL_SEQ_SCAN"},
		{iproto.ER_NO_SUCH_EVENT, "ER_NO_SUCH_EVENT"},
		{iproto.ER_BOOTSTRAP_NOT_UNANIMOUS, "ER_BOOTSTRAP_NOT_UNANIMOUS"},
		{iproto.ER_CANT_CHECK_BOOTSTRAP_LEADER, "ER_CANT_CHECK_BOOTSTRAP_LEADER"},
		{iproto.ER_BOOTSTRAP_CONNECTION_NOT_TO_ALL, "ER_BOOTSTRAP_CONNECTION_NOT_TO_ALL"},
		{iproto.ER_NIL_UUID, "ER_NIL_UUID"},
		{iproto.ER_WRONG_FUNCTION_OPTIONS, "ER_WRONG_FUNCTION_OPTIONS"},
		{iproto.ER_MISSING_SYSTEM_SPACES, "ER_MISSING_SYSTEM_SPACES"},
		{iproto.ER_CLUSTER_NAME_MISMATCH, "ER_CLUSTER_NAME_MISMATCH"},
		{iproto.ER_REPLICASET_NAME_MISMATCH, "ER_REPLICASET_NAME_MISMATCH"},
		{iproto.ER_INSTANCE_NAME_DUPLICATE, "ER_INSTANCE_NAME_DUPLICATE"},
		{iproto.ER_INSTANCE_NAME_MISMATCH, "ER_INSTANCE_NAME_MISMATCH"},
		{iproto.ER_SCHEMA_NEEDS_UPGRADE, "ER_SCHEMA_NEEDS_UPGRADE"},
		{iproto.ER_SCHEMA_UPGRADE_IN_PROGRESS, "ER_SCHEMA_UPGRADE_IN_PROGRESS"},
		{iproto.ER_DEPRECATED, "ER_DEPRECATED"},
		{iproto.ER_UNCONFIGURED, "ER_UNCONFIGURED"},
		{iproto.ER_CREATE_DEFAULT_FUNC, "ER_CREATE_DEFAULT_FUNC"},
		{iproto.ER_DEFAULT_FUNC_FAILED, "ER_DEFAULT_FUNC_FAILED"},
		{iproto.ER_INVALID_DEC, "ER_INVALID_DEC"},
		{iproto.ER_IN_ANOTHER_PROMOTE, "ER_IN_ANOTHER_PROMOTE"},
		{iproto.ER_SHUTDOWN, "ER_SHUTDOWN"},
		{iproto.ER_FIELD_VALUE_OUT_OF_RANGE, "ER_FIELD_VALUE_OUT_OF_RANGE"},
		{iproto.ER_REPLICASET_NOT_FOUND, "ER_REPLICASET_NOT_FOUND"},
		{iproto.ER_REPLICASET_NO_WRITABLE, "ER_REPLICASET_NO_WRITABLE"},
		{iproto.ER_REPLICASET_MORE_THAN_ONE_WRITABLE, "ER_REPLICASET_MORE_THAN_ONE_WRITABLE"},
		{iproto.ER_TXN_COMMIT, "ER_TXN_COMMIT"},
		{iproto.ER_READ_VIEW_BUSY, "ER_READ_VIEW_BUSY"},
		{iproto.ER_READ_VIEW_CLOSED, "ER_READ_VIEW_CLOSED"},
		{iproto.ER_WAL_QUEUE_FULL, "ER_WAL_QUEUE_FULL"},
		{iproto.ER_INVALID_VCLOCK, "ER_INVALID_VCLOCK"},
		{iproto.ER_SYNC_QUEUE_FULL, "ER_SYNC_QUEUE_FULL"},
		{iproto.ER_KEY_PART_VALUE_OUT_OF_RANGE, "ER_KEY_PART_VALUE_OUT_OF_RANGE"},
		{iproto.ER_REPLICA_GC, "ER_REPLICA_GC"},
	}

	for i, tc := range cases {
		t.Run(tc.Str, func(t *testing.T) {
			if tc.Err.String() != tc.Str {
				t.Errorf("Got %s, expected %s", tc.Err.String(), tc.Str)
			}
			if int(tc.Err) != i {
				t.Errorf("Got %d, expected %d", tc.Err, i)
			}
		})
	}
}
