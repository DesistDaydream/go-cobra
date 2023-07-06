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
	viper.BindPFlags(viperCmd.Flags())
}

func runViper(cmd *cobra.Command, args []string) {
	fmt.Println("viper called")

	fmt.Println(viper.GetString("test"))
}
