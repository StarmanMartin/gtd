package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
    "os/user"
	"github.com/starmanmartin/simple-fs"
    "flag"
)

const (
    gdtDir = "/.gtdConfig"
)

var (
	isAdd, isRest bool
    goconfig *container
)

func init(){
    flag.BoolVar(&isAdd, "a", false, "Add package to deployment list")
	flag.BoolVar(&isRest, "r", false, "Reset config File")
}

func getPakckageList() ([]string, []string) {
	cwd := goconfig.Cwd
	list := goconfig.List	

	pathList := make([]string, len(list))
	for i, val := range list {
		pathList[i] = strings.Replace(cwd+"/"+val+"/src", "//", "/", -1)
	}

	return pathList, list
}

func getGoPath() string {
	gpath := goconfig.Gopath
	return strings.Replace(gpath, "//", "/", -1) + "/src"
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
	cmdCommand := goconfig.Install
	if cmdCommand == "" {
		return errors.New("No install command")
	}

    cmdCommand = "go install " + cmdCommand +  "/" + packagename

	cmd := getCmd(cmdCommand)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func getGdtDir() string {
    cUser, err := user.Current()
    
    if err != nil {
        log.Panicln(err)
    }
    
    gdtPath := cUser.HomeDir + gdtDir
    
    if exists, readErr := fs.Exists(gdtPath); !exists {
        os.MkdirAll(gdtPath, os.ModePerm)
    } else if readErr != nil {
        log.Panicln(readErr)
    }
    
    return gdtPath
}

func main() {
    flag.Parse()    
    var err error
    if goconfig, err = readJSON(getGdtDir() + "/packages.json"); err != nil || isRest {	
        inputAll(goconfig)
        goconfig.saveJSON(getGdtDir() + "/packages.json")
		return
	}
    
    if isAdd {
        inputAddPackage(goconfig)
        goconfig.saveJSON(getGdtDir() + "/packages.json")
    } 
    
	log.Println("Started test deployment")	

	pathList, packageList := getPakckageList()
	gopath := getGoPath()

	for i, v := range pathList {
		if err := fs.SyncFolderAndIngonre(v, gopath, ".git"); err != nil {
			log.Println(err)
		}

		log.Println("Copied", packageList[i])
        
		if err := exeInstall(packageList[i]); err != nil {
			log.Println("Error: ", err)
		}
	}

}
