package check

import (
	"bootstrap-table/tools/go/api/file"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func init() {

}

func APIDocsIsOK(docsFileSlice []string) bool {
	isOK := true
	for _, curFilePath := range docsFileSlice {
		curFile := file.MustFile(curFilePath)
		scanner := bufio.NewScanner(curFile)

		lineNumber := 1
		regex := regexp.MustCompile(`\*\*Attribute:\*\* \x60(?P<prefix>.+-.+)\x60`) // \x60 = character `
		groupIndex := regex.SubexpIndex("prefix")

		for scanner.Scan() { // https://golang.org/pkg/bufio/#Scanner.Scan
			curRowString := scanner.Text()
			matchSlice := regex.FindStringSubmatch(curRowString)
			if len(matchSlice) > (groupIndex - 1) {
				matchData := matchSlice[groupIndex]
				if !strings.HasPrefix(matchData, "data") {
					fmt.Printf("%s | %d | %s | %s\n", curFile.Name(), lineNumber, curRowString, matchData)
					isOK = false
				}
			}
			lineNumber++
		}
		if err := curFile.Close(); err != nil {
			panic(err)
		}
	}
	return isOK
}
