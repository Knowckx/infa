package infa

import (
	"github.com/Knowckx/infa/format"
	"github.com/Knowckx/infa/parallel"
	"github.com/Knowckx/infa/path"
	"github.com/Knowckx/infa/scatter"
	"github.com/Knowckx/infa/util"
	"k8s.io/apimachinery/pkg/util/rand"
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



func NewParallel(max int) *parallel.WaitPool {
	return parallel.NewWaitPool(max)
}

// 泛型的三元表达式函数
func Return[T any](boolExpression bool, trueReturnValue, falseReturnValue T) T {
	if boolExpression {
		return trueReturnValue
	} else {
		return falseReturnValue
	}
}

// 泛型的数组随机数
func GetRandomElement[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}
