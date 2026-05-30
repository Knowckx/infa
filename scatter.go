package infa

import (
	"os"
	"time"

	"log/slog"
)



func Getenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		slog.Warn("Get environment variable failed. return nil.", "key", key)
	}
	return value
}

func FormatTime(in time.Time) string {
	return in.Format(time.RFC3339)
}

// server time use zone Shanghai.
func SHTime(in time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	return in.In(loc)
}
