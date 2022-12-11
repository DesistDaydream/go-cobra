/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package viper

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ViperCmd 封装
func NewViperCmd() *cobra.Command {
	// viperCmd 表示 viper 命令
	viperCmd := &cobra.Command{
		Use:   "viper",
		Short: "viper 库使用示例",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("viper called")
		},
	}

	return viperCmd
}

func init() {
	// TODO: Viper 从各种地方读取配置的示例
	// viper.AutomaticEnv() 读取匹配的环境变量
}
