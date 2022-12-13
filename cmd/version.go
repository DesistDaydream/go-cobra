/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd 封装
func NewVersionCmd() *cobra.Command {
	// versionCmd 表示 version 命令
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "这个命令的简要描述",
		Long:  `横跨多行的较长描述，可能包含示例和使用命令的用法。`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version called")
		},
	}

	return versionCmd
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
