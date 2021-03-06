package main

import (
	"fmt"

	"github.com/commnerd/govel/cli/new"

	"github.com/spf13/cobra"
)

var desc = "Govel is a web framework written in golang based on Laravel"

var rootCmd = &cobra.Command{
	Use:   "govel",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(desc)
	},
}

func init() {
	rootCmd.AddCommand(new.Cmd)
}

func main() {
	rootCmd.Execute()
}
