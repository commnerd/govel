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

//go:generate bash -c "cd ../../cmd && go build ."
//go:generate bash -c "cd ../../scripts && go run ."
//go:generate bash -c "rm ../../cmd/cmd"

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
		addCmd(baseDir)
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
	}

	for _, subDir := range subDirs {
		path := strings.Join(strings.Split(baseDir+"/"+subDir, "/"), string(os.PathSeparator))
		fmt.Printf("Creating subdir %s.\n", path)
		cobra.CheckErr(os.Mkdir(fmt.Sprintf("%s", path), 0755))
	}
}

func addBaseFiles(baseDir string) {
	writeFile(baseDir+"/main.go", MAIN_GO)
	writeFile(baseDir+"/config/main.yml", CONFIG_MAIN)
	writeFile()
}

func addCmd(baseDir string) {
	binary, _ := hex.DecodeString(cmdTool)
	fmt.Printf("Placing 'cmd' into '%s' project.\n", baseDir)
	err := ioutil.WriteFile(
		strings.Join(strings.Split(baseDir+"/cmd", "/"), string(os.PathSeparator)),
		binary, 0755)
	cobra.CheckErr(err)
}

func writeFile(path, contents string) {
	path = strings.Join(strings.Split(path, "/"), string(os.PathSeparator))
	f, err := os.Create(path)
	cobra.CheckErr(err)

	defer f.Close()

	byteContents := []byte(contents)
	_, err = f.Write(byteContents)
	cobra.CheckErr(err)
}
