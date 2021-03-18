package test

import (
	"os/exec"
	_ "log"
	"fmt"
	"github.com/spf13/cobra"
)

var desc = "Test the application."

var Cmd = &cobra.Command{
	Use:   "test",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("go", "test", "tests").Output()
		cobra.CheckErr(err)
		fmt.Println(out)
	},
}
