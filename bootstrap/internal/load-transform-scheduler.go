package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// transformScheduler.gp 相对于程序根目录的路径
const transformSchedulerGpFilepath = "./kernel/schedulers/transform.gp"

// transformScheduler 的构造函数名
const transformSchedulerConstructorName = "NewTransformScheduler"

// TransformSchedulerConstructorType transformScheduler 的构造函数类型
type TransformSchedulerConstructorType = func () core.TransformScheduler

// LoadTransformSchedulerModule 加载 core.TransformScheduler 的构造函数
func LoadTransformSchedulerModule (root string) TransformSchedulerConstructorType {
	// 获取 transformScheduler 模块
	gp := common.GetGP(root, transformSchedulerGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(transformSchedulerConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(TransformSchedulerConstructorType)
}