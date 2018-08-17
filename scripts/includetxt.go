package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"path"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	dir := "data"
	os.MkdirAll(dir, os.ModeDir | 0775)
	fs, _ := ioutil.ReadDir(dir)
	out, _ := os.Create(path.Join(dir,"textfiles.go"))
	out.Write([]byte("package data \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".yaml") {
			out.Write([]byte(strings.TrimSuffix(f.Name(), ".yaml") + " = `"))
			f, _ := os.Open(path.Join(dir, f.Name()))
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}