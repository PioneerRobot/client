package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// handleScheduler.gp 相对于程序根目录的路径
const handleSchedulerGpFilepath = "./kernel/schedulers/handle.gp"

// handleScheduler 的构造函数名
const handleSchedulerConstructorName = "NewHandleScheduler"

// HandleSchedulerConstructorType handleScheduler 的构造函数类型
type HandleSchedulerConstructorType = func () core.HandleScheduler

// LoadHandleSchedulerModule 加载 core.HandleScheduler 的构造函数
func LoadHandleSchedulerModule (root string) HandleSchedulerConstructorType {
	// 获取 handleScheduler 模块
	gp := common.GetGP(root, handleSchedulerGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(handleSchedulerConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(HandleSchedulerConstructorType)
}