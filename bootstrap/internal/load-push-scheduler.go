package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// pushScheduler.gp 相对于程序根目录的路径
const pushSchedulerGpFilepath = "./kernel/schedulers/push.gp"

// pushScheduler 的构造函数名
const pushSchedulerConstructorName = "NewPushScheduler"

// PushSchedulerConstructorType pushScheduler 的构造函数类型
type PushSchedulerConstructorType = func () core.PushScheduler

// LoadPushSchedulerModule 加载 core.PushScheduler 的构造函数
func LoadPushSchedulerModule (root string) PushSchedulerConstructorType {
	// 获取 pushScheduler 模块
	gp := common.GetGP(root, pushSchedulerGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(pushSchedulerConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(PushSchedulerConstructorType)
}