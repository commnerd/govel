package generate

import (
	"fmt"

	"github.com/commnerd/govel/tool/generate/controller"
	"github.com/commnerd/govel/tool/generate/route"
	"github.com/spf13/cobra"
)

var desc = "Generate various entities for the framework."

var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(desc)
	},
}

func init() {
	rootCmd.AddCommand(route.Cmd)
	rootCmd.AddCommand(controller.Cmd)
}

func main() {
	rootCmd.Execute()
}
