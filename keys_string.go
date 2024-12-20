// Code generated by "stringer -type=Key,MetadataKey,BallotKey,RaftKey,SqlInfoKey -output=keys_string.go"; DO NOT EDIT.

package iproto

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_REQUEST_TYPE-0]
	_ = x[IPROTO_SYNC-1]
	_ = x[IPROTO_REPLICA_ID-2]
	_ = x[IPROTO_LSN-3]
	_ = x[IPROTO_TIMESTAMP-4]
	_ = x[IPROTO_SCHEMA_VERSION-5]
	_ = x[IPROTO_SERVER_VERSION-6]
	_ = x[IPROTO_GROUP_ID-7]
	_ = x[IPROTO_TSN-8]
	_ = x[IPROTO_FLAGS-9]
	_ = x[IPROTO_STREAM_ID-10]
	_ = x[IPROTO_SPACE_ID-16]
	_ = x[IPROTO_INDEX_ID-17]
	_ = x[IPROTO_LIMIT-18]
	_ = x[IPROTO_OFFSET-19]
	_ = x[IPROTO_ITERATOR-20]
	_ = x[IPROTO_INDEX_BASE-21]
	_ = x[IPROTO_FETCH_POSITION-31]
	_ = x[IPROTO_KEY-32]
	_ = x[IPROTO_TUPLE-33]
	_ = x[IPROTO_FUNCTION_NAME-34]
	_ = x[IPROTO_USER_NAME-35]
	_ = x[IPROTO_INSTANCE_UUID-36]
	_ = x[IPROTO_REPLICASET_UUID-37]
	_ = x[IPROTO_VCLOCK-38]
	_ = x[IPROTO_EXPR-39]
	_ = x[IPROTO_OPS-40]
	_ = x[IPROTO_BALLOT-41]
	_ = x[IPROTO_TUPLE_META-42]
	_ = x[IPROTO_OPTIONS-43]
	_ = x[IPROTO_OLD_TUPLE-44]
	_ = x[IPROTO_NEW_TUPLE-45]
	_ = x[IPROTO_AFTER_POSITION-46]
	_ = x[IPROTO_AFTER_TUPLE-47]
	_ = x[IPROTO_DATA-48]
	_ = x[IPROTO_ERROR_24-49]
	_ = x[IPROTO_METADATA-50]
	_ = x[IPROTO_BIND_METADATA-51]
	_ = x[IPROTO_BIND_COUNT-52]
	_ = x[IPROTO_POSITION-53]
	_ = x[IPROTO_ARROW-54]
	_ = x[IPROTO_SQL_TEXT-64]
	_ = x[IPROTO_SQL_BIND-65]
	_ = x[IPROTO_SQL_INFO-66]
	_ = x[IPROTO_STMT_ID-67]
	_ = x[IPROTO_REPLICA_ANON-80]
	_ = x[IPROTO_ID_FILTER-81]
	_ = x[IPROTO_ERROR-82]
	_ = x[IPROTO_TERM-83]
	_ = x[IPROTO_VERSION-84]
	_ = x[IPROTO_FEATURES-85]
	_ = x[IPROTO_TIMEOUT-86]
	_ = x[IPROTO_EVENT_KEY-87]
	_ = x[IPROTO_EVENT_DATA-88]
	_ = x[IPROTO_TXN_ISOLATION-89]
	_ = x[IPROTO_VCLOCK_SYNC-90]
	_ = x[IPROTO_AUTH_TYPE-91]
	_ = x[IPROTO_REPLICASET_NAME-92]
	_ = x[IPROTO_INSTANCE_NAME-93]
	_ = x[IPROTO_SPACE_NAME-94]
	_ = x[IPROTO_INDEX_NAME-95]
	_ = x[IPROTO_TUPLE_FORMATS-96]
	_ = x[IPROTO_IS_SYNC-97]
	_ = x[IPROTO_IS_CHECKPOINT_JOIN-98]
	_ = x[IPROTO_CHECKPOINT_VCLOCK-99]
	_ = x[IPROTO_CHECKPOINT_LSN-100]
}

