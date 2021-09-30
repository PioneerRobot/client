package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// cidGenerator.gp 相对于程序根目录的路径
const generatorGpFilepath = "./kernel/cid-generator.gp"

// cidGenerator 的构造函数名
const generatorConstructorName = "NewCidGenerator"

// CidGeneratorConstructorType cidGenerator 的构造函数类型
type CidGeneratorConstructorType = func () core.CidGenerator

// LoadCidGeneratorModule 加载 core.CidGenerator 的构造函数
func LoadCidGeneratorModule (root string) CidGeneratorConstructorType {
	// 获取 cidGenerator 模块
	gp := common.GetGP(root, generatorGpFilepath)
	// 获取构造函数
	method, err := gp.Lookup(generatorConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(CidGeneratorConstructorType)
}