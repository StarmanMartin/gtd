package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"github.com/kardianos/osext"
	"github.com/starmanmartin/simple-fs"
	"github.com/starmanmartin/goconfig"
)

const (
	packagePlaceholder = "%package%"
)

func getPakckageList() ([]string, []string) {
	cwd, _ := goconfig.GetString("cwd")
	list, ok := goconfig.GetArrayString("list")
	if !ok {
		log.Panicln("No list in package.json")
	}

	pathList := make([]string, len(list))
	for i, val := range list {
		pathList[i] = strings.Replace(cwd+"/"+val+"/src", "//", "/", -1)
	}

	return pathList, list
}

func getGoPath() string {
	gpath, ok := goconfig.GetString("gopath")
	if !ok {
		gpath = os.Getenv("GOPATH")
	}
	if runtime.GOOS == "windows" {
		gpath = strings.Split(gpath, ";")[0] + "/src"
	} else {
		gpath = strings.Split(gpath, ":")[0] + "/src"
	}

	return strings.Replace(gpath, "//", "/", -1)
}

func prepareInstall() string {
	gpath, ok := goconfig.GetString("install")
	if !ok {
		gpath = os.Getenv("GOPATH")
	}

	gpath = strings.Split(gpath, ";")[0] + "/src"
	return strings.Replace(gpath, "//", "/", -1)
}

func getCmd(cmdCommand string) *exec.Cmd {
	parts := strings.Split(cmdCommand, " ")
	head := parts[0]
	parts = parts[1:]

	cmd := exec.Command(head, parts...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

func exeInstall(packagename string) error {
	cmdCommand, ok := goconfig.GetString("install")
	if !ok {
		return errors.New("No install command")
	}

	cmdCommand = strings.Replace(cmdCommand, packagePlaceholder, packagename, -1)

	cmd := getCmd(cmdCommand)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	log.Println("Started test deployment")
	cwd, _ := osext.ExecutableFolder()
    
	if err := goconfig.InitConficOnce(cwd + "/gdtConfig/packages.json"); err != nil {
		log.Panicln(err)
		return
	}

	pathList, packageList := getPakckageList()
	gopath := getGoPath()

	for i, v := range pathList {
		if err := fs.CopyFolderAndIngonre(v, gopath, ".git"); err != nil {
			log.Println(err)
		}
		if err := exeInstall(packageList[i]); err != nil {
			log.Println("Error: ", err)
		}

		log.Println("Copied", packageList[i])
	}

}
