package main

import (
	"encoding/hex"
	"io/ioutil"
	"os"
)

func main() {
	dat, _ := ioutil.ReadFile("../cmd/cmd")
	hx := hex.EncodeToString(dat)

	f, _ := os.Create("../govel/new/cmdTool.go")

	defer f.Close()

	var prefix = []byte("package new\n\nconst CMD_TOOL = \"")
	var suffix = []byte("\"")
	byteContents := append(prefix, hx...)
	byteContents = append(byteContents, suffix...)
	_, _ = f.Write(byteContents)
}
