package app

import (
	"encoding/json"
	"fmt"

	"github.com/MortalSC/FastGO/cmd/fg-apiserver/app/options"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// NewFastG0Command 创建一个*cobra.Command 对象,用于启动应用程序
func NewFastG0Command() *cobra.Command {

	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		//指定命令的名字,该名字会出现在帮助信息中
		Use: "fg-apiserver",
		//命令的简短描述
		Short: "A very lightweight ful1 go project",
		Long:  "A very lightweight full go project,designed to help beginners quickly learn Go project development.",
		//命令出错时,不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		//指定调用cmd.Execute()时,执行的Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello FastGo!")

			if err := viper.Unmarshal(opts); err != nil {
				return err
			}

			if err := opts.Validate(); err != nil {
				return err
			}

			fmt.Printf("Read MySQL host from Viper: %s\n\n", viper.GetString("mysql.host"))

			jsonData, _ := json.MarshalIndent(opts, "", "  ")
			fmt.Printf("MySQL options: %s\n", string(jsonData))

			return nil
		},
		//设置命令运行时的参数检查,不需要指定命令行参数。例如:./fg-apiserver paraml param2
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver config file")

	return cmd
}
