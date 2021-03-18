package serve

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var desc = "Serve up the application using the framework."

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("go", "build", ".")
		cobra.CheckErr(c.Run())

		path, err := os.Getwd()
		cobra.CheckErr(err)

		splitPath := strings.Split(path, string(os.PathSeparator))
		path = splitPath[len(splitPath)-1]

		toolCmd := "." + string(os.PathSeparator) + path
		fmt.Println("Running " + toolCmd)
		c = exec.Command(toolCmd)
		cobra.CheckErr(c.Run())
	},
}
