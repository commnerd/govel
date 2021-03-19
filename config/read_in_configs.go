package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func (c *Config) ReadInConfigs(path string) {
	config := buildConfig(path)

	c.ReadConfig(bytes.NewBuffer([]byte(config)))
}

func ReadInConfigs(path string) {
	cfg := Get()

	cfg.ReadInConfigs(path)
}

func buildConfig(path string) string {
	config := ""

	for _, fileInfo := range getYamlFileList(path) {

		// Get key from filename
		key := keyFromYamlFilename(fileInfo.Name())
		if key == "" {
			continue
		}

		config += key + ":\n"

		// Read file
		if string(path[len(path)-1]) != string(os.PathSeparator) {
			path += string(os.PathSeparator)
		}
		dat, err := ioutil.ReadFile(path + fileInfo.Name())
		if err != nil {
			panic(err)
		}

		contents := string(dat)

		// Trim whitespace
		contents = cleanupWhitespace(contents)

		// Indent contents and add file key
		contents = indentContents(contents)

		// Prepend major key
		config += contents + "\n"
	}

	return config
}

func cleanupWhitespace(content string) string {
	lines := strings.Split(content, "\n")

	for index, line := range lines {
		regex := regexp.MustCompile(`\t`)

		line = regex.ReplaceAllString(line, "    ")

		regex = regexp.MustCompile(`\s*$`)

		lines[index] = regex.ReplaceAllString(line, "")
	}

	return strings.Join(lines, "\n")
}

func indentContents(content string) string {
	contentLines := strings.Split(content, "\n")

	regex := regexp.MustCompile(`^\s*$`)

	for index, line := range contentLines {
		line = "    " + line
		if regex.MatchString(line) {
			line = ""
		}
		contentLines[index] = line
	}

	return strings.Join(contentLines, "\n")
}

func keyFromYamlFilename(name string) string {
	regex := regexp.MustCompile(`^(.*)\.(yaml|yml)$`)

	matches := regex.FindStringSubmatch(name)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func getYamlFileList(path string) []os.FileInfo {
	files := make([]os.FileInfo, 0)

	allFiles := getFileList(path)

	for _, f := range allFiles {
		if !f.IsDir() {
			regex := regexp.MustCompile(`\.(yaml|yml)$`)
			matches := regex.MatchString(f.Name())
			if matches {
				files = append(files, f)
			}
		}
	}

	return files
}

func getFileList(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}
