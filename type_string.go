// Code generated by "stringer -type=Type"; DO NOT EDIT.

package iproto

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_OK-0]
	_ = x[IPROTO_SELECT-1]
	_ = x[IPROTO_INSERT-2]
	_ = x[IPROTO_REPLACE-3]
	_ = x[IPROTO_UPDATE-4]
	_ = x[IPROTO_DELETE-5]
	_ = x[IPROTO_CALL_16-6]
	_ = x[IPROTO_AUTH-7]
	_ = x[IPROTO_EVAL-8]
	_ = x[IPROTO_UPSERT-9]
	_ = x[IPROTO_CALL-10]
	_ = x[IPROTO_EXECUTE-11]
	_ = x[IPROTO_NOP-12]
	_ = x[IPROTO_PREPARE-13]
	_ = x[IPROTO_BEGIN-14]
	_ = x[IPROTO_COMMIT-15]
	_ = x[IPROTO_ROLLBACK-16]
	_ = x[IPROTO_RAFT-30]
	_ = x[IPROTO_RAFT_PROMOTE-31]
	_ = x[IPROTO_RAFT_DEMOTE-32]
	_ = x[IPROTO_RAFT_CONFIRM-40]
	_ = x[IPROTO_RAFT_ROLLBACK-41]
	_ = x[IPROTO_PING-64]
	_ = x[IPROTO_JOIN-65]
	_ = x[IPROTO_SUBSCRIBE-66]
	_ = x[IPROTO_VOTE_DEPRECATED-67]
	_ = x[IPROTO_VOTE-68]
	_ = x[IPROTO_FETCH_SNAPSHOT-69]
	_ = x[IPROTO_REGISTER-70]
	_ = x[IPROTO_JOIN_META-71]
	_ = x[IPROTO_JOIN_SNAPSHOT-72]
	_ = x[IPROTO_ID-73]
	_ = x[IPROTO_WATCH-74]
	_ = x[IPROTO_UNWATCH-75]
	_ = x[IPROTO_EVENT-76]
	_ = x[VY_INDEX_RUN_INFO-100]
	_ = x[VY_INDEX_PAGE_INFO-101]
	_ = x[VY_RUN_ROW_INDEX-102]
	_ = x[IPROTO_CHUNK-128]
	_ = x[IPROTO_TYPE_ERROR-32768]
	_ = x[IPROTO_UNKNOWN - -1]
}

const (
	_Type_name_0 = "IPROTO_UNKNOWNIPROTO_OKIPROTO_SELECTIPROTO_INSERTIPROTO_REPLACEIPROTO_UPDATEIPROTO_DELETEIPROTO_CALL_16IPROTO_AUTHIPROTO_EVALIPROTO_UPSERTIPROTO_CALLIPROTO_EXECUTEIPROTO_NOPIPROTO_PREPAREIPROTO_BEGINIPROTO_COMMITIPROTO_ROLLBACK"
	_Type_name_1 = "IPROTO_RAFTIPROTO_RAFT_PROMOTEIPROTO_RAFT_DEMOTE"
	_Type_name_2 = "IPROTO_RAFT_CONFIRMIPROTO_RAFT_ROLLBACK"
	_Type_name_3 = "IPROTO_PINGIPROTO_JOINIPROTO_SUBSCRIBEIPROTO_VOTE_DEPRECATEDIPROTO_VOTEIPROTO_FETCH_SNAPSHOTIPROTO_REGISTERIPROTO_JOIN_METAIPROTO_JOIN_SNAPSHOTIPROTO_IDIPROTO_WATCHIPROTO_UNWATCHIPROTO_EVENT"
	_Type_name_4 = "VY_INDEX_RUN_INFOVY_INDEX_PAGE_INFOVY_RUN_ROW_INDEX"
	_Type_name_5 = "IPROTO_CHUNK"
	_Type_name_6 = "IPROTO_TYPE_ERROR"
)

var (
	_Type_index_0 = [...]uint8{0, 14, 23, 36, 49, 63, 76, 89, 103, 114, 125, 138, 149, 163, 173, 187, 199, 212, 227}
	_Type_index_1 = [...]uint8{0, 11, 30, 48}
	_Type_index_2 = [...]uint8{0, 19, 39}
	_Type_index_3 = [...]uint8{0, 11, 22, 38, 60, 71, 92, 107, 123, 143, 152, 164, 178, 190}
	_Type_index_4 = [...]uint8{0, 17, 35, 51}
)

func (i Type) String() string {
	switch {
	case -1 <= i && i <= 16:
		i -= -1
		return _Type_name_0[_Type_index_0[i]:_Type_index_0[i+1]]
	case 30 <= i && i <= 32:
		i -= 30
		return _Type_name_1[_Type_index_1[i]:_Type_index_1[i+1]]
	case 40 <= i && i <= 41:
		i -= 40
		return _Type_name_2[_Type_index_2[i]:_Type_index_2[i+1]]
	case 64 <= i && i <= 76:
		i -= 64
		return _Type_name_3[_Type_index_3[i]:_Type_index_3[i+1]]
	case 100 <= i && i <= 102:
		i -= 100
		return _Type_name_4[_Type_index_4[i]:_Type_index_4[i+1]]
	case i == 128:
		return _Type_name_5
	case i == 32768:
		return _Type_name_6
	default:
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}