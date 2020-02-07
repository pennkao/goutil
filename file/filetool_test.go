package file

import (
	"fmt"
	"testing"
)

var filePath = "/Users/Penn/"

func TestReadDirFiles(t *testing.T) {
	rs, err := ReadDirFiles(filePath)
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}

func TestReadDirDirs(t *testing.T) {
	rs, err := ReadDirDirs(filePath)
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}

func TestReadDir(t *testing.T) {
	rs, err := ReadDir(filePath)
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}

func TestSearchFile(t *testing.T) {
	rs, err := SearchFile("Cellar", filePath, "/usr/local")
	if err != nil {
		t.Error("DirsUnder: ", err.Error())
	}
	fmt.Println(rs)
}
