package utils

import (
	"path/filepath"
	"runtime"
)

// 获取基于当前源文件的配置文件路径
func GetDefaultConfigPath(configPath string) string {
	// 获取当前源文件的绝对路径
	_, filename, _, _ := runtime.Caller(1)
	// 获取当前源文件所在目录
	dir := filepath.Dir(filename)
	// 使用传入的相对路径构建完整路径
	return filepath.Join(dir, configPath)
}
