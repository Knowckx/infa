package web

import (
	"testing"

	infa "github.com/Knowckx/infa"
)

func Test_PrintCookies(t *testing.T) {
	out := GetHostCookies("confluent.cloud")
	infa.PrintToml(out)
}
