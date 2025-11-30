package infa

import (
	"log/slog"
	"os"

	pgk "github.com/Knowckx/infa/pkg"
)

// 对slog的日志格式进行自定义
func InitSimpleLogger() {
	handler := pgk.NewSimpleHandler(os.Stdout)
	slog.SetDefault(slog.New(handler))
}
