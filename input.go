package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func gloablGoPath() string {
	gpath := os.Getenv("GOPATH")

	if runtime.GOOS == "windows" {
		gpath = strings.Split(gpath, ";")[0]
	} else {
		gpath = strings.Split(gpath, ":")[0]
	}

	return gpath
}

func readInput(text, defaultVal string) string {
	if defaultVal == "" {
		fmt.Print(text, ":  ")
	} else {
		fmt.Printf("%s[%s]: ", text, defaultVal)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	value := scanner.Text()
	fmt.Println()
	if value == "" {
		return defaultVal
	}

	return value
}

func inputGoPath(c *container) {
	c.Gopath = readInput("Input go path", gloablGoPath())
}

func inputAddPackage(c *container) {
	packageNew := readInput("New package name", "")
	if packageNew != "" {
		c.List = append(c.List, packageNew)
	}
}

func inputCwd(c *container) {
	c.Cwd = readInput("Input cwd of your project", c.Cwd)
}

func inputInstall(c *container) {
	c.Install = readInput("Input package prefix (github.com/starmanmartin)", c.Install)
}

func inputPackageList(c *container) {
	c.List = strings.Split(readInput("package List to deploy (seperated by space)", strings.Join(c.List, " ")), " ")
}

func inputAll(c *container) {
	inputGoPath(c)
	inputCwd(c)
	inputInstall(c)
	inputPackageList(c)
}
