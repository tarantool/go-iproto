package iproto_test

import (
	"fmt"

	"github.com/tarantool/go-iproto"
)

func Example() {
	fmt.Printf("%s=%d\n", iproto.ER_READONLY, iproto.ER_READONLY)
	fmt.Printf("%s=%d\n", iproto.IPROTO_FEATURE_WATCHERS, iproto.IPROTO_FEATURE_WATCHERS)
	fmt.Printf("%s=%d\n", iproto.IPROTO_FLAG_COMMIT, iproto.IPROTO_FLAG_COMMIT)
	fmt.Printf("%s=%d\n", iproto.IPROTO_SYNC, iproto.IPROTO_SYNC)
	fmt.Printf("%s=%d\n", iproto.IPROTO_SELECT, iproto.IPROTO_SELECT)
	// Output:
	// ER_READONLY=7
	// IPROTO_FEATURE_WATCHERS=3
	// IPROTO_FLAG_COMMIT=1
	// IPROTO_SYNC=1
	// IPROTO_SELECT=1
}
