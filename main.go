package main

import (
	"fmt"
	"path/filepath"

	"github.com/d0ughb0yy/goExtensionCheck/checks"
)

func main() {
	var userFile string
	fmt.Println("Enter full path to the file:")
	fmt.Scanf("%v", &userFile)

	newFile := checks.File{
		Path:      userFile,
		Name:      filepath.Base(userFile),
		Extension: filepath.Ext(userFile),
	}

	newFile.CheckHealth()

}
