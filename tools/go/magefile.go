// +build mage

/*
mage -d "." -v check
*/

package main

import (
	"bootstrap-table/tools/go/api/utils"
	"bootstrap-table/tools/go/check"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	GoEXE   = os.Getenv("GOEXE")
	RootDir string
)

// Env
func init() {
	utils.Must(os.Setenv("GO111MODULE", "on"))
}

// const Variable
func init() {
	RootDir, _ = filepath.Abs("../..")
	fmt.Printf("⚠❗🐬 Please, make sure the root path is right: %s 🐬❗⚠\n", RootDir)
}

func Check() error {
	apiDir := filepath.Join(RootDir, "site/docs/api")
	collectFiles := func(dir string, extList []string) (fileList []string, err error) {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if utils.In(filepath.Ext(path), extList) { // .md, .rst
				fileList = append(fileList, path)
			}
			return nil
		})
		if err != nil {
			log.Fatalf("walk error [%v]\n", err)
			return nil, err
		}
		return fileList, nil
	}

	apiFilepathSlice, _ := collectFiles(apiDir, []string{".md"})
	if !check.APIDocsIsOK(apiFilepathSlice) {
		return errors.New("api documentation may not ok  please check again")
	}
	return nil
}

/*
func main() {
	Check()
}
*/
