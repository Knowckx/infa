package path

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetProjRootPath(t *testing.T) {
	res := GetProjRootPath("in-fa")
	fmt.Println("result:", res)
	return
}

func Test_GetProjFilePath(t *testing.T) {
	root := GetAppPath()
	res := GetFilePath(root, "config")
	fmt.Println(res)
	return
}

func Test_GetAppPath(t *testing.T) {
	res := GetAppPath()
	fmt.Println(res)
	return
}

func Test_SaveDataToFile(t *testing.T) {
	err := SaveDataToFile("test_file", "data123")
	assert.Nil(t, err)
}
