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

// Execute 将所有子命令添加到根命令并设置 Flags。这由 main.main() 调用。它只需要对 rootCmd 发生一次。
func Execute() {
	app := newApp()
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}

}

type RootFlags struct {
	// 这里定义的变量，可以在下面的 init 函数中，通过 rootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "指定配置文件(默认在$HOME/.cobracli.yaml)") 进行绑定
	// 也可以通过 viper 进行绑定
	CfgFile string
}

var rootFlags RootFlags

func newApp() *cobra.Command {
	// rootCmd 表示在没有任何子命令调用的情况时的基本命令。
	// 通常来说，我们会将命令实例化的操作封装到一个函数中，以便对 XXXCmd 变量执行更多操作。
	// 我们只要让这个函数返回 *cobra.Command 类型的变量即可。
	var rootCmd = &cobra.Command{
		Use:   "go-cobra",
		Short: "这个应用简要的描述",
		Long: `横跨多行的较长描述，可能包含示例和使用应用程序的用法。 例如：
当我运行程序时，会显示该描述内容
	如果使用缩进，这行在界面展示时有缩进。`,
		// 如果这个应用没有任何子命令，直接使用 go-cobra 执行的话，将会执行下面 Run 字段指定的函数
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("主程序运行后执行的代码块。如果注销 Run，则运行主程序会显示上面Long上的信息")
			fmt.Println("让我们看看命令行 Flags 的值：", rootFlags.CfgFile)
		},
	}

	// 我们可以在这里定义命令行 Flags 和 配置设置。
	// 比如初始化数据库连接、初始化日志、读取配置文件等等

	// ######## 添加 命令行Flags ########
	// Cobra 支持 持久性flags (i.e. Global Flags)，如果在这个位置定义，则这些 flags 对应用程序来说是全局的。
	// 第一个参数是变量，用于存储该flag的值；第二个参数为该flag的名字；第三个参数为该flag的默认值,无默认值可以为空；第四个参数是该flag的描述信息
	// 比如我现在使用如下命令: go-cobra --config abc 。那么 cfgFile 的值为abc。
	rootCmd.PersistentFlags().StringVarP(&rootFlags.CfgFile, "config", "c", "", "指定配置文件")
	// Cobra 还支持本地 flags ，仅在直接调用此命令时才有意义。
	rootCmd.Flags().BoolP("toggle", "t", false, "关于toggle标志的帮助信息")

	// ！！！注意！！！：Cobra 只有在上面的 Run 字段定义的函数运行时才会解析手动指定的命令行 Flags，否则只有代码中设置默认值可以获取到。
	// 比如运行 go run main.go --config="abc.yaml" 时，rootFlags.CfgFile 并不会被赋值为 abc.yaml，而是默认值。
	fmt.Printf("手动解析之前，检查手动指定的 Flags 的值为：%s\n", rootFlags.CfgFile)
	// 如果想要在这里就获取到手动指定的 Flags 中的 rootFlags.CfgFile 的值，必须要提前解析一下，如下：
	rootCmd.PersistentFlags().Parse(os.Args)
	fmt.Printf("手动解析之后，检查手动指定的 Flags 的值为：%s\n", rootFlags.CfgFile)

	// ######## 添加 配置文件 ########
	// 使用 Viper 简化处理配置文件的过程。Viper 可以从 JSON、TOML、YAML、HCL、环境变量和命令行参数等等地方中读取配置。
	config.NewConfig(rootFlags.CfgFile)

	// ######## 添加 子命令 ########
	// 为了更好的管理子命令，我们通常会将子命令放在不同的文件中，然后在这里进行注册
	rootCmd.AddCommand(
		NewHelloWorldCmd(),
		vipercmd.NewViperCmd(),
	)

	return rootCmd
}
