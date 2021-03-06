package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	conf = &Config{Viper: viper.New()}

	sep := string(os.PathSeparator)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configPath = path + sep + "config"
	populateConfig()
}

func trimExtension(filename string) string {
	var extension = filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

func getConfigFileList() []string {
	var files []string

	err := filepath.Walk(configPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func populateConfig() {
	config := ""

	files := getConfigFileList()
	for _, file := range files {

		pathParts := strings.Split(file, string(os.PathSeparator))
		fileName := pathParts[len(pathParts)-1]

		config += trimExtension(fileName) + ":\n"
		dat, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		data := string(dat)

		regex := regexp.MustCompile("\t")
		data = regex.ReplaceAllString(data, "    ")

		regex = regexp.MustCompile("^")
		data = regex.ReplaceAllString(data, "    ")

		config += string(data)
	}

	fmt.Println(config)
	// viper.SetConfigType("yaml")

	// (*Config).ReadConfig(bytes.NewBuffer([]byte(config)))

	conf.Viper.ReadConfig(bytes.NewBuffer([]byte(config)))

	// fmt.Println(file)
	// fmt.Println(viper.AllKeys())
}
