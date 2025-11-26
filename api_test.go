package infa

import (
	"fmt"
	"testing"
)



func Test_QuickTest(t *testing.T) {
	avger := Averager{}
	avger.AddNumber(1)
	avger.AddNumber(2.1)
	avger.AddNumber(5)
	fmt.Println(avger.Avg)
}
