package infa

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FindProjectRoot 通过向上递归查找 go.mod 找到项目的根目录
func FindProjectRoot() (string, error) {
	// 1. 获取起始点：先尝试获取可执行文件所在目录
	// 注意：这里复用了之前的逻辑，或者直接用 os.Executable
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 处理 go run 生成的临时文件路径
	// 如果是 go run，exePath 会在临时目录，我们改用当前工作目录作为搜索起点
	dir := filepath.Dir(exePath)
	if strings.Contains(dir, "go-build") || strings.Contains(dir, os.TempDir()) {
		dir, _ = os.Getwd()
	}

	// 2. 递归向上查找 go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			// 找到了！这就是项目根目录
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// 已经到达磁盘根目录 (如 C:\) 还是没找到
			// 说明这是生产环境（没有源码），直接返回可执行文件所在目录
			return filepath.Dir(exePath), nil
		}
		dir = parent
	}
}

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
