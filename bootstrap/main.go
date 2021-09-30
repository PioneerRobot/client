package main

import (
	"github.com/PioneerRobot/client/bootstrap/internal"
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/plugin"
	"path/filepath"
)

// 插件存放目录名
const pluginDirName = "plugins"

// 初始化全局对象
func initGlobal (root string, debug bool) core.Global {
	// 加载「配置」模块
	configModule := internal.LoadConfigModule(root)
	// TODO: 解析配置文件
	configRoot := configModule(nil)
	// 加载「日志」模块
	// loggerModule := internal.LoadLoggerModule(root)
	// 构建日志处理器实例
	// logger := loggerModule()
	// 加载「错误管理器」模块
	errorManagerModule := internal.LoadErrorManagerModule(root)
	// 构建错误管理器实例
	errorManager := errorManagerModule()
	// 加载「全局对象」模块
	globalModule := internal.LoadGlobalModule(root)
	// 构建全局对象实例
	return globalModule(configRoot, errorManager, debug)
}

// 初始化插件加载器
func initPluginLoader (root string, pluginPath string) plugin.Loader {
	// 加载「插件加载器」模块
	loaderModule := internal.LoadLoaderModule(root)
	// 构建插件加载器实例
	loader := loaderModule(pluginPath)
	// 加载所有插件包解析器工厂
	for _, parserName := range internal.GetAllParserModules(root) {
		// 获取当前循环的「插件包解析器工厂」模块
		parserFactoryModule := internal.LoadParserFactoryModule(root, parserName)
		// 构造插件包解析器工厂实例
		parserFactory := parserFactoryModule()
		// 将插件包解析器工厂注册到插件加载器中
		loader.RegisterReaderFactory(parserFactory)
	}
	// 返回插件加载器
	return loader
}

// 加载所有插件，并注入组件池中
func injectAllPlugins(loader plugin.Loader, poolPointer *core.ComponentPool) {
	components, err := loader.GetAllComponents()
	if err != nil {
		panic(err)
	}
	*poolPointer = append(*poolPointer, *components...)
}

// 加载所有调度器，并注入组件池中
func injectAllSchedulers(root string, poolPointer *core.ComponentPool) {
	// 加载「内容调度器」模块
	contentSchedulerModule := internal.LoadContentSchedulerModule(root)
	// 构造内容调度器实例
	contentScheduler := contentSchedulerModule()

	// 加载「处理调度器」模块
	handleSchedulerModule := internal.LoadHandleSchedulerModule(root)
	// 构造处理调度器实例
	handleScheduler := handleSchedulerModule()

	// 加载「内容调度器」模块
	pushSchedulerModule := internal.LoadPushSchedulerModule(root)
	// 构造内容调度器实例
	pushScheduler := pushSchedulerModule()

	// 加载「内容调度器」模块
	transformSchedulerModule := internal.LoadTransformSchedulerModule(root)
	// 构造内容调度器实例
	transformScheduler := transformSchedulerModule()

	// 将调度器存入组件池中
	*poolPointer = append(*poolPointer, contentScheduler, handleScheduler, pushScheduler, transformScheduler)
}

// 初始化引擎
func initEngine (root string, global core.Global, pool core.ComponentPool) core.Engine {
	// 加载「CID 生成器」模块
	cidGeneratorModule := internal.LoadCidGeneratorModule(root)
	// 构建 CID 生成器实例
	cidGenerator := cidGeneratorModule()
	// 加载「引擎」模块
	engineModule := internal.LoadEngineModule(root)
	// 构建引擎实例
	engine := engineModule()
	// 初始化引擎
	engine.Init(global, cidGenerator, pool)
	// 返回引擎实例
	return engine
}

// Init 此函数需要完成如下操作：
// 1. 加载各基本组件
// 2. 构建并初始化引擎
func Init(root string, debug bool) (core.Global, core.Engine, plugin.Loader) {
	// 初始化「全局对象」
	global := initGlobal(root, debug)
	// 构造组件池
	pool := make(core.ComponentPool, 0)
	// 加载并注入所有调度器
	injectAllSchedulers(root, &pool)
	// 计算完整的插件存放路径
	pluginSavePath := filepath.Join(root, pluginDirName)
	// 初始化「插件加载器」
	loader := initPluginLoader(root, pluginSavePath)
	// 注入所有插件
	injectAllPlugins(loader, &pool)
	// 初始化引擎
	engine := initEngine(root, global, pool)
	// 返回全局对象、引擎、插件加载器
	return global, engine, loader
}