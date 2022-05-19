package infa

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func Test_TimeNow(t *testing.T) {
	out, err := TimeNow()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	assert.Nil(t, err)
	fmt.Println(out)
}
func TimeNow() (int, error) {
	now := time.Now()
	res := now.Format(time.RFC3339)
	fmt.Println(res)
	return 0, nil
}
