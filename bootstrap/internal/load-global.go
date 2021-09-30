package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// global.gp 相对于程序根目录的路径
const globalGpFilepath = "./kernel/global.gp"

// global 的构造函数名
const globalConstructorName = "NewGlobal"

// GlobalConstructorType global 的构造函数类型
// core.ConfigMap     根配置
// core.ErrorManager  异常管理器
// bool               是否启用 Debug 模式
type GlobalConstructorType = func (core.ConfigMap, core.ErrorManager, bool) core.Global

// LoadGlobalModule 加载 core.Global 的构造函数
func LoadGlobalModule (root string) GlobalConstructorType {
	// 获取 global 模块
	gp := common.GetGP(root, globalGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(globalConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(GlobalConstructorType)
}