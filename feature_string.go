// Code generated by "stringer -type=Feature"; DO NOT EDIT.

package iproto

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IPROTO_FEATURE_STREAMS-0]
	_ = x[IPROTO_FEATURE_TRANSACTIONS-1]
	_ = x[IPROTO_FEATURE_ERROR_EXTENSION-2]
	_ = x[IPROTO_FEATURE_WATCHERS-3]
	_ = x[IPROTO_FEATURE_PAGINATION-4]
	_ = x[IPROTO_FEATURE_SPACE_AND_INDEX_NAMES-5]
	_ = x[IPROTO_FEATURE_WATCH_ONCE-6]
	_ = x[IPROTO_FEATURE_DML_TUPLE_EXTENSION-7]
	_ = x[IPROTO_FEATURE_CALL_RET_TUPLE_EXTENSION-8]
	_ = x[IPROTO_FEATURE_CALL_ARG_TUPLE_EXTENSION-9]
	_ = x[IPROTO_FEATURE_FETCH_SNAPSHOT_CURSOR-10]
	_ = x[IPROTO_FEATURE_IS_SYNC-11]
	_ = x[IPROTO_FEATURE_INSERT_ARROW-12]
}

const _Feature_name = "IPROTO_FEATURE_STREAMSIPROTO_FEATURE_TRANSACTIONSIPROTO_FEATURE_ERROR_EXTENSIONIPROTO_FEATURE_WATCHERSIPROTO_FEATURE_PAGINATIONIPROTO_FEATURE_SPACE_AND_INDEX_NAMESIPROTO_FEATURE_WATCH_ONCEIPROTO_FEATURE_DML_TUPLE_EXTENSIONIPROTO_FEATURE_CALL_RET_TUPLE_EXTENSIONIPROTO_FEATURE_CALL_ARG_TUPLE_EXTENSIONIPROTO_FEATURE_FETCH_SNAPSHOT_CURSORIPROTO_FEATURE_IS_SYNCIPROTO_FEATURE_INSERT_ARROW"

var _Feature_index = [...]uint16{0, 22, 49, 79, 102, 127, 163, 188, 222, 261, 300, 336, 358, 385}

func (i Feature) String() string {
	if i < 0 || i >= Feature(len(_Feature_index)-1) {
		return "Feature(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Feature_name[_Feature_index[i]:_Feature_index[i+1]]
}
