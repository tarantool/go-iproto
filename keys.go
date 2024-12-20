// Code generated by generate.sh; DO NOT EDIT.

package iproto

// IPROTO key constants, generated from
// tarantool/src/box/iproto_constants.h
type Key int

const (
	IPROTO_REQUEST_TYPE Key = 0x00
	IPROTO_SYNC         Key = 0x01
	// Replication keys (header)
	IPROTO_REPLICA_ID     Key = 0x02
	IPROTO_LSN            Key = 0x03
	IPROTO_TIMESTAMP      Key = 0x04
	IPROTO_SCHEMA_VERSION Key = 0x05
	IPROTO_SERVER_VERSION Key = 0x06
	IPROTO_GROUP_ID       Key = 0x07
	IPROTO_TSN            Key = 0x08
	IPROTO_FLAGS          Key = 0x09
	IPROTO_STREAM_ID      Key = 0x0a
	// Leave a gap for other keys in the header.
	IPROTO_SPACE_ID   Key = 0x10
	IPROTO_INDEX_ID   Key = 0x11
	IPROTO_LIMIT      Key = 0x12
	IPROTO_OFFSET     Key = 0x13
	IPROTO_ITERATOR   Key = 0x14
	IPROTO_INDEX_BASE Key = 0x15
	// Leave a gap between integer values and other keys
	// Flag indicating the need to send position of
	// last selected tuple in response.
	IPROTO_FETCH_POSITION Key = 0x1f
	IPROTO_KEY            Key = 0x20
	IPROTO_TUPLE          Key = 0x21
	IPROTO_FUNCTION_NAME  Key = 0x22
	IPROTO_USER_NAME      Key = 0x23
	// Replication keys (body).
	// Unfortunately, there is no gap between request and
	// replication keys (between USER_NAME and INSTANCE_UUID).
	// So imagine, that OPS, EXPR and FIELD_NAME keys follows
	// the USER_NAME key.
	IPROTO_INSTANCE_UUID   Key = 0x24
	IPROTO_REPLICASET_UUID Key = 0x25
	IPROTO_VCLOCK          Key = 0x26
	// Also request keys. See the comment above.
	IPROTO_EXPR Key = 0x27
	// UPSERT but not UPDATE ops, because of legacy
	IPROTO_OPS        Key = 0x28
	IPROTO_BALLOT     Key = 0x29
	IPROTO_TUPLE_META Key = 0x2a
	IPROTO_OPTIONS    Key = 0x2b
	// Old tuple (i.e. before DML request is applied).
	IPROTO_OLD_TUPLE Key = 0x2c
	// New tuple (i.e. result of DML request).
	IPROTO_NEW_TUPLE Key = 0x2d
	// Position of last selected tuple to start iteration after it.
	IPROTO_AFTER_POSITION Key = 0x2e
	// Last selected tuple to start iteration after it.
	IPROTO_AFTER_TUPLE Key = 0x2f
	// Response keys.
	IPROTO_DATA     Key = 0x30
	IPROTO_ERROR_24 Key = 0x31
	// IPROTO_METADATA: [
	//      { IPROTO_FIELD_NAME: name },
	//      { ... },
	//      ...
	// ]
	IPROTO_METADATA      Key = 0x32
	IPROTO_BIND_METADATA Key = 0x33
	IPROTO_BIND_COUNT    Key = 0x34
	// Position of last selected tuple in response.
	IPROTO_POSITION Key = 0x35
	// Also request keys. See the comment above.
	// The data in Arrow format.
	IPROTO_ARROW Key = 0x36
	// Leave a gap between response keys and SQL keys.
	IPROTO_SQL_TEXT Key = 0x40
	IPROTO_SQL_BIND Key = 0x41
	// IPROTO_SQL_INFO: {
	//      SQL_INFO_ROW_COUNT: number
	// }
	IPROTO_SQL_INFO Key = 0x42
	IPROTO_STMT_ID  Key = 0x43
	// Leave a gap between SQL keys and additional request keys
	IPROTO_REPLICA_ANON Key = 0x50
	IPROTO_ID_FILTER    Key = 0x51
	IPROTO_ERROR        Key = 0x52
	// Term. Has the same meaning as IPROTO_RAFT_TERM, but is an iproto
	// key, rather than a raft key. Used for PROMOTE request, which needs
	// both iproto (e.g. REPLICA_ID) and raft (RAFT_TERM) keys.
	IPROTO_TERM Key = 0x53
	// Protocol version.
	IPROTO_VERSION Key = 0x54
	// Protocol features.
	IPROTO_FEATURES Key = 0x55
	// Operation timeout. Specific to request type.
	IPROTO_TIMEOUT Key = 0x56
	// Key name and data sent to a remote watcher.
	IPROTO_EVENT_KEY  Key = 0x57
	IPROTO_EVENT_DATA Key = 0x58
	// Isolation level, is used only by IPROTO_BEGIN request.
	IPROTO_TXN_ISOLATION Key = 0x59
	// A vclock synchronisation request identifier.
	IPROTO_VCLOCK_SYNC Key = 0x5a
	// Name of the authentication method that is currently used on
	// the server (value of box.cfg.auth_type). It's sent in reply
	// to IPROTO_ID request. A client can use it as the default
	// authentication method.
	IPROTO_AUTH_TYPE       Key = 0x5b
	IPROTO_REPLICASET_NAME Key = 0x5c
	IPROTO_INSTANCE_NAME   Key = 0x5d
	// Space name used instead of identifier (IPROTO_SPACE_ID) in DML
	// requests. Preferred when identifier is present (i.e., the identifier
	// is ignored).
	IPROTO_SPACE_NAME Key = 0x5e
	// Index name used instead of identifier (IPROTO_INDEX_ID) in
	// IPROTO_SELECT, IPROTO_UPDATE, and IPROTO_DELETE requests. Preferred
	// when identifier is present (i.e., the identifier is ignored).
	IPROTO_INDEX_NAME Key = 0x5f
	// Mapping of format identifier to format clause consisting of field
	// names and field types.
	IPROTO_TUPLE_FORMATS Key = 0x60
	// Flag indicating whether the transaction is synchronous.
	IPROTO_IS_SYNC Key = 0x61
	// Flag indicating whether checkpoint join should be done.
	IPROTO_IS_CHECKPOINT_JOIN Key = 0x62
	// Shows the signature of the checkpoint to read from.
	// Requires CHECKPOINT_JOIN to be true.
	IPROTO_CHECKPOINT_VCLOCK Key = 0x63
	// Shows the lsn to start sending from. Server sends all rows
	// >= IPROTO_CHECKPOINT_LSN. Requires CHECKPOINT_JOIN to be
	// true and CHECKPOINT_VCLOCK to be set.
	IPROTO_CHECKPOINT_LSN Key = 0x64
)

