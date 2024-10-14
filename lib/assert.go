package lib

import "fmt"

func Assert(condition bool, msg string) {
	if !condition {
		panic(fmt.Sprintf("ASSERT_ERR: %s", msg))
	}
}
