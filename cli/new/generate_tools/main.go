package main

import (
	"encoding/hex"
	"io/ioutil"
	"os"
)

func main() {
	dat, _ := ioutil.ReadFile("../../tool/tool")
	hx := hex.EncodeToString(dat)

	f, _ := os.Create("tool_bin.go")

	defer f.Close()

	var prefix = []byte("package new\n\nconst TOOL_BIN = \"")
	var suffix = []byte("\"")
	byteContents := append(prefix, hx...)
	byteContents = append(byteContents, suffix...)
	_, _ = f.Write(byteContents)
}
