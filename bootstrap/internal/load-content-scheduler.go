package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// contentScheduler.gp 相对于程序根目录的路径
const contentSchedulerGpFilepath = "./kernel/schedulers/content.gp"

// contentScheduler 的构造函数名
const contentSchedulerConstructorName = "NewContentScheduler"

// ContentSchedulerConstructorType contentScheduler 的构造函数类型
type ContentSchedulerConstructorType = func () core.ContentScheduler

// LoadContentSchedulerModule 加载 core.ContentScheduler 的构造函数
func LoadContentSchedulerModule (root string) ContentSchedulerConstructorType {
	// 获取 contentScheduler 模块
	gp := common.GetGP(root, contentSchedulerGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(contentSchedulerConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(ContentSchedulerConstructorType)
}