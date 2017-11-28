package main

import (
	"fmt"
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func main() {
	path, _ := os.Getwd()

	if ostype == "windows" {
		fmt.Fprintln(os.Stdout, "ostype: %s", ostype)
		fmt.Fprintln(os.Stdout, "path: %s", path)
	} else if ostype == "linux" {
		fmt.Fprintln(os.Stdout, "ostype: %s", ostype)
		fmt.Fprintln(os.Stdout, "path: %s", path)
	}

	os.Create("./bb/aa.txt")
}