// IPROTO metadata key constants, generated from
// tarantool/src/box/iproto_constants.h
type MetadataKey int

const (
	IPROTO_FIELD_NAME             MetadataKey = 0
	IPROTO_FIELD_TYPE             MetadataKey = 1
	IPROTO_FIELD_COLL             MetadataKey = 2
	IPROTO_FIELD_IS_NULLABLE      MetadataKey = 3
	IPROTO_FIELD_IS_AUTOINCREMENT MetadataKey = 4
	IPROTO_FIELD_SPAN             MetadataKey = 5
)

// IPROTO ballot key constants, generated from
// tarantool/src/box/iproto_constants.h
type BallotKey int

const (
	IPROTO_BALLOT_IS_RO_CFG                BallotKey = 0x01
	IPROTO_BALLOT_VCLOCK                   BallotKey = 0x02
	IPROTO_BALLOT_GC_VCLOCK                BallotKey = 0x03
	IPROTO_BALLOT_IS_RO                    BallotKey = 0x04
	IPROTO_BALLOT_IS_ANON                  BallotKey = 0x05
	IPROTO_BALLOT_IS_BOOTED                BallotKey = 0x06
	IPROTO_BALLOT_CAN_LEAD                 BallotKey = 0x07
	IPROTO_BALLOT_BOOTSTRAP_LEADER_UUID    BallotKey = 0x08
	IPROTO_BALLOT_REGISTERED_REPLICA_UUIDS BallotKey = 0x09
	IPROTO_BALLOT_INSTANCE_NAME            BallotKey = 0x0a
)

// IPROTO raft key constants, generated from
// tarantool/src/box/iproto_constants.h
type RaftKey int

const (
	IPROTO_RAFT_TERM           RaftKey = 0
	IPROTO_RAFT_VOTE           RaftKey = 1
	IPROTO_RAFT_STATE          RaftKey = 2
	IPROTO_RAFT_VCLOCK         RaftKey = 3
	IPROTO_RAFT_LEADER_ID      RaftKey = 4
	IPROTO_RAFT_IS_LEADER_SEEN RaftKey = 5
)

// IPROTO SQL info key constants, generated from
// tarantool/src/box/execute.h
type SqlInfoKey int

const (
	SQL_INFO_ROW_COUNT         SqlInfoKey = 0
	SQL_INFO_AUTOINCREMENT_IDS SqlInfoKey = 1
)
