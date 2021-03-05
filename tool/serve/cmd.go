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

		cmdCmd := strings.Join(strings.Split(fmt.Sprintf("./%s", path), "/"), string(os.PathSeparator))
		fmt.Println("Running " + cmdCmd)
		c = exec.Command(cmdCmd)
		cobra.CheckErr(c.Run())
	},
}
