/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	vipercmd "github.com/DesistDaydream/go-cobra/cmd/viper"
	"github.com/DesistDaydream/go-cobra/config"
	"github.com/spf13/cobra"
)

type RootFlags struct {
	// 这里定义的变量，可以在下面的 init 函数中，通过 rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "指定配置文件(默认在$HOME/.cobracli.yaml)") 进行绑定
	// 也可以通过 viper 进行绑定
	CfgFile string
}

var rootFlags RootFlags

// Execute 将所有子命令添加到根命令并设置 Flags。这由 main.main() 调用。它只需要对 rootCmd 发生一次。
func Execute() {
	app := newApp()
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func newApp() *cobra.Command {
	// rootCmd 表示在没有任何子命令调用的情况时的基本命令。
	var rootCmd = &cobra.Command{
		Use:   "go-cobra",
		Short: "这个应用简要的描述",
		Long: `横跨多行的较长描述，可能包含示例和使用应用程序的用法。 例如：
当我运行程序时，会显示该描述内容
	如果使用缩进，这行在界面展示时有缩进。`,
		// 如果这个应用没有任何子命令，直接使用 go-cobra 执行的话，将会执行下面 Run 字段指定的函数
		Run: rootRun,
	}

	// 我们可以在这里定义命令行 Flags 和 配置设置。
	// 这里可以做一些初始化的工作，比如初始化数据库连接、初始化日志、读取配置文件等等

	// ######## 添加 命令行Flags ########
	// Cobra 支持 持久性flags (i.e. Global Flags)，如果在这个位置定义，则这些 flags 对应用程序来说是全局的。
	// 第一个参数是变量，用于存储该flag的值；第二个参数为该flag的名字；第三个参数为该flag的默认值,无默认值可以为空；第四个参数是该flag的描述信息
	// 比如我现在使用如下命令: go-cobra --config abc 。那么 cfgFile 的值为abc。
	rootCmd.PersistentFlags().StringVarP(&rootFlags.CfgFile, "config", "c", "", "指定配置文件")
	// Cobra 还支持本地 flags ，仅在直接调用此命令时才有意义。
	rootCmd.Flags().BoolP("toggle", "t", false, "关于toggle标志的帮助信息")

	// ######## 添加 配置 ########
	// ！！！注意！！！：Cobra 只有在上面的 Run 字段定义的函数运行之前才会解析手动指定的命令行 Flags，否则只能获取到代码中设置的 Flags 默认值。
	// 比如运行 go run main.go --config="abc.yaml" 时，rootFlags.CfgFile 并不会被赋值为 abc.yaml，而是默认值。
	// 此时有两种方式解决这个问题：
	// 1. 使用 Prase() 函数，提前解析 Flags：
	// rootCmd.PersistentFlags().Parse(os.Args)
	// 2. 使用 OnInitialize() 函数，该函数会在 Command.Run 字段指定的函数执行前，先执行 initConfig 函数。
	// 查看 Cobra 源码，OnInitialize() 中的 initializers 变量会在 preRun() 函数中被执行。
	cobra.OnInitialize(initConfig)
	// 假如我现在在这里执加了一行 config.NewConfig(rootFlags.CfgFile)，那么这个函数其实是会在 OnInitialize 函数执行之前执行的。
	// config.NewConfig(rootFlags.CfgFile)

	// ######## 添加 子命令 ########
	// 为了更好的管理子命令，我们通常会将子命令放在不同的文件中，然后在这里进行注册
	rootCmd.AddCommand(
		NewVersionCmd(),
		vipercmd.NewViperCmd(),
	)

	return rootCmd
}

func initConfig() {
	// 使用 Viper 简化处理配置文件的过程。Viper 可以从 JSON、TOML、YAML、HCL、环境变量和命令行参数等等地方中读取配置。
	config.NewConfig(rootFlags.CfgFile)
}

func rootRun(cmd *cobra.Command, args []string) {
	fmt.Println("主程序运行后执行的代码块。如果注销 Run，则运行主程序会显示上面Long上的信息")
	fmt.Println("在 Run 字段指定的函数中，我们可以获取到 Flags 的值：", rootFlags.CfgFile)
}
