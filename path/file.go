package path

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func SaveDataToFile(fileName, data string) error {
	path := "./" + fileName
	f, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	f.WriteString(data)
	log.Info().Str("fileName", fileName).Msg("write data to file success.")
	return nil
}
