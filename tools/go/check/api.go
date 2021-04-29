package check

import (
	"bootstrap-table/tools/go/api/file"
	"bootstrap-table/tools/go/api/regexhelper"
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type cError struct {
	MissHTMLSuffixError error
	AttributeError      error
}

var Error cError

func init() {
	Error.AttributeError = errors.New("missing prefix data")
	Error.MissHTMLSuffixError = errors.New("missing suffix html")
}

func APIDocsIsOK(docsFileSlice []string) bool {
	isOK := true

	// regexAttr := regexp.MustCompile(`\*\*Attribute:\*\* \x60(?P<prefix>.+-.+)\x60`) // \x60 = character `
	// regexExample := regexp.MustCompile(`\[.*\]\(.*\/(.*\.html)\)`) // Too verbose. Integrate to one as below
	regex := regexp.MustCompile(`(\*\*(?P<name>.*):\*\*){1} (?:(\[.*\]\((?P<examplelink>.*)\))|(\x60(?P<attrprefix>.+-.+)\x60))`) // https://regex101.com/r/62LqIi/2

	attrCheck := func(text string, matchSlice []string) (string, error) {
		// check "The data-* attributes is used to store custom data private to the page or application." https://www.w3schools.com/tags/att_global_data.asp
		matchData := regexhelper.GetGroupValue(regex, matchSlice, "attrprefix")
		if !strings.HasPrefix(matchData, "data") {
			return matchData, Error.AttributeError
		}
		return "", nil
	}

	exampleLinkCheck := func(text string, matchSlice []string) (string, error) {
		// check [](.html)
		matchData := regexhelper.GetGroupValue(regex, matchSlice, "examplelink")
		if !strings.HasSuffix(matchData, ".html") {
			return "", Error.MissHTMLSuffixError
		}
		return matchData, nil
	}

	checkFuncMap := map[string]func(text string, matchSlice []string) (string, error){
		"Attribute": attrCheck,
		"Example":   exampleLinkCheck,
	}

	for _, curFilePath := range docsFileSlice {
		curFile := file.MustFile(curFilePath)
		scanner := bufio.NewScanner(curFile)

		lineNumber := 1
		for scanner.Scan() { // https://golang.org/pkg/bufio/#Scanner.Scan
			curRowString := scanner.Text()
			matchSlice := regex.FindStringSubmatch(curRowString)
			if len(matchSlice) == 0 {
				continue
			}
			name := regexhelper.GetGroupValue(regex, matchSlice, "name")
			if name == "" {
				continue
			}

			if curCheckFunc, keyExists := checkFuncMap[name]; keyExists {
				if matchData, err := curCheckFunc(curRowString, matchSlice); err != nil {
					fmt.Printf("%s | %d | %s | %s | %s\n", curFile.Name(), lineNumber, curRowString, err.Error(), matchData)
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
