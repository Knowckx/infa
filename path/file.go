package path

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"log/slog"
)

func SaveDataToFile(fileName, data string) error {
	path := fileName // "./" +
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.WriteString(data)
	slog.Info("write data to file success.", "fileName", fileName)
	return nil
}

func ReadFile(f string) (string, error) {
	res, err := ioutil.ReadFile(f)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(res), nil
}
