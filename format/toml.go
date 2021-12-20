package format

import (
	"io/ioutil"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog/log"
)

func LoadTomlFile(f string, data interface{}) error {
	out, err := ioutil.ReadFile(f)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("read file error")
		return err
	}

	err = toml.Unmarshal(out, data)
	if err != nil {
		log.Error().Str("file", f).Err(err).Msg("yaml.Unmarshal error")
		return err
	}
	return err
}
