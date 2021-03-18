package route

import (
	"github.com/spf13/cobra"
)

var desc = "Generate a route for the framework."

var Cmd = &cobra.Command{
	Use:   "route",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		// path/to/whatever does not exist
	},
}
