package main

import (
	"fmt"
	"github.com/PioneerRobot/client/execute/internal"
	"github.com/PioneerRobot/core"
	"github.com/PioneerRobot/plugin"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Debug 调试模式
const Debug = false

// 程序停止信号
var signalChan = make(chan os.Signal, 1)

// ProgramRootPath 程序运行路径
var ProgramRootPath = internal.GetProgramRootPath()

// 检查异常
func checkError (err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}

// 运行引擎
func runEngine(_ core.Global, engine core.Engine, _ plugin.Loader) {
	fmt.Printf("初始状态：%s\n", engine.State())

	fmt.Printf("准备引擎...\n")
	// 准备引擎
	err := engine.Prepare()
	checkError(err)
	time.Sleep(100 * time.Microsecond)
	fmt.Printf("状态：%s\n", engine.State())

	// 运行引擎
	fmt.Printf("运行引擎...\n")
	engine.Run()
	time.Sleep(100 * time.Microsecond)
	fmt.Printf("状态：%s\n", engine.State())
}

func destroy (engine core.Engine, loader plugin.Loader) {
	// 停止引擎
	fmt.Printf("停止引擎...\n")
	engine.Stop()
	time.Sleep(3 * time.Second)
	fmt.Printf("状态：%s\n", engine.State())

	// 销毁引擎
	fmt.Printf("销毁引擎...\n")
	engine.Destroy()
	time.Sleep(3 * time.Second)
	fmt.Printf("状态：%s\n", engine.State())

	// 销毁插件加载器
	err := loader.Destroy()
	checkError(err)
}

func main () {
	// 加载 bootstrap.gp
	bootstrap := internal.LoadBootstrap(ProgramRootPath)
	// 执行初始化
	global, engine, loader := bootstrap(ProgramRootPath, Debug)

	// 接收系统信号
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// 运行引擎
	runEngine(global, engine, loader)

	// 阻塞后续执行，直到收到退出程序的指令
	<- signalChan
	fmt.Println("正在停止引擎，请耐心等待...")

	// 销毁程序
	destroy(engine, loader)

	// 退出程序
	os.Exit(0)
}