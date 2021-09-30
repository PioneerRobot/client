package internal

import (
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/simple/common"
)

// config.gp 相对于程序根目录的路径
const configGpFilepath = "./kernel/config.gp"

// config 的解析函数名
const configParseFuncName = "ParseConfig"

// ConfigParseFuncType config 的解析函数类型
// []byte: 配置内容
type ConfigParseFuncType = func ([]byte) core.ConfigMap

// LoadConfigModule 加载 core.Config 的构造函数
func LoadConfigModule (root string) ConfigParseFuncType {
	// 获取 config 模块
	gp := common.GetGP(root, configGpFilepath)
	// 获取解析函数
	method, err := gp.Lookup(configParseFuncName)
	if err != nil {
		panic(err)
	}
	return method.(ConfigParseFuncType)
}