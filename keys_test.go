// Code generated by generate.sh; DO NOT EDIT.

package iproto_test

import (
	"testing"

	"github.com/tarantool/go-iproto"
)

func TestKey(t *testing.T) {
	cases := []struct {
		Key iproto.Key
		Str string
		Val int
	}{

		{iproto.IPROTO_REQUEST_TYPE, "IPROTO_REQUEST_TYPE", 0x00},
		{iproto.IPROTO_SYNC, "IPROTO_SYNC", 0x01},
		{iproto.IPROTO_REPLICA_ID, "IPROTO_REPLICA_ID", 0x02},
		{iproto.IPROTO_LSN, "IPROTO_LSN", 0x03},
		{iproto.IPROTO_TIMESTAMP, "IPROTO_TIMESTAMP", 0x04},
		{iproto.IPROTO_SCHEMA_VERSION, "IPROTO_SCHEMA_VERSION", 0x05},
		{iproto.IPROTO_SERVER_VERSION, "IPROTO_SERVER_VERSION", 0x06},
		{iproto.IPROTO_GROUP_ID, "IPROTO_GROUP_ID", 0x07},
		{iproto.IPROTO_TSN, "IPROTO_TSN", 0x08},
		{iproto.IPROTO_FLAGS, "IPROTO_FLAGS", 0x09},
		{iproto.IPROTO_STREAM_ID, "IPROTO_STREAM_ID", 0x0a},
		{iproto.IPROTO_SPACE_ID, "IPROTO_SPACE_ID", 0x10},
		{iproto.IPROTO_INDEX_ID, "IPROTO_INDEX_ID", 0x11},
		{iproto.IPROTO_LIMIT, "IPROTO_LIMIT", 0x12},
		{iproto.IPROTO_OFFSET, "IPROTO_OFFSET", 0x13},
		{iproto.IPROTO_ITERATOR, "IPROTO_ITERATOR", 0x14},
		{iproto.IPROTO_INDEX_BASE, "IPROTO_INDEX_BASE", 0x15},
		{iproto.IPROTO_FETCH_POSITION, "IPROTO_FETCH_POSITION", 0x1f},
		{iproto.IPROTO_KEY, "IPROTO_KEY", 0x20},
		{iproto.IPROTO_TUPLE, "IPROTO_TUPLE", 0x21},
		{iproto.IPROTO_FUNCTION_NAME, "IPROTO_FUNCTION_NAME", 0x22},
		{iproto.IPROTO_USER_NAME, "IPROTO_USER_NAME", 0x23},
		{iproto.IPROTO_INSTANCE_UUID, "IPROTO_INSTANCE_UUID", 0x24},
		{iproto.IPROTO_REPLICASET_UUID, "IPROTO_REPLICASET_UUID", 0x25},
		{iproto.IPROTO_VCLOCK, "IPROTO_VCLOCK", 0x26},
		{iproto.IPROTO_EXPR, "IPROTO_EXPR", 0x27},
		{iproto.IPROTO_OPS, "IPROTO_OPS", 0x28},
		{iproto.IPROTO_BALLOT, "IPROTO_BALLOT", 0x29},
		{iproto.IPROTO_TUPLE_META, "IPROTO_TUPLE_META", 0x2a},
		{iproto.IPROTO_OPTIONS, "IPROTO_OPTIONS", 0x2b},
		{iproto.IPROTO_OLD_TUPLE, "IPROTO_OLD_TUPLE", 0x2c},
		{iproto.IPROTO_NEW_TUPLE, "IPROTO_NEW_TUPLE", 0x2d},
		{iproto.IPROTO_AFTER_POSITION, "IPROTO_AFTER_POSITION", 0x2e},
		{iproto.IPROTO_AFTER_TUPLE, "IPROTO_AFTER_TUPLE", 0x2f},
		{iproto.IPROTO_DATA, "IPROTO_DATA", 0x30},
		{iproto.IPROTO_ERROR_24, "IPROTO_ERROR_24", 0x31},
		{iproto.IPROTO_METADATA, "IPROTO_METADATA", 0x32},
		{iproto.IPROTO_BIND_METADATA, "IPROTO_BIND_METADATA", 0x33},
		{iproto.IPROTO_BIND_COUNT, "IPROTO_BIND_COUNT", 0x34},
		{iproto.IPROTO_POSITION, "IPROTO_POSITION", 0x35},
		{iproto.IPROTO_SQL_TEXT, "IPROTO_SQL_TEXT", 0x40},
		{iproto.IPROTO_SQL_BIND, "IPROTO_SQL_BIND", 0x41},
		{iproto.IPROTO_SQL_INFO, "IPROTO_SQL_INFO", 0x42},
		{iproto.IPROTO_STMT_ID, "IPROTO_STMT_ID", 0x43},
		{iproto.IPROTO_REPLICA_ANON, "IPROTO_REPLICA_ANON", 0x50},
		{iproto.IPROTO_ID_FILTER, "IPROTO_ID_FILTER", 0x51},
		{iproto.IPROTO_ERROR, "IPROTO_ERROR", 0x52},
		{iproto.IPROTO_TERM, "IPROTO_TERM", 0x53},
		{iproto.IPROTO_VERSION, "IPROTO_VERSION", 0x54},
		{iproto.IPROTO_FEATURES, "IPROTO_FEATURES", 0x55},
		{iproto.IPROTO_TIMEOUT, "IPROTO_TIMEOUT", 0x56},
		{iproto.IPROTO_EVENT_KEY, "IPROTO_EVENT_KEY", 0x57},
		{iproto.IPROTO_EVENT_DATA, "IPROTO_EVENT_DATA", 0x58},
		{iproto.IPROTO_TXN_ISOLATION, "IPROTO_TXN_ISOLATION", 0x59},
		{iproto.IPROTO_VCLOCK_SYNC, "IPROTO_VCLOCK_SYNC", 0x5a},
		{iproto.IPROTO_AUTH_TYPE, "IPROTO_AUTH_TYPE", 0x5b},
	}

	for _, tc := range cases {
		t.Run(tc.Str, func(t *testing.T) {
			if tc.Key.String() != tc.Str {
				t.Errorf("Got %s, expected %s", tc.Key.String(), tc.Str)
			}
			if int(tc.Key) != tc.Val {
				t.Errorf("Got %d, expected %d", tc.Key, tc.Val)
			}
		})
	}
}

