package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command {
    Use: "clone url [destination]",
    Short: "Clone a repository into a new directory",
    Run: func(cmd *cobra.Command, args []string) {
        output, err := ExecuteCommand("git", "clone", args...)
        if err != nil {
            Error(cmd, args, err)
        }
        
        fmt.Fprintf(os.Stdout, output)
    },
}

func init() {
    rootCmd.AddCommand(cloneCmd)
}