const (
	_Key_name_0 = "IPROTO_REQUEST_TYPEIPROTO_SYNCIPROTO_REPLICA_IDIPROTO_LSNIPROTO_TIMESTAMPIPROTO_SCHEMA_VERSIONIPROTO_SERVER_VERSIONIPROTO_GROUP_IDIPROTO_TSNIPROTO_FLAGSIPROTO_STREAM_ID"
	_Key_name_1 = "IPROTO_SPACE_IDIPROTO_INDEX_IDIPROTO_LIMITIPROTO_OFFSETIPROTO_ITERATORIPROTO_INDEX_BASE"
	_Key_name_2 = "IPROTO_FETCH_POSITIONIPROTO_KEYIPROTO_TUPLEIPROTO_FUNCTION_NAMEIPROTO_USER_NAMEIPROTO_INSTANCE_UUIDIPROTO_REPLICASET_UUIDIPROTO_VCLOCKIPROTO_EXPRIPROTO_OPSIPROTO_BALLOTIPROTO_TUPLE_METAIPROTO_OPTIONSIPROTO_OLD_TUPLEIPROTO_NEW_TUPLEIPROTO_AFTER_POSITIONIPROTO_AFTER_TUPLEIPROTO_DATAIPROTO_ERROR_24IPROTO_METADATAIPROTO_BIND_METADATAIPROTO_BIND_COUNTIPROTO_POSITIONIPROTO_ARROW"
	_Key_name_3 = "IPROTO_SQL_TEXTIPROTO_SQL_BINDIPROTO_SQL_INFOIPROTO_STMT_ID"
	_Key_name_4 = "IPROTO_REPLICA_ANONIPROTO_ID_FILTERIPROTO_ERRORIPROTO_TERMIPROTO_VERSIONIPROTO_FEATURESIPROTO_TIMEOUTIPROTO_EVENT_KEYIPROTO_EVENT_DATAIPROTO_TXN_ISOLATIONIPROTO_VCLOCK_SYNCIPROTO_AUTH_TYPEIPROTO_REPLICASET_NAMEIPROTO_INSTANCE_NAMEIPROTO_SPACE_NAMEIPROTO_INDEX_NAMEIPROTO_TUPLE_FORMATSIPROTO_IS_SYNCIPROTO_IS_CHECKPOINT_JOINIPROTO_CHECKPOINT_VCLOCKIPROTO_CHECKPOINT_LSN"
)

var (
	_Key_index_0 = [...]uint8{0, 19, 30, 47, 57, 73, 94, 115, 130, 140, 152, 168}
	_Key_index_1 = [...]uint8{0, 15, 30, 42, 55, 70, 87}
	_Key_index_2 = [...]uint16{0, 21, 31, 43, 63, 79, 99, 121, 134, 145, 155, 168, 185, 199, 215, 231, 252, 270, 281, 296, 311, 331, 348, 363, 375}
	_Key_index_3 = [...]uint8{0, 15, 30, 45, 59}
	_Key_index_4 = [...]uint16{0, 19, 35, 47, 58, 72, 87, 101, 117, 134, 154, 172, 188, 210, 230, 247, 264, 284, 298, 323, 347, 368}
)

func (i Key) String() string {
	switch {
	case 0 <= i && i <= 10:
		return _Key_name_0[_Key_index_0[i]:_Key_index_0[i+1]]
	case 16 <= i && i <= 21:
		i -= 16
		return _Key_name_1[_Key_index_1[i]:_Key_index_1[i+1]]
	case 31 <= i && i <= 54:
		i -= 31
		return _Key_name_2[_Key_index_2[i]:_Key_index_2[i+1]]
	case 64 <= i && i <= 67:
		i -= 64
		return _Key_name_3[_Key_index_3[i]:_Key_index_3[i+1]]
	case 80 <= i && i <= 100:
		i -= 80
		return _Key_name_4[_Key_index_4[i]:_Key_index_4[i+1]]
	default:
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_FIELD_NAME-0]
	_ = x[IPROTO_FIELD_TYPE-1]
	_ = x[IPROTO_FIELD_COLL-2]
	_ = x[IPROTO_FIELD_IS_NULLABLE-3]
	_ = x[IPROTO_FIELD_IS_AUTOINCREMENT-4]
	_ = x[IPROTO_FIELD_SPAN-5]
}

