package new

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:generate bash -c "go run ./generate_tools/main.go"

var desc = "Generate a new Govel application."

var Cmd = &cobra.Command{
	Use:   "new",
	Short: desc,
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 || len(args) > 1 {
			cobra.CheckErr(errors.New("Govel's \"new\" command's only argument should be the name of the project you want to create."))
		}
		baseDir := args[0]
		bootstrapDirectoryStructure(baseDir)
		addBaseFiles(baseDir)
	},
}

func bootstrapDirectoryStructure(baseDir string) {
	fmt.Printf("Creating subdir %s.\n", baseDir)
	cobra.CheckErr(os.Mkdir(baseDir, 0755))

	subDirs := []string{
		"models",
		"controllers",
		"config",
		"database",
		"public",
		"resources",
		"routes",
		"storage",
		"tests",
		"tests/controllers",
	}

	for _, subDir := range subDirs {
		path := strings.Join(strings.Split(baseDir+"/"+subDir, "/"), string(os.PathSeparator))
		fmt.Printf("Creating subdir %s.\n", path)
		cobra.CheckErr(os.Mkdir(fmt.Sprintf("%s", path), 0755))
	}
}

func addBaseFiles(baseDir string) {
	writeFile(baseDir+"/main.go", fmt.Sprintf(MAIN_GO, baseDir))
	writeFile(baseDir+"/config/main.yml", CONFIG_MAIN)
	writeFile(baseDir+"/controllers/root.go", CONTROLLERS_ROOT)
	writeFile(baseDir+"/controllers/init.go", CONTROLLERS_INIT)
	writeFile(baseDir+"/routes/web.go", ROUTES_WEB)
	writeFile(baseDir+"/go.mod", fmt.Sprintf(GOMOD, baseDir, baseDir))
	writeFile(baseDir+"/tests/controllers/root_test.go", fmt.Sprintf(TESTS_CONTROLLERS_ROOT, baseDir))
	addTool(baseDir)

}

func addTool(baseDir string) {
	binary, _ := hex.DecodeString(TOOL_BIN)
	fmt.Printf("Placing 'cmd' into '%s' project.\n", baseDir)
	err := ioutil.WriteFile(
		strings.Join(strings.Split(baseDir+"/tool", "/"), string(os.PathSeparator)),
		binary, 0755)
	cobra.CheckErr(err)
}

func writeFile(path, contents string) {
	splitPath := strings.Split(path, "/")
	fileName := splitPath[len(splitPath)-1]
	baseDir := strings.Join(splitPath[0:len(splitPath)-1], string(os.PathSeparator))
	path = strings.Join(splitPath, string(os.PathSeparator))
	f, err := os.Create(path)
	cobra.CheckErr(err)

	defer f.Close()

	fmt.Printf("Placing '%s' into '%s' project.\n", fileName, baseDir)
	byteContents := []byte(contents)
	_, err = f.Write(byteContents)
	cobra.CheckErr(err)
}
