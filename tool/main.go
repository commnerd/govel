package main

import (
	"fmt"

	"github.com/commnerd/govel/tool/serve"
	"github.com/commnerd/govel/tool/test"

	"github.com/spf13/cobra"
)

var desc = "A Swiss Army Knife tool for Govel"

var rootCmd = &cobra.Command{
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(desc)
	},
}

func init() {
	rootCmd.AddCommand(serve.Cmd)
	rootCmd.AddCommand(test.Cmd)
}

func main() {
	rootCmd.Execute()
}
