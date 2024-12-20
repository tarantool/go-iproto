// Code generated by generate.sh; DO NOT EDIT.

package iproto

// IPROTO error code constants, generated from
// tarantool/src/box/errcode.h
type Error int

const (
	// "Unknown error"
	ER_UNKNOWN Error = 0
	// "%s"
	ER_ILLEGAL_PARAMS Error = 1
	// "Failed to allocate %u bytes in %s for %s"
	ER_MEMORY_ISSUE Error = 2
	// "Duplicate key exists in unique index \"%s\" in space \"%s\" with old tuple - %s and new tuple - %s"
	ER_TUPLE_FOUND Error = 3
	// Unused
	ER_TUPLE_NOT_FOUND Error = 4
	// "%s does not support %s"
	ER_UNSUPPORTED Error = 5
	// Unused
	ER_NONMASTER Error = 6
	// "Can't modify data on a read-only instance"
	ER_READONLY Error = 7
	// "Error injection '%s'"
	ER_INJECTION Error = 8
	// "Failed to create space '%s': %s"
	ER_CREATE_SPACE Error = 9
	// "Space '%s' already exists"
	ER_SPACE_EXISTS Error = 10
	// "Can't drop space '%s': %s"
	ER_DROP_SPACE Error = 11
	// "Can't modify space '%s': %s"
	ER_ALTER_SPACE Error = 12
	// "Unsupported index type supplied for index '%s' in space '%s'"
	ER_INDEX_TYPE Error = 13
	// "Can't create or modify index '%s' in space '%s': %s"
	ER_MODIFY_INDEX Error = 14
	// "Can't drop the primary key in a system space, space '%s'"
	ER_LAST_DROP Error = 15
	// "Tuple format limit reached: %u"
	ER_TUPLE_FORMAT_LIMIT Error = 16
	// "Can't drop primary key in space '%s' while secondary keys exist"
	ER_DROP_PRIMARY_KEY Error = 17
	// "Supplied key type of part %u does not match index part type: expected %s"
	ER_KEY_PART_TYPE Error = 18
	// "Invalid key part count in an exact match (expected %u, got %u)"
	ER_EXACT_MATCH Error = 19
	// "Invalid MsgPack - %s"
	ER_INVALID_MSGPACK Error = 20
	// Unused
	ER_PROC_RET Error = 21
	// "Tuple/Key must be MsgPack array"
	ER_TUPLE_NOT_ARRAY Error = 22
	// "Tuple field %s type does not match one required by operation: expected %s, got %s"
	ER_FIELD_TYPE Error = 23
	// "Field %s has type '%s' in one index, but type '%s' in another"
	ER_INDEX_PART_TYPE_MISMATCH Error = 24
	// "SPLICE error on field %s: %s"
	ER_UPDATE_SPLICE Error = 25
	// "Argument type in operation '%c' on field %s does not match field type: expected %s"
	ER_UPDATE_ARG_TYPE Error = 26
	// "Field %s has type '%s' in space format, but type '%s' in index definition"
	ER_FORMAT_MISMATCH_INDEX_PART Error = 27
	// "Unknown UPDATE operation #%d: %s"
	ER_UNKNOWN_UPDATE_OP Error = 28
	// "Field %s UPDATE error: %s"
	ER_UPDATE_FIELD Error = 29
	// "Transaction is active at return from function"
	ER_FUNCTION_TX_ACTIVE Error = 30
	// "Invalid key part count (expected [0..%u], got %u)"
	ER_KEY_PART_COUNT Error = 31
	// "%s"
	ER_PROC_LUA Error = 32
	// "Procedure '%s' is not defined"
	ER_NO_SUCH_PROC Error = 33
	// "Trigger '%s' doesn't exist"
	ER_NO_SUCH_TRIGGER Error = 34
	// "No index #%u is defined in space '%s'"
	ER_NO_SUCH_INDEX_ID Error = 35
	// "Space '%s' does not exist"
	ER_NO_SUCH_SPACE Error = 36
	// "Field %d was not found in the tuple"
	ER_NO_SUCH_FIELD_NO Error = 37
	// "Tuple field count %u does not match space field count %u"
	ER_EXACT_FIELD_COUNT Error = 38
	// "Tuple field %s required by space format is missing"
	ER_FIELD_MISSING Error = 39
	// "Failed to write to disk"
	ER_WAL_IO Error = 40
	// "Get() doesn't support partial keys and non-unique indexes"
	ER_MORE_THAN_ONE_TUPLE Error = 41
	// "%s access to %s '%s' is denied for user '%s'"
	ER_ACCESS_DENIED Error = 42
	// "Failed to create user '%s': %s"
	ER_CREATE_USER Error = 43
	// "Failed to drop user or role '%s': %s"
	ER_DROP_USER Error = 44
	// "User '%s' is not found"
	ER_NO_SUCH_USER Error = 45
	// "User '%s' already exists"
	ER_USER_EXISTS Error = 46
	// "User not found or supplied credentials are invalid"
	ER_CREDS_MISMATCH Error = 47
	// "Unknown request type %u"
	ER_UNKNOWN_REQUEST_TYPE Error = 48
	// "Unknown object type '%s'"
	ER_UNKNOWN_SCHEMA_OBJECT Error = 49
	// "Failed to create function '%s': %s"
	ER_CREATE_FUNCTION Error = 50
	// "Function '%s' does not exist"
	ER_NO_SUCH_FUNCTION Error = 51
	// "Function '%s' already exists"
	ER_FUNCTION_EXISTS Error = 52
	// "Invalid return value of space:before_replace trigger: expected tuple or nil"
	ER_BEFORE_REPLACE_RET Error = 53
	// "Can not perform %s in a multi-statement transaction"
	ER_MULTISTATEMENT_TRANSACTION Error = 54
	// "Trigger '%s' already exists"
	ER_TRIGGER_EXISTS Error = 55
	// "A limit on the total number of users has been reached: %u"
	ER_USER_MAX Error = 56
	// "Space engine '%s' does not exist"
	ER_NO_SUCH_ENGINE Error = 57
	// "Can't set option '%s' dynamically"
	ER_RELOAD_CFG Error = 58
	// "Incorrect value for option '%s': %s"
	ER_CFG Error = 59
	// "Can not set a savepoint in an empty transaction"
	ER_SAVEPOINT_EMPTY_TX Error = 60
	// "Can not rollback to savepoint: the savepoint does not exist"
	ER_NO_SUCH_SAVEPOINT Error = 61
	// "Replica %s is not registered with replica set %s"
	ER_UNKNOWN_REPLICA Error = 62
	// "Replica set UUID mismatch: expected %s, got %s"
	ER_REPLICASET_UUID_MISMATCH Error = 63
	// "Invalid UUID: %s"
	ER_INVALID_UUID Error = 64
	// "Can't reset replica set UUID: it is already assigned"
	ER_REPLICASET_UUID_IS_RO Error = 65
	// "Instance UUID mismatch: expected %s, got %s"
	ER_INSTANCE_UUID_MISMATCH Error = 66
	// "Can't initialize replica id with a reserved value %u"
	ER_REPLICA_ID_IS_RESERVED Error = 67
	// Unused
	ER_INVALID_ORDER Error = 68
	// "Missing mandatory field '%s' in request"
	ER_MISSING_REQUEST_FIELD Error = 69
	// "Invalid identifier '%s' (expected printable symbols only or it is too long)"
	ER_IDENTIFIER Error = 70
	// "Can't drop function %u: %s"
	ER_DROP_FUNCTION Error = 71
	// "Unknown iterator type '%s'"
	ER_ITERATOR_TYPE Error = 72
	// "Replica count limit reached: %u"
	ER_REPLICA_MAX Error = 73
	// Unused
	ER_INVALID_XLOG Error = 74
	// Unused
	ER_INVALID_XLOG_NAME Error = 75
	// Unused
	ER_INVALID_XLOG_ORDER Error = 76
	// "Connection is not established"
	ER_NO_CONNECTION Error = 77
	// "Timeout exceeded"
	ER_TIMEOUT Error = 78
	// "Operation is not permitted when there is an active transaction "
	ER_ACTIVE_TRANSACTION Error = 79
	// "The transaction the cursor belongs to has ended"
	ER_CURSOR_NO_TRANSACTION Error = 80
	// "Storage engine '%s' does not support cross-engine transactions"
	ER_CROSS_ENGINE_TRANSACTION Error = 81
	// "Role '%s' is not found"
	ER_NO_SUCH_ROLE Error = 82
	// "Role '%s' already exists"
	ER_ROLE_EXISTS Error = 83
	// "Failed to create role '%s': %s"
	ER_CREATE_ROLE Error = 84
	// "Index '%s' already exists"
	ER_INDEX_EXISTS Error = 85
	// "Session is closed"
	ER_SESSION_CLOSED Error = 86
	// "Granting role '%s' to role '%s' would create a loop"
	ER_ROLE_LOOP Error = 87
	// "Incorrect grant arguments: %s"
	ER_GRANT Error = 88
	// "User '%s' already has %s access on %s"
	ER_PRIV_GRANTED Error = 89
	// "User '%s' already has role '%s'"
	ER_ROLE_GRANTED Error = 90
	// "User '%s' does not have %s access on %s '%s'"
	ER_PRIV_NOT_GRANTED Error = 91
	// "User '%s' does not have role '%s'"
	ER_ROLE_NOT_GRANTED Error = 92
	// "Can't find snapshot"
	ER_MISSING_SNAPSHOT Error = 93
	// "Attempt to modify a tuple field which is part of primary index in space '%s'"
	ER_CANT_UPDATE_PRIMARY_KEY Error = 94
	// "Integer overflow when performing '%c' operation on field %s"
	ER_UPDATE_INTEGER_OVERFLOW Error = 95
	// "Setting password for guest user has no effect"
	ER_GUEST_USER_PASSWORD Error = 96
	// "Transaction has been aborted by conflict"
	ER_TRANSACTION_CONFLICT Error = 97
	// "Unsupported %s privilege '%s'"
	ER_UNSUPPORTED_PRIV Error = 98
	// "Failed to dynamically load function '%s': %s"
	ER_LOAD_FUNCTION Error = 99
	// "Unsupported language '%s' specified for function '%s'"
	ER_FUNCTION_LANGUAGE Error = 100
	// "RTree: %s must be an array with %u (point) or %u (rectangle/box) numeric coordinates"
	ER_RTREE_RECT Error = 101
	// "%s"
	ER_PROC_C Error = 102
	// Unused
	ER_UNKNOWN_RTREE_INDEX_DISTANCE_TYPE Error = 103
	// "%s"
	ER_PROTOCOL Error = 104
	// Unused
	ER_UPSERT_UNIQUE_SECONDARY_KEY Error = 105
	// "Wrong record in _index space: got {%s}, expected {%s}"
	ER_WRONG_INDEX_RECORD Error = 106
	// "Wrong index part %u: %s"
	ER_WRONG_INDEX_PARTS Error = 107
	// "Wrong index options: %s"
	ER_WRONG_INDEX_OPTIONS Error = 108
	// "Wrong schema version, current: %d, in request: %llu"
	ER_WRONG_SCHEMA_VERSION Error = 109
	// "Failed to allocate %u bytes for tuple: tuple is too large. Check 'memtx_max_tuple_size' configuration option."
	ER_MEMTX_MAX_TUPLE_SIZE Error = 110
	// "Wrong space options: %s"
	ER_WRONG_SPACE_OPTIONS Error = 111
	// "Index '%s' (%s) of space '%s' (%s) does not support %s"
	ER_UNSUPPORTED_INDEX_FEATURE Error = 112
	// "View '%s' is read-only"
	ER_VIEW_IS_RO Error = 113
	// "No active transaction"
	ER_NO_TRANSACTION Error = 114
	// "%s"
	ER_SYSTEM Error = 115
	// "Instance bootstrap hasn't finished yet"
	ER_LOADING Error = 116
	// "Connection to self"
	ER_CONNECTION_TO_SELF Error = 117
	// Unused
	ER_KEY_PART_IS_TOO_LONG Error = 118
	// "Compression error: %s"
	ER_COMPRESSION Error = 119
	// "Snapshot is already in progress"
	ER_CHECKPOINT_IN_PROGRESS Error = 120
	// "Can not execute a nested statement: nesting limit reached"
	ER_SUB_STMT_MAX Error = 121
	// "Can not commit transaction in a nested statement"
	ER_COMMIT_IN_SUB_STMT Error = 122
	// "Rollback called in a nested statement"
	ER_ROLLBACK_IN_SUB_STMT Error = 123
	// "Decompression error: %s"
	ER_DECOMPRESSION Error = 124
	// "Invalid xlog type: expected %s, got %s"
	ER_INVALID_XLOG_TYPE Error = 125
	// "Failed to lock WAL directory %s and hot_standby mode is off"
	ER_ALREADY_RUNNING Error = 126
	// "Indexed field count limit reached: %d indexed fields"
	ER_INDEX_FIELD_COUNT_LIMIT Error = 127
	// "The local instance id %u is read-only"
	ER_LOCAL_INSTANCE_ID_IS_READ_ONLY Error = 128
	// "Backup is already in progress"
	ER_BACKUP_IN_PROGRESS Error = 129
	// "The read view is aborted"
	ER_READ_VIEW_ABORTED Error = 130
	// "Invalid INDEX file %s: %s"
	ER_INVALID_INDEX_FILE Error = 131
	// "Invalid RUN file: %s"
	ER_INVALID_RUN_FILE Error = 132
	// "Invalid VYLOG file: %s"
	ER_INVALID_VYLOG_FILE Error = 133
	// "WAL has a rollback in progress"
	ER_CASCADE_ROLLBACK Error = 134
	// "Timed out waiting for Vinyl memory quota"
	ER_VY_QUOTA_TIMEOUT Error = 135
	// "%s index  does not support selects via a partial key (expected %u parts, got %u). Please Consider changing index type to TREE."
	ER_PARTIAL_KEY Error = 136
	// "Can't truncate a system space, space '%s'"
	ER_TRUNCATE_SYSTEM_SPACE Error = 137
	// "Failed to dynamically load module '%s': %s"
	ER_LOAD_MODULE Error = 138
	// "Failed to allocate %u bytes for tuple: tuple is too large. Check 'vinyl_max_tuple_size' configuration option."
	ER_VINYL_MAX_TUPLE_SIZE Error = 139
	// "Wrong _schema version: expected 'major.minor[.patch]'"
	ER_WRONG_DD_VERSION Error = 140
	// "Wrong space format field %u: %s"
	ER_WRONG_SPACE_FORMAT Error = 141
	// "Failed to create sequence '%s': %s"
	ER_CREATE_SEQUENCE Error = 142
	// "Can't modify sequence '%s': %s"
	ER_ALTER_SEQUENCE Error = 143
	// "Can't drop sequence '%s': %s"
	ER_DROP_SEQUENCE Error = 144
	// "Sequence '%s' does not exist"
	ER_NO_SUCH_SEQUENCE Error = 145
	// "Sequence '%s' already exists"
	ER_SEQUENCE_EXISTS Error = 146
	// "Sequence '%s' has overflowed"
	ER_SEQUENCE_OVERFLOW Error = 147
	// "No index '%s' is defined in space '%s'"
	ER_NO_SUCH_INDEX_NAME Error = 148
	// "Space field '%s' is duplicate"
	ER_SPACE_FIELD_IS_DUPLICATE Error = 149
	// "Failed to initialize collation: %s."
	ER_CANT_CREATE_COLLATION Error = 150
	// "Wrong collation options: %s"
	ER_WRONG_COLLATION_OPTIONS Error = 151
	// "Primary index of space '%s' can not contain nullable parts"
	ER_NULLABLE_PRIMARY Error = 152
	// "Field '%s' was not found in space '%s' format"
	ER_NO_SUCH_FIELD_NAME_IN_SPACE Error = 153
	// "Transaction has been aborted by a fiber yield"
	ER_TRANSACTION_YIELD Error = 154
	// "Replication group '%s' does not exist"
	ER_NO_SUCH_GROUP Error = 155
	// Unused
	ER_SQL_BIND_VALUE Error = 156
	// "Bind value type %s for parameter %s is not supported"
	ER_SQL_BIND_TYPE Error = 157
	// "SQL bind parameter limit reached: %d"
	ER_SQL_BIND_PARAMETER_MAX Error = 158
	// "Failed to execute SQL statement: %s"
	ER_SQL_EXECUTE Error = 159
	// "Decimal overflow when performing operation '%c' on field %s"
	ER_UPDATE_DECIMAL_OVERFLOW Error = 160
	// "Parameter %s was not found in the statement"
	ER_SQL_BIND_NOT_FOUND Error = 161
	// "Field %s contains %s on conflict action, but %s in index parts"
	ER_ACTION_MISMATCH Error = 162
	// "Space declared as a view must have SQL statement"
	ER_VIEW_MISSING_SQL Error = 163
	// "Can not commit transaction: deferred foreign keys violations are not resolved"
	ER_FOREIGN_KEY_CONSTRAINT Error = 164
	// "Module '%s' does not exist"
	ER_NO_SUCH_MODULE Error = 165
	// "Collation '%s' does not exist"
	ER_NO_SUCH_COLLATION Error = 166
	// "Failed to create foreign key constraint '%s': %s"
	ER_CREATE_FK_CONSTRAINT Error = 167
	// Unused
	ER_DROP_FK_CONSTRAINT Error = 168
	// "Constraint '%s' does not exist in space '%s'"
	ER_NO_SUCH_CONSTRAINT Error = 169
	// "%s constraint '%s' already exists in space '%s'"
	ER_CONSTRAINT_EXISTS Error = 170
	// "Type mismatch: can not convert %s to %s"
	ER_SQL_TYPE_MISMATCH Error = 171
	// "Rowid is overflowed: too many entries in ephemeral space"
	ER_ROWID_OVERFLOW Error = 172
	// "Can't drop collation '%s': %s"
	ER_DROP_COLLATION Error = 173
	// "Illegal mix of collations"
	ER_ILLEGAL_COLLATION_MIX Error = 174
	// "Pragma '%s' does not exist"
	ER_SQL_NO_SUCH_PRAGMA Error = 175
	// "Can't resolve field '%s'"
	ER_SQL_CANT_RESOLVE_FIELD Error = 176
	// "Index '%s' already exists in space '%s'"
	ER_INDEX_EXISTS_IN_SPACE Error = 177
	// "Inconsistent types: expected %s got %s"
	ER_INCONSISTENT_TYPES Error = 178
	// "Syntax error at line %d at or near position %d: %s"
	ER_SQL_SYNTAX_WITH_POS Error = 179
	// "Failed to parse SQL statement: parser stack limit reached"
	ER_SQL_STACK_OVERFLOW Error = 180
	// "Failed to expand '*' in SELECT statement without FROM clause"
	ER_SQL_SELECT_WILDCARD Error = 181
	// "Failed to execute an empty SQL statement"
	ER_SQL_STATEMENT_EMPTY Error = 182
	// "At line %d at or near position %d: keyword '%s' is reserved. Please use double quotes if '%s' is an identifier."
	ER_SQL_KEYWORD_IS_RESERVED Error = 183
	// "Syntax error at line %d near '%s'"
	ER_SQL_SYNTAX_NEAR_TOKEN Error = 184
	// "At line %d at or near position %d: unrecognized token '%s'"
	ER_SQL_UNKNOWN_TOKEN Error = 185
	// "%s"
	ER_SQL_PARSER_GENERIC Error = 186
	// Unused
	ER_SQL_ANALYZE_ARGUMENT Error = 187
	// "Failed to create space '%s': space column count %d exceeds the limit (%d)"
	ER_SQL_COLUMN_COUNT_MAX Error = 188
	// "Hex literal %s length %d exceeds the supported limit (%d)"
	ER_HEX_LITERAL_MAX Error = 189
	// "Integer literal %s exceeds the supported range [-9223372036854775808, 18446744073709551615]"
	ER_INT_LITERAL_MAX Error = 190
	// "%s %d exceeds the limit (%d)"
	ER_SQL_PARSER_LIMIT Error = 191
	// "%s are prohibited in an index definition"
	ER_INDEX_DEF_UNSUPPORTED Error = 192
	// Unused
	ER_CK_DEF_UNSUPPORTED Error = 193
	// "Field %s is used as multikey in one index and as single key in another"
	ER_MULTIKEY_INDEX_MISMATCH Error = 194
	// "Failed to create check constraint '%s': %s"
	ER_CREATE_CK_CONSTRAINT Error = 195
	// Unused
	ER_CK_CONSTRAINT_FAILED Error = 196
	// "Unequal number of entries in row expression: left side has %u, but right side - %u"
	ER_SQL_COLUMN_COUNT Error = 197
	// "Failed to build a key for functional index '%s' of space '%s': %s"
	ER_FUNC_INDEX_FUNC Error = 198
	// "Key format doesn't match one defined in functional index '%s' of space '%s': %s"
	ER_FUNC_INDEX_FORMAT Error = 199
	// "Wrong functional index definition: %s"
	ER_FUNC_INDEX_PARTS Error = 200
	// "Field '%s' was not found in the tuple"
	ER_NO_SUCH_FIELD_NAME Error = 201
	// "Wrong number of arguments is passed to %s(): expected %s, got %d"
	ER_FUNC_WRONG_ARG_COUNT Error = 202
	// "Trying to bootstrap a local read-only instance as master"
	ER_BOOTSTRAP_READONLY Error = 203
	// "SQL expects exactly one argument returned from %s, got %d"
	ER_SQL_FUNC_WRONG_RET_COUNT Error = 204
	// "Function '%s' returned value of invalid type: expected %s got %s"
	ER_FUNC_INVALID_RETURN_TYPE Error = 205
	// "At line %d at or near position %d: %s"
	ER_SQL_PARSER_GENERIC_WITH_POS Error = 206
	// "Replica '%s' is not anonymous and cannot register."
	ER_REPLICA_NOT_ANON Error = 207
	// "Couldn't find an instance to register this replica on."
	ER_CANNOT_REGISTER Error = 208
	// "Session setting %s expected a value of type %s"
	ER_SESSION_SETTING_INVALID_VALUE Error = 209
	// "Failed to prepare SQL statement: %s"
	ER_SQL_PREPARE Error = 210
	// "Prepared statement with id %u does not exist"
	ER_WRONG_QUERY_ID Error = 211
	// "Sequence '%s' is not started"
	ER_SEQUENCE_NOT_STARTED Error = 212
	// "Session setting %s doesn't exist"
	ER_NO_SUCH_SESSION_SETTING Error = 213
	// "Found uncommitted sync transactions from other instance with id %u"
	ER_UNCOMMITTED_FOREIGN_SYNC_TXNS Error = 214
	// Unused
	ER_SYNC_MASTER_MISMATCH Error = 215
	// "Quorum collection for a synchronous transaction is timed out"
	ER_SYNC_QUORUM_TIMEOUT Error = 216
	// "A rollback for a synchronous transaction is received"
	ER_SYNC_ROLLBACK Error = 217
	// "Can't create tuple: metadata size %u is too big"
	ER_TUPLE_METADATA_IS_TOO_BIG Error = 218
	// "%s"
	ER_XLOG_GAP Error = 219
	// "Can't subscribe non-anonymous replica %s until join is done"
	ER_TOO_EARLY_SUBSCRIBE Error = 220
	// "Can't add AUTOINCREMENT: space %s can't feature more than one AUTOINCREMENT field"
	ER_SQL_CANT_ADD_AUTOINC Error = 221
	// "Couldn't wait for quorum %d: %s"
	ER_QUORUM_WAIT Error = 222
	// "Instance with replica id %u was promoted first"
	ER_INTERFERING_PROMOTE Error = 223
	// "Elections were turned off"
	ER_ELECTION_DISABLED Error = 224
	// "Transaction was rolled back"
	ER_TXN_ROLLBACK Error = 225
	// "The instance is not a leader. New leader is %u"
	ER_NOT_LEADER Error = 226
	// "The synchronous transaction queue doesn't belong to any instance"
	ER_SYNC_QUEUE_UNCLAIMED Error = 227
	// "The synchronous transaction queue belongs to other instance with id %u"
	ER_SYNC_QUEUE_FOREIGN Error = 228
	// "Unable to process %s request in stream"
	ER_UNABLE_TO_PROCESS_IN_STREAM Error = 229
	// "Unable to process %s request out of stream"
	ER_UNABLE_TO_PROCESS_OUT_OF_STREAM Error = 230
	// "Transaction has been aborted by timeout"
	ER_TRANSACTION_TIMEOUT Error = 231
	// "Operation is not permitted if timer is already running"
	ER_ACTIVE_TIMER Error = 232
	// "Tuple field count limit reached: see box.schema.FIELD_MAX"
	ER_TUPLE_FIELD_COUNT_LIMIT Error = 233
	// "Failed to create constraint '%s' in space '%s': %s"
	ER_CREATE_CONSTRAINT Error = 234
	// "Check constraint '%s' failed for field '%s'"
	ER_FIELD_CONSTRAINT_FAILED Error = 235
	// "Check constraint '%s' failed for a tuple"
	ER_TUPLE_CONSTRAINT_FAILED Error = 236
	// "Failed to create foreign key '%s' in space '%s': %s"
	ER_CREATE_FOREIGN_KEY Error = 237
	// "Foreign key '%s' integrity check failed: %s"
	ER_FOREIGN_KEY_INTEGRITY Error = 238
	// "Foreign key constraint '%s' failed for field '%s': %s"
	ER_FIELD_FOREIGN_KEY_FAILED Error = 239
	// "Foreign key constraint '%s' failed: %s"
	ER_COMPLEX_FOREIGN_KEY_FAILED Error = 240
	// "Wrong space upgrade options: %s"
	ER_WRONG_SPACE_UPGRADE_OPTIONS Error = 241
	// "Not enough peers connected to start elections: %d out of minimal required %d"
	ER_NO_ELECTION_QUORUM Error = 242
	// "%s"
	ER_SSL Error = 243
	// "Split-Brain discovered: %s"
	ER_SPLIT_BRAIN Error = 244
	// "The term is outdated: old - %llu, new - %llu"
	ER_OLD_TERM Error = 245
	// "Interfering elections started"
	ER_INTERFERING_ELECTIONS Error = 246
	// "Iterator position is invalid"
	ER_ITERATOR_POSITION Error = 247
	// "Type of the default value does not match tuple field %s type: expected %s, got %s"
	ER_DEFAULT_VALUE_TYPE Error = 248
	// "Unknown authentication method '%s'"
	ER_UNKNOWN_AUTH_METHOD Error = 249
	// "Invalid '%s' data: %s"
	ER_INVALID_AUTH_DATA Error = 250
	// "Invalid '%s' request: %s"
	ER_INVALID_AUTH_REQUEST Error = 251
	// "Password doesn't meet security requirements: %s"
	ER_WEAK_PASSWORD Error = 252
	// "Password must differ from last %d passwords"
	ER_OLD_PASSWORD Error = 253
	// "Session %llu does not exist"
	ER_NO_SUCH_SESSION Error = 254
	// "Session '%s' is not supported"
	ER_WRONG_SESSION_TYPE Error = 255
	// "Password expired"
	ER_PASSWORD_EXPIRED Error = 256
	// "Too many authentication attempts"
	ER_AUTH_DELAY Error = 257
	// "Authentication required"
	ER_AUTH_REQUIRED Error = 258
	// "Scanning is not allowed for %s"
	ER_SQL_SEQ_SCAN Error = 259
	// "Unknown event %s"
	ER_NO_SUCH_EVENT Error = 260
	// "Replica %s chose a different bootstrap leader %s"
	ER_BOOTSTRAP_NOT_UNANIMOUS Error = 261
	// "Can't check who replica %s chose its bootstrap leader"
	ER_CANT_CHECK_BOOTSTRAP_LEADER Error = 262
	// "Some replica set members were not specified in box.cfg.replication"
	ER_BOOTSTRAP_CONNECTION_NOT_TO_ALL Error = 263
	// "Nil UUID is reserved and can't be used in replication"
	ER_NIL_UUID Error = 264
	// "Wrong function options: %s"
	ER_WRONG_FUNCTION_OPTIONS Error = 265
	// "Snapshot has no system spaces"
	ER_MISSING_SYSTEM_SPACES Error = 266
	// "Cluster name mismatch: name '%s' provided in config confilcts with the instance one '%s'"
	ER_CLUSTER_NAME_MISMATCH Error = 267
	// "Replicaset name mismatch: name '%s' provided in config confilcts with the instance one '%s'"
	ER_REPLICASET_NAME_MISMATCH Error = 268
	// "Duplicate replica name %s, already occupied by %s"
	ER_INSTANCE_NAME_DUPLICATE Error = 269
	// "Instance name mismatch: name '%s' provided in config confilcts with the instance one '%s'"
	ER_INSTANCE_NAME_MISMATCH Error = 270
	// "Your schema version is %s while Tarantool feature %s requires schema version %s or higher. Please, consider using box.schema.upgrade()."
	ER_SCHEMA_NEEDS_UPGRADE Error = 271
	// "Schema upgrade is already in progress"
	ER_SCHEMA_UPGRADE_IN_PROGRESS Error = 272
	// "%s is deprecated"
	ER_DEPRECATED Error = 273
	// "Please call box.cfg{} first"
	ER_UNCONFIGURED Error = 274
	// "Failed to create field default function '%s': %s"
	ER_CREATE_DEFAULT_FUNC Error = 275
	// "Error calling field default function '%s': %s"
	ER_DEFAULT_FUNC_FAILED Error = 276
	// "Invalid decimal: '%s'"
	ER_INVALID_DEC Error = 277
	// "box.ctl.promote() is already running"
	ER_IN_ANOTHER_PROMOTE Error = 278
	// "Server is shutting down"
	ER_SHUTDOWN Error = 279
	// "The value of field %s exceeds the supported range for type '%s': %s"
	ER_FIELD_VALUE_OUT_OF_RANGE Error = 280
	// "The replicaset was not found by its name"
	ER_REPLICASET_NOT_FOUND Error = 281
	// "Writable instance was not found in replicaset"
	ER_REPLICASET_NO_WRITABLE Error = 282
	// "More than one writable was found in replicaset"
	ER_REPLICASET_MORE_THAN_ONE_WRITABLE Error = 283
	// "Transaction was committed"
	ER_TXN_COMMIT Error = 284
	// "The read view is busy"
	ER_READ_VIEW_BUSY Error = 285
	// "The read view is closed"
	ER_READ_VIEW_CLOSED Error = 286
	// "The WAL queue is full"
	ER_WAL_QUEUE_FULL Error = 287
	// "Invalid vclock"
	ER_INVALID_VCLOCK Error = 288
	// "The synchronous transaction queue is full"
	ER_SYNC_QUEUE_FULL Error = 289
	// "The value of key part exceeds the supported range for type"
	ER_KEY_PART_VALUE_OUT_OF_RANGE Error = 290
	// "Cannot clean up replica's resources"
	ER_REPLICA_GC Error = 291
)
