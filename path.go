package infa

import (
	"os"
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
