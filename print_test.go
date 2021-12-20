package infa

import (
	"testing"
	"time"
)

func Test_Printf(t *testing.T) {
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
}

func Test_RecordExecTime(t *testing.T) {
	defer RecordExecTime("Test_RecordExecTime")()
	time.Sleep(2 * time.Second)
}
