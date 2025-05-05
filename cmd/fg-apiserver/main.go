package main

import (
	"os"
	"github.com/MortalSC/FastGO/cmd/fg-apiserver/app"

	_ "go.uber.org/automaxprocs" // 自动设置最大可用的 CPU 核心数
)

func main() {
	// 创建 Go 项目
	command := app.NewFastG0Command()

	// 执行命令并处理错误
	if err := command.Execute(); err != nil {
		// 如果发生错误，则退出程序
		// 返回退出码，可以使其他程序根据退出码来判断服务运行状态
		os.Exit(1)
	}
}
