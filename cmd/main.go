package main

import (
	"fmt"

	"cmd/serve"

	"github.com/spf13/cobra"
)

var desc = "A Swiss Army Knife tool for Govel"

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(desc)
	},
}

func init() {
	rootCmd.AddCommand(serve.Cmd)
}

func main() {
	rootCmd.Execute()
}
