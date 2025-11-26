package path

import (
	"os"
	"path/filepath"
	"strings"
)

const InfaProjName = "infa"

// GetAppPath 返回可执行文件所在的目录绝对路径。
// 它通过 os.Executable() 获取程序路径，这是最可靠的方法。
// 为了便于在开发环境中使用 `go run` 或 `go test`，它会检测临时构建路径 (go-build) 或测试执行环境 (.test)，并回退到使用 os.Getwd()。
func GetAppPath() (string, error) {
	// os.Executable() 会返回启动当前进程的可执行文件的绝对路径
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// 在开发时使用 `go run` 或 `go test`，可执行文件会被放在一个临时目录中。
	// 在这些情况下，返回当前工作目录是更合理的行为。
	if strings.Contains(exePath, "go-build") || strings.Contains(exePath, ".test") {
		return os.Getwd()
	}

	// 使用 filepath.Dir 获取可执行文件所在的目录
	return filepath.Dir(exePath), nil
}

// GetProjRootPath 通过在程序路径中搜索项目名称来定位项目的根目录。
// 它现在返回一个错误，以便调用者可以妥善处理路径无法找到的情况。
func GetProjRootPath(projName string) (string, error) {
	appPath, err := GetAppPath()
	if err != nil {
		return "", err
	}

	// 如果 projName 为空，或者路径中不包含 projName，直接返回可执行文件所在目录
	if projName == "" || !strings.Contains(appPath, projName) {
		return appPath, nil
	}

	// 使用 strings.Split 来定位项目根目录。
	// 这会找到第一个出现 projName 的地方，并将其作为根目录。
	paths := strings.Split(appPath, projName)
	rPath := filepath.Join(paths[0], projName)

	return rPath, nil
}

func GetFilePath(root string, mids ...string) string {
	rest := filepath.Join(mids...)
	out := filepath.Join(root, rest)
	return out
}

// LocFilePath 结合 GetProjRootPath 和 GetFilePath，以定位项目内的特定文件或目录。
// 已更新为传播 GetProjRootPath 可能返回的错误。
func LocFilePath(projName string, mids ...string) (string, error) {
	rootPath, err := GetProjRootPath(projName)
	if err != nil {
		return "", err
	}
	rest := filepath.Join(mids...)
	out := filepath.Join(rootPath, rest)
	return out, nil
}

// GetInfaPath 是 LocFilePath 针对 "in-fa" 项目的便捷封装。
// 已更新为传播错误。
func GetInfaPath(mids ...string) (string, error) {
	rootPath, err := GetProjRootPath(InfaProjName)
	if err != nil {
		return "", err
	}
	rest := filepath.Join(mids...)
	out := filepath.Join(rootPath, rest)
	return out, nil
}
