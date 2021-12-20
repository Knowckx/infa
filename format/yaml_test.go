package format

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SaveJson(t *testing.T) {
	testMap := map[string]string{
		"AA":  "BB",
		"AA1": "BB1",
	}
	err := SaveYamlFile("asu", testMap)
	assert.Nil(t, err)
}
