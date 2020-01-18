package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type ErrorHandling int

const (
	ContinueOnParseError 	ErrorHandling = 1 // 解析错误尝试继续处理
	ExitOnParseError 		ErrorHandling = 2 // 解析错误程序停止
	PanicOnParseError		ErrorHandling = 3 // 解析错误 panic
	ReturnOnDividedByZero	ErrorHandling = 4 // 除0返回
	PanicOnDividedByZero	ErrorHandling = 5 // 除0 painc
)

type OpType int

const (
	ADD 		OpType = 1
	MINUS 		OpType = 2
	MULTIPLY 	OpType = 3
	DIVIDE 		OpType = 4
)

var (
	parseHandling int
)

var rootCmd = &cobra.Command {
	Use: "math",
	Short: "Math calc the accumulative result.",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized subcommand"))
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&parseHandling, "parse_error", "p", int(ContinueOnParseError), "do what when parse arg error")
}

func Execute() {
	rootCmd.Execute()
}