package format

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"log/slog"
	"gopkg.in/yaml.v2"
)

func LoadYamlFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		slog.Error("read file error", "file", f, "error", err)
		return err
	}
	err = yaml.Unmarshal(out, data)
	if err != nil {
		slog.Error("yaml.Unmarshal error", "file", f, "error", err)
		return err
	}
	return err
}

func SaveYamlFile(fileName string, in interface{}) error {
	data, err := yaml.Marshal(in)
	fileName = fileName + ".yaml"
	fMod := os.FileMode(0644)
	err = os.WriteFile(fileName, data, fMod)
	return errors.WithStack(err)
}
