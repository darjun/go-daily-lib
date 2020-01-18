package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	dividedByZeroHanding int // 除 0 如何处理
)

var divideCmd = &cobra.Command {
	Use: "divide",
	Short: "Divide subcommand divide all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calc(values, DIVIDE)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "/"), result)
	},
}

func init() {
	divideCmd.Flags().IntVarP(&dividedByZeroHanding, "divide_by_zero", "d", int(PanicOnDividedByZero), "do what when divided by zero")

	rootCmd.AddCommand(divideCmd)
}