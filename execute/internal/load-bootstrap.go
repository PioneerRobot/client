package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/plugin"
	"os"
	"path/filepath"
	gp "plugin"
)

// bootstrap.gp 相对于程序根目录的路径
const bootstrapGpFilepath = "bootstrap.gp"

// bootstrap 的初始化函数名
const bootstrapInitFuncName = "Init"

// BootstrapInitFuncType bootstrap 的初始化函数类型
// root:  程序运行路径
// debug: 是否以 Debug 模式启动
type BootstrapInitFuncType = func (root string, debug bool) (core.Global, core.Engine, plugin.Loader)

// 获取指定的 gp 模块
func getGP (root string, path string) *gp.Plugin {
	// 计算完整路径
	relPath := filepath.Join(root, path)
	// 判断指定的文件是否存在
	stat, err := os.Stat(relPath)
	if err != nil {
		if os.IsNotExist(err) {
			panic("指定的 gp 模块不存在：" + relPath)
		}
		panic(err)
	}
	// 判断目标是否是文件
	if stat.IsDir() {
		panic("指定的 gp 模块不是一个文件：" + relPath)
	}
	// 获取模块
	module, err := gp.Open(relPath)
	if err != nil {
		panic(err)
	}
	return module
}

// LoadBootstrap 加载 core.Bootstrap 的构造函数
func LoadBootstrap (root string) BootstrapInitFuncType {
	// 获取 bootstrap 模块
	bootstrap := getGP(root, bootstrapGpFilepath)
	// 获取解析函数
	method, err := bootstrap.Lookup(bootstrapInitFuncName)
	if err != nil {
		panic(err)
	}
	return method.(BootstrapInitFuncType)
}