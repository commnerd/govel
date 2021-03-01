package main

import (
	"fmt"

	"github.com/commnerd/govel/govel/new"

	"github.com/spf13/cobra"
)

//go:generate bash -c "ls -d -- */)"

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
