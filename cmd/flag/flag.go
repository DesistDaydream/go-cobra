package flag

import (
	"fmt"

	"github.com/spf13/cobra"
)

type flagFlag struct {
	stringFlag      string
	stringSliceFlag []string
}

var f flagFlag

func NewFlagCmd() *cobra.Command {
	flagCmd := &cobra.Command{
		Use:   "flag",
		Short: "命令行标志 Demo",
		Run:   runFlag,
	}

	flagCmd.Flags().StringVar(&f.stringFlag, "string-flag", "", "字符串标志")
	flagCmd.PersistentFlags().StringSliceVar(&f.stringSliceFlag, "slice-flag", nil, "字符串切片标志")

	// 将指定的 Flag 设置为必须的，若不指定标志的值，则会报错。就算有默认值也会报错，报错内容:
	// Error: required flag(s) "string-flag" not set
	flagCmd.MarkFlagRequired("string-flag")
	// 如果想要将指定的 Persistent Flag 设置为必须的，则使用 MarkPersistentFlagRequired 方法
	flagCmd.MarkPersistentFlagRequired("slice-flag")
	// TODO: 暂时没有找到判断 Flag 标志为空的方法，只能自己写这个逻辑

	return flagCmd
}

func runFlag(cmd *cobra.Command, args []string) {
	fmt.Println("call flag cmd")
}
