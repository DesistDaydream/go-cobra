/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloWorldCmd 封装
func NewHelloWorldCmd() *cobra.Command {
	// helloWorldCmd 表示 helloWorld 命令
	helloWorldCmd := &cobra.Command{
		Use:   "helloWorld",
		Short: "这个命令的简要描述",
		Long:  `横跨多行的较长描述，可能包含示例和使用命令的用法。`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("helloWorld called")
		},
	}

	return helloWorldCmd
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloWorldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloWorldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
