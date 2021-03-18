package controller

import (
	"github.com/spf13/cobra"
)

var desc = "Generate a route for the framework."

var Cmd = &cobra.Command{
	Use:   "controller",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		// viper.New()
		// controller := arg[0]
		// pathSplit := strings.Split(controller, "/")

		// if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		// 	// path/to/whatever does not exist
		// }
	},
}
