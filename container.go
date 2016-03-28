package main

import (
	"encoding/json"
	"io/ioutil"
    "os"
)

func readFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}


	return data, nil
}

type container struct {
	Cwd, Gopath string
	List        []string
	Install     string
}

func readJSON(file string) (*container, error) {
    c := &container{}
    data, err := readFile(file)
    if err != nil {
		return c, err
	}
    
    err = json.Unmarshal(data, c)    
    return c, err    
}

func (c *container) saveJSON(file string) error{
    fileContent, err := json.Marshal(c)
    if err != nil {
        return err
    }
    
    return ioutil.WriteFile(file, fileContent, os.ModePerm)
}