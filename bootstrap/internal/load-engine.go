package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// engine.gp 相对于程序根目录的路径
const engineGpFilepath = "./kernel/engine.gp"

// engine 的构造函数名
const engineConstructorName = "NewEngine"

// EngineConstructorType engine 的构造函数类型
type EngineConstructorType = func () core.Engine

// LoadEngineModule 加载 core.Engine 的构造函数
func LoadEngineModule (root string) EngineConstructorType {
	// 获取 engine 模块
	gp := common.GetGP(root, engineGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(engineConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(EngineConstructorType)
}