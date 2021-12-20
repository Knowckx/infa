package web

import (
	"testing"

	infa "github.tools.sap/aeolia/in-fa"
)

func Test_PrintCookies(t *testing.T) {
	out := GetHostCookies("confluent.cloud")
	infa.PrintToml(out)
}
