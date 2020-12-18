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

	"github.com/spf13/cobra"
)

// testCmd 表示test子命令
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "在执行主程序的--help时，子命令的短描述就在这里",
	Long:  `关于使用 test 子命令时的帮助信息和提示`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("执行 test 子命令时，执行该代码块")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// 这里可以定义该子命令独有的 flags 和 配置设置

	// Cobra 支持Persistent Flags，这将对该子命令和该子命令下的所有子命令起作用，例如：
	testCmd.PersistentFlags().String("foo", "", "这是关于test子命令中--foo标签的帮助信息")

	// Cobra 支持仅在直接调用此命令时运行的本地 Flags，e.g.：
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
