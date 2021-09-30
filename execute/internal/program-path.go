package internal

import (
	"os"
	"path/filepath"
)

// GetProgramRootPath 获取程序运行路径
func GetProgramRootPath () string {
	// 获取程序运行目录
	root := filepath.Dir(os.Args[0])
	// 计算绝对路径
	abs, err := filepath.Abs(root)
	if err != nil {
		panic(err)
	}
	return abs
}