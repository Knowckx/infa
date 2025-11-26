package format

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
	"log/slog"
)

func LoadTomlFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		slog.Error("read file error", "file", f, "error", err)
		return err
	}

	err = toml.Unmarshal(out, data)
	if err != nil {
		slog.Error("toml.Unmarshal error", "file", f, "error", err)
		return err
	}
	return err
}
