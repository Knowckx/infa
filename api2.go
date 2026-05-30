package infa

import (
	"log/slog"
	"os/exec"
)

// OpenFileBySystem 调用系统默认程序打开文件 (相当于双击)
func OpenFileBySystem(filePath string) {
	// cmd /c start "" "你的文件路径"
	// 第一个 "" 是必须的，它是 start 命令的 Title 参数，防止路径被误认为标题
	cmd := exec.Command("cmd", "/c", "start", "", filePath)
	err := cmd.Start() // Start 是异步的，不会阻塞你的程序
	if err != nil {
		slog.Error("打开音乐文件失败", "path", filePath, "err", err)
	} else {
		slog.Info("🎵 已唤起播放器", "file", filePath)
	}
}
