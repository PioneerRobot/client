package internal

import (
	"github.com/PioneerRobot/plugin"
	"github.com/PioneerRobot/simple/common"
)

// loader.gp 相对于程序根目录的路径
const loaderGpFilepath = "./kernel/loader.gp"

// loader 的构造函数名
const loaderConstructorName = "NewLoader"

// LoaderConstructorType loader 的构造函数类型
// string  插件存放根目录
type LoaderConstructorType = func (string) plugin.Loader

// LoadLoaderModule 加载 plugin.Loader 的构造函数
func LoadLoaderModule (root string) LoaderConstructorType {
	// 获取 loader 模块
	gp := common.GetGP(root, loaderGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(loaderConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(LoaderConstructorType)
}