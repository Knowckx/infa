package infa

import (
	"fmt"
	"testing"
)

func Test_FindProjectRoot(t *testing.T) {
	res, err := FindProjectRoot()
	fmt.Println(res, err)
}
