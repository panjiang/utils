package xtime

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	now := time.Now()
	fmt.Println(GetTS(now))
	fmt.Println(GetDateNumStr(now))
	fmt.Println(TimeFromTS(2147483648))
}