func TestMetadataKey(t *testing.T) {
	cases := []struct {
		MetadataKey iproto.MetadataKey
		Str         string
		Val         int
	}{

		{iproto.IPROTO_FIELD_NAME, "IPROTO_FIELD_NAME", 0},
		{iproto.IPROTO_FIELD_TYPE, "IPROTO_FIELD_TYPE", 1},
		{iproto.IPROTO_FIELD_COLL, "IPROTO_FIELD_COLL", 2},
		{iproto.IPROTO_FIELD_IS_NULLABLE, "IPROTO_FIELD_IS_NULLABLE", 3},
		{iproto.IPROTO_FIELD_IS_AUTOINCREMENT, "IPROTO_FIELD_IS_AUTOINCREMENT", 4},
		{iproto.IPROTO_FIELD_SPAN, "IPROTO_FIELD_SPAN", 5},
	}

	for _, tc := range cases {
		t.Run(tc.Str, func(t *testing.T) {
			if tc.MetadataKey.String() != tc.Str {
				t.Errorf("Got %s, expected %s", tc.MetadataKey.String(), tc.Str)
			}
			if int(tc.MetadataKey) != tc.Val {
				t.Errorf("Got %d, expected %d", tc.MetadataKey, tc.Val)
			}
		})
	}
}

func TestBallotKey(t *testing.T) {
	cases := []struct {
		BallotKey iproto.BallotKey
		Str       string
		Val       int
	}{

		{iproto.IPROTO_BALLOT_IS_RO_CFG, "IPROTO_BALLOT_IS_RO_CFG", 0x01},
		{iproto.IPROTO_BALLOT_VCLOCK, "IPROTO_BALLOT_VCLOCK", 0x02},
		{iproto.IPROTO_BALLOT_GC_VCLOCK, "IPROTO_BALLOT_GC_VCLOCK", 0x03},
		{iproto.IPROTO_BALLOT_IS_RO, "IPROTO_BALLOT_IS_RO", 0x04},
		{iproto.IPROTO_BALLOT_IS_ANON, "IPROTO_BALLOT_IS_ANON", 0x05},
		{iproto.IPROTO_BALLOT_IS_BOOTED, "IPROTO_BALLOT_IS_BOOTED", 0x06},
		{iproto.IPROTO_BALLOT_CAN_LEAD, "IPROTO_BALLOT_CAN_LEAD", 0x07},
		{iproto.IPROTO_BALLOT_BOOTSTRAP_LEADER_UUID, "IPROTO_BALLOT_BOOTSTRAP_LEADER_UUID", 0x08},
		{iproto.IPROTO_BALLOT_REGISTERED_REPLICA_UUIDS, "IPROTO_BALLOT_REGISTERED_REPLICA_UUIDS", 0x09},
	}

	for _, tc := range cases {
		t.Run(tc.Str, func(t *testing.T) {
			if tc.BallotKey.String() != tc.Str {
				t.Errorf("Got %s, expected %s", tc.BallotKey.String(), tc.Str)
			}
			if int(tc.BallotKey) != tc.Val {
				t.Errorf("Got %d, expected %d", tc.BallotKey, tc.Val)
			}
		})
	}
}

func TestRaftKey(t *testing.T) {
	cases := []struct {
		RaftKey iproto.RaftKey
		Str     string
		Val     int
	}{

		{iproto.IPROTO_RAFT_TERM, "IPROTO_RAFT_TERM", 0},
		{iproto.IPROTO_RAFT_VOTE, "IPROTO_RAFT_VOTE", 1},
		{iproto.IPROTO_RAFT_STATE, "IPROTO_RAFT_STATE", 2},
		{iproto.IPROTO_RAFT_VCLOCK, "IPROTO_RAFT_VCLOCK", 3},
		{iproto.IPROTO_RAFT_LEADER_ID, "IPROTO_RAFT_LEADER_ID", 4},
		{iproto.IPROTO_RAFT_IS_LEADER_SEEN, "IPROTO_RAFT_IS_LEADER_SEEN", 5},
	}

	for _, tc := range cases {
		t.Run(tc.Str, func(t *testing.T) {
			if tc.RaftKey.String() != tc.Str {
				t.Errorf("Got %s, expected %s", tc.RaftKey.String(), tc.Str)
			}
			if int(tc.RaftKey) != tc.Val {
				t.Errorf("Got %d, expected %d", tc.RaftKey, tc.Val)
			}
		})
	}
}
