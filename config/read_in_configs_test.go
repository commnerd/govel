package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileList(t *testing.T) {
	writeDefaultTestFiles()
	defer cleanupDefaultTestFiles()

	files := getFileList("./test")

	assert.Equal(t, 3, len(files))
	assert.Equal(t, "f1.yaml", files[0].Name())
	assert.Equal(t, "f2.yml", files[1].Name())
	assert.Equal(t, "f3.tmp", files[2].Name())
}

func TestGetYamlFileList(t *testing.T) {
	writeDefaultTestFiles()
	defer cleanupDefaultTestFiles()

	files := getYamlFileList("./test")

	assert.Equal(t, 2, len(files))
	assert.Equal(t, "f1.yaml", files[0].Name())
	assert.Equal(t, "f2.yml", files[1].Name())
}

func TestKeyFromFilename(t *testing.T) {
	name1 := "f1.yaml"
	name2 := "f2.yml"
	name3 := "f3.tmp"

	assert.Equal(t, "f1", keyFromYamlFilename(name1))
	assert.Equal(t, "f2", keyFromYamlFilename(name2))
	assert.Equal(t, "", keyFromYamlFilename(name3))
}

func TestIndentContents(t *testing.T) {
	content := `foo: "bar"
    
    biz: "baz"`

	assert.Equal(t, `    foo: "bar"

        biz: "baz"`,
		indentContents(content))
}

func TestCleanupWhitespace(t *testing.T) {
	content := `foo: "bar"     

	biz: "baz"`

	assert.Equal(t, `foo: "bar"

    biz: "baz"`,
		cleanupWhitespace(content))
}

func TestBuildConfig(t *testing.T) {
	writeDefaultTestFiles()
	defer cleanupDefaultTestFiles()

	config := buildConfig("./test")

	assert.Equal(t, `f1:
    foo: "bar"
f2:
    foo: "bar"
`, config)
}

func TestReadInConfigs(t *testing.T) {
	writeDefaultTestFiles()
	defer cleanupDefaultTestFiles()

	instance.ReadInConfigs("./test")

	assert.Equal(t, "bar", instance.GetString("f1.foo"))
	assert.Equal(t, "bar", instance.GetString("f2.foo"))
}

func writeDefaultTestFiles() {
	createDir("./test")

	writeFile("./test/f1.yaml", `foo: "bar"`)
	writeFile("./test/f2.yml", `foo: "bar"`)
	writeFile("./test/f3.tmp", `foo: "bar"`)
}

func cleanupDefaultTestFiles() {
	os.RemoveAll("./test")
}

func createDir(dirName string) {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		panic(err)
	}
}

func writeFile(path, content string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	byteContents := []byte(content)
	_, err = f.Write(byteContents)
	if err != nil {
		panic(err)
	}
}