const _MetadataKey_name = "IPROTO_FIELD_NAMEIPROTO_FIELD_TYPEIPROTO_FIELD_COLLIPROTO_FIELD_IS_NULLABLEIPROTO_FIELD_IS_AUTOINCREMENTIPROTO_FIELD_SPAN"

var _MetadataKey_index = [...]uint8{0, 17, 34, 51, 75, 104, 121}

func (i MetadataKey) String() string {
	if i < 0 || i >= MetadataKey(len(_MetadataKey_index)-1) {
		return "MetadataKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MetadataKey_name[_MetadataKey_index[i]:_MetadataKey_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_BALLOT_IS_RO_CFG-1]
	_ = x[IPROTO_BALLOT_VCLOCK-2]
	_ = x[IPROTO_BALLOT_GC_VCLOCK-3]
	_ = x[IPROTO_BALLOT_IS_RO-4]
	_ = x[IPROTO_BALLOT_IS_ANON-5]
	_ = x[IPROTO_BALLOT_IS_BOOTED-6]
	_ = x[IPROTO_BALLOT_CAN_LEAD-7]
	_ = x[IPROTO_BALLOT_BOOTSTRAP_LEADER_UUID-8]
	_ = x[IPROTO_BALLOT_REGISTERED_REPLICA_UUIDS-9]
	_ = x[IPROTO_BALLOT_INSTANCE_NAME-10]
}

const _BallotKey_name = "IPROTO_BALLOT_IS_RO_CFGIPROTO_BALLOT_VCLOCKIPROTO_BALLOT_GC_VCLOCKIPROTO_BALLOT_IS_ROIPROTO_BALLOT_IS_ANONIPROTO_BALLOT_IS_BOOTEDIPROTO_BALLOT_CAN_LEADIPROTO_BALLOT_BOOTSTRAP_LEADER_UUIDIPROTO_BALLOT_REGISTERED_REPLICA_UUIDSIPROTO_BALLOT_INSTANCE_NAME"

var _BallotKey_index = [...]uint8{0, 23, 43, 66, 85, 106, 129, 151, 186, 224, 251}

func (i BallotKey) String() string {
	i -= 1
	if i < 0 || i >= BallotKey(len(_BallotKey_index)-1) {
		return "BallotKey(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BallotKey_name[_BallotKey_index[i]:_BallotKey_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_RAFT_TERM-0]
	_ = x[IPROTO_RAFT_VOTE-1]
	_ = x[IPROTO_RAFT_STATE-2]
	_ = x[IPROTO_RAFT_VCLOCK-3]
	_ = x[IPROTO_RAFT_LEADER_ID-4]
	_ = x[IPROTO_RAFT_IS_LEADER_SEEN-5]
}

const _RaftKey_name = "IPROTO_RAFT_TERMIPROTO_RAFT_VOTEIPROTO_RAFT_STATEIPROTO_RAFT_VCLOCKIPROTO_RAFT_LEADER_IDIPROTO_RAFT_IS_LEADER_SEEN"

var _RaftKey_index = [...]uint8{0, 16, 32, 49, 67, 88, 114}

func (i RaftKey) String() string {
	if i < 0 || i >= RaftKey(len(_RaftKey_index)-1) {
		return "RaftKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RaftKey_name[_RaftKey_index[i]:_RaftKey_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SQL_INFO_ROW_COUNT-0]
	_ = x[SQL_INFO_AUTOINCREMENT_IDS-1]
}

const _SqlInfoKey_name = "SQL_INFO_ROW_COUNTSQL_INFO_AUTOINCREMENT_IDS"

var _SqlInfoKey_index = [...]uint8{0, 18, 44}

func (i SqlInfoKey) String() string {
	if i < 0 || i >= SqlInfoKey(len(_SqlInfoKey_index)-1) {
		return "SqlInfoKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SqlInfoKey_name[_SqlInfoKey_index[i]:_SqlInfoKey_index[i+1]]
}
