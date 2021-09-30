package internal

import (
	"github.com/PioneerRobot/plugin"
	"github.com/PioneerRobot/simple/common"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// parser 模块的存放目录，相对于程序根目录的路径
const parserGpRootPath = "./kernel/parsers"

// ParserFactory 的构造函数名
const parserFactoryConstructorName = "NewParserFactory"

// ParserFactoryConstructorType ParserFactory 的构造函数类型
type ParserFactoryConstructorType = func () plugin.ParserFactory

// GetAllParserModules 获取所有 Parser 模块名
func GetAllParserModules (root string) []string {
	// 计算绝对路径
	path := filepath.Join(root, parserGpRootPath)
	// 读取 parser 模块存放目录下的所有文件
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	// 获取所有的 .gp 文件的文件名
	modules := make([]string, 0)
	for _, file := range files {
		// 获取当前遍历的文件的文件名
		name := file.Name()
		// 排除所有非 .gp 的文件
		if !strings.HasSuffix(name, ".gp") {
			continue
		}
		// 截取文件名
		_, name = filepath.Split(name)
		// 存入模块名数组中
		modules = append(modules, name)
	}
	return modules
}

// LoadParserFactoryModule 加载 plugin.ParserFactory 的构造函数
func LoadParserFactoryModule (root string, name string) ParserFactoryConstructorType {
	// 获取 parser 模块
	gp := common.GetGP(root, filepath.Join(parserGpRootPath, name))
	// 获取构造函数
	method, err := gp.Lookup(parserFactoryConstructorName)
	if err != nil {
		panic(err)
	}
	return method.(ParserFactoryConstructorType)
}