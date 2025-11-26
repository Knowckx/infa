package format

import (
	"encoding/json"
	"io/ioutil"

	"log/slog"
)

// save json | load json
func LoadJsonFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		slog.Error("read file error", "file", f, "error", err)
		return err
	}
	err = json.Unmarshal(out, data)
	if err != nil {
		slog.Error("json.Unmarshal error", "file", f, "error", err)
		return err
	}
	return err
}
