//ff:func feature=data type=helper control=iteration dimension=1
//ff:what useTables — put the tables given (name to text) where the package reads its tables from, in place of the ones held in the program, until the test ends
package data

import (
	"testing"
	"testing/fstest"
)

func useTables(t *testing.T, files map[string]string) {
	t.Helper()
	fsys := fstest.MapFS{}
	for name, text := range files {
		fsys["tables/"+name+".tsv"] = &fstest.MapFile{Data: []byte(text)}
	}
	old := tables
	tables = fsys
	t.Cleanup(func() { tables = old })
}
