package infa

import (
	"github.com/Knowckx/infa/format"
	"github.com/Knowckx/infa/parallel"
	"github.com/Knowckx/infa/path"
	"github.com/Knowckx/infa/scatter"
	"github.com/Knowckx/infa/util"
)

func ShortStr(in string) string {
	return util.ShortStr(in)
}

func SaveDataToFile(fileName, data string) error {
	return path.SaveDataToFile(fileName, data)
}

func LoadJsonFile(f string, data interface{}) error {
	return format.LoadJsonFile(f, data)
}

func ReadFile(f string) (string, error) {
	return path.ReadFile(f)
}

func ExecCmd(in string) error {
	return scatter.ExecCmd(in)
}

func LocFilePath(projName string, mids ...string) string {
	return path.LocFilePath(projName, mids...)
}

func NewParallel(max int) *parallel.WaitPool {
	return parallel.NewWaitPool(max)
}
