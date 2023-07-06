/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package viper

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var viperCmd *cobra.Command

type ViperFlags struct {
	Test string
}

var viperFlags ViperFlags

// ViperCmd 封装
func NewViperCmd() *cobra.Command {
	// viperCmd 表示 viper 命令
	viperCmd = &cobra.Command{
		Use:   "viper",
		Short: "viper 库使用示例",
		Long:  ``,
		Run:   runViper,
	}

	cobra.OnInitialize(initConfig)

	viperCmd.Flags().StringVar(&viperFlags.Test, "test", "", "测试标志")

	return viperCmd
}

func initConfig() {
	// TODO: Viper 从各种地方读取配置的示例
	// viper.AutomaticEnv() 读取匹配的环境变量

	// 自动绑定配置文件中的 Key 与 Flags。BindPFlags 方法将 Flags 绑定到配置文件中，使用 Flags 的长名称作为配置文件中的 Key。
	viper.BindPFlags(viperCmd.Flags())

	// 手动绑定配置文件中的 Key 与 Flags。与 BindPFlags 不同，BindPFlag 方法可以指定配置文件中的 Key 名与指定的 Flags 绑定。
	viper.BindPFlag("manual-test", viperCmd.Flags().Lookup("test"))
	// 带层级的 Key 绑定。只用标准的 . 符号引用 Key 即可。
	viper.BindPFlag("level-one.level-two", viperCmd.Flags().Lookup("test"))
}

func runViper(cmd *cobra.Command, args []string) {
	fmt.Println("viper called")

	// 这里获取配置文件中的 test 键的值，如果 Flags 中使用 --test 指定了值，则覆盖配置文件中的值。
	fmt.Println(viper.GetString("test"))
	fmt.Println(viper.GetString("manual-test"))
	fmt.Println(viper.GetString("level-one.level-two"))
}
