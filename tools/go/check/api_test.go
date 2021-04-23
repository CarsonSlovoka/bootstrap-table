package check

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func removeTestFile(filepath string) {
	if _, err := os.Stat(filepath); err == nil { // check file exists
		err := os.Remove(filepath)
		if err != nil {
			panic(err)
		}
	}
}

func TestAPIDocsIsOK(t *testing.T) {
	tmpFilePath := ".temp.test-APIDocsIsOK"
	defer removeTestFile(tmpFilePath) // delete test file after the job is finished

	type CData struct {
		data         string
		expectResult bool
	}

	dOK := `## columns
- **Attribute:** 🐬-🐬
- **Type:** 🐬Array🐬
- **Attribute:** 🐬data-show-button-text🐬
- **Attribute:** 🐬data-show-button🐬
- **Type:** 🐬Boolean🐬
- **Detail:**
`
	dOK = strings.ReplaceAll(dOK, "🐬", "`")
	dataOK := CData{dOK, true}

	dFailed := `## columns
- **Attribute:** 🐬-🐬
- **Attribute:** 🐬show-button-text🐬
- **Attribute:** 🐬data-show-button-text🐬
- **Attribute:** 🐬error-2🐬
- **Type:** 🐬Boolean🐬
- **Detail:**
`
	dFailed = strings.ReplaceAll(dFailed, "🐬", "`")
	dataFailed := CData{dFailed, false}

	testDataSlice := []CData{dataOK, dataFailed}

	for _, curData := range testDataSlice {
		if file, _ := os.Create(tmpFilePath); file != nil {
			_, _ = file.Write([]byte(curData.data))
			_ = file.Close()
			resultState := APIDocsIsOK([]string{file.Name()})
			assert.Equal(t, resultState, curData.expectResult)
		}
	}
}
