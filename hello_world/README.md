# cobra 介绍
是一个golang的库，其提供简单的接口来创建强大现代的CLI接口，类似于git或者go工具。同时，它也是一个应用，用来生成个人应用框架，从而开发以Cobra为基础的应用。热门的docker和k8s源码中都使用了Cobra。

比如kubeadm和kubectl这种命令行工具，就是根据cobra来实现的，也可以把cobra当做一种框架，是一种想开发命令行工具所用的框架。

# cobra 基础使用
首先需要安装cobra  
`go get -u github.com/spf13/cobra/cobra`  
安装完成后会创建一个二进制文件 cobra 在 $GOPATH/bin 目录中，然后可以使用 cobra 命令来生成代码模板  
`cobra init --pkg-name cobracli`  该命令会在当前目录创建基本的代码，目录结构如下
```
[root@lichenhao cobra]# tree
.
├── cmd
│   └── root.go
├── LICENSE
└── main.go
```
默认情况下，Cobra将添加Apache许可证(LICENSE文件)。如果不想这样，可以将标志添加-l none到所有生成器命令。但是，它会在每个文件（// Copyright © 2018 NAME HERE ）的顶部添加版权声明。如果通过选项，-a YOUR NAME 则索赔将包含姓名。这些标志是可选的。  
其中 main.go 的初始代码如下：
```
package main

import "cobracli/cmd"

func main() {
	cmd.Execute()
}
```
其中 root.go 的初始代码如下：
```
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd 表示在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "cobracli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// 如果裸应用程序具有与之关联的操作，则取消注释以下行并在 {} 中写入要执行操作的代码：
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute 将所有子命令添加到root命令并适当设置flags
// 这由 main 包中的 main() 调用。 它只需要对 rootCmd 发生一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// 这里可以定义 flags 和 配置设置。
	// Cobra 支持 持久性flags ，如果在这个位置定义，则这些 flags 对应用程序来说是全局的
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobracli.yaml)")

    // Cobra 还支持本地 flags ，仅在直接调用此操作时才运行。
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig 读取配置文件和ENV变量(如果已设置).
func initConfig() {
	if cfgFile != "" {
		// 使用 flag 传递过来的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找 home 目录.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 在主目录中搜索名称为“ .cobracli”的配置（无扩展名）。
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobracli")
	}

	viper.AutomaticEnv() // 读取匹配的环境变量

	// 如果找到配置文件，则读入它。
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
```
使用`go mod init cobracli`命令初始化mod之后，即可使用`go run main.go`运行代码了

## cobra 子命令
使用`cobra add test`命令即可为程序添加名为 test 的子命令，并在 cmd 目录下添加 test.go 代码文件。
test.go 代码初始如下:
```
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// 这里可以定义该子命令独有的 flags 和 配置设置
	
	// Cobra支持Persistent Flags，这将对该子命令和该子命令下的所有子命令起作用，例如：
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra支持仅在直接调用此命令时运行的本地标志，例如：
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
```
# cobra 初体验总结
在 cobra 生成的代码中会有类似这样的变量 XXXCmd ，比如rootCmd，testCmd，其中root该程序的就是主命令(比如改程序为cobracli，则在命令行执行的时候，就是用`cobracli --help`这样的形式来使用)；test就是子命令。这种变量的用处类似于结构体，作用在这些变量的方法都是对这些命令或者子命令来进行操作。