/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd 表示在没有任何子命令的情况下调用时的基本命令
var rootCmd = &cobra.Command{
	Use:   "cobracli",
	Short: "A brief description of your application",
	Long: `横跨多行的较长描述，可能包含示例和使用应用程序的用法。 例如：
当我运行程序时，会显示该描述内容
	如果使用缩进，这行在界面展示时有缩进`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("主程序运行后执行的代码块。如果注销Run，则运行主程序会显示上面Long上的信息")
	},
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
	// 有两种 Flags 类型可用。在使用中的显示就是 Falgs 和 Global Flags 的区别。

	// Cobra 支持 持久性flags (i.e. Global Flags)，如果在这个位置定义，则这些 flags 对应用程序来说是全局的。
	// 第一个参数是变量，用于存储该flag的值；第二个参数为该flag的名字；第三个参数为该flag的默认值,无默认值可以为空；第四个参数是该flag的描述信息
	// 比如我现在使用如下命令：cobracli --config abc 。那么cfgFile的值为abc。
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "指定配置文件(默认在$HOME/.cobracli.yaml)")

	// Cobra 还支持本地 flags ，仅在直接调用此命令时才有意义。
	rootCmd.Flags().BoolP("toggle", "t", false, "关于toggle标志的帮助信息")
}

// initConfig 读取配置文件和ENV变量(如果已设置).
func initConfig() {
	if cfgFile != "" {
		// 使用 flag 传递过来的配置文件
		viper.SetConfigFile(cfgFile)
		fmt.Println("当使用 --config 标签时，打印该内容，config的值为:", cfgFile)
	} else {
		// 查找 home 目录.
		home, err := homedir.Dir()
		fmt.Println("不使用 --config 标签时，打印该内容,config的默认值为：", cfgFile)
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
