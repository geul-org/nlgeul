//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what useCases — put the rows of a case exception table written in a file (testdata) where the index reads them from, in place of the table held in the program, until the test ends; the file is read as the data package reads a table (the first line names the columns, a line that begins with "# " or is empty is skipped, line numbers count every line)
package parse

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/data"
)

func useCases(t *testing.T, path string) {
	t.Helper()
	text, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(string(text), "\n"), "\n")
	names := strings.Split(strings.TrimPrefix(lines[0], "# 칸(탭으로 나눔): "), ", ")
	var rows []data.Row
	for i, line := range lines {
		if strings.HasPrefix(line, "# ") || line == "" {
			continue
		}
		rows = append(rows, data.Row{File: filepath.Base(path), Line: i + 1, Names: names, Cells: strings.Split(line, "\t")})
	}
	cases, err := loadCases(rows)
	if err != nil {
		t.Fatal(err)
	}
	old := caseExceptions
	caseExceptions = cases
	t.Cleanup(func() { caseExceptions = old })
}
