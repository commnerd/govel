package route

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var desc = "Generate a route for the framework."

var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		viper.New()
		route := arg[0]
		routeSplit := strings.Split(route, "/")

		config := "routes/" + routeSplit[0] + ".yml"

		if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
			// path/to/whatever does not exist
		  }

		viper.Read
	},
}
