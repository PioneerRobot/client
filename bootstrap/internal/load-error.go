package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// error.gp 相对于程序根目录的路径
const errorGpFilepath = "./kernel/error.gp"

// error 的构造函数名
const errorConstructorName = "NewErrorManager"

// ErrorManagerConstructorType error 的构造函数类型
type ErrorManagerConstructorType = func () core.ErrorManager

// LoadErrorManagerModule 加载 core.ErrorManager 的构造函数
func LoadErrorManagerModule (root string) ErrorManagerConstructorType {
	// 获取 error 模块
	gp := common.GetGP(root, errorGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(errorConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(ErrorManagerConstructorType)
}