package infa

import (
	"os"

	"github.com/rs/zerolog/log"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Warn().Str("key", key).Msg("Get environment variable failed. return nil.")
	}
	return value
}
