//ff:func feature=data type=reader control=iteration dimension=1
//ff:what TestTables — every table held in the program is read and has at least one row, and the tests know every table held: no table is held that the list lacks, none is listed that is not held
package data

import (
	"io/fs"
	"strings"
	"testing"
)

func TestTables(t *testing.T) {
	cols := map[string]int{
		"closed-words-more": 2, "closed-words": 3, "closing-counts": 2, "closing-pieces": 2,
		"closing-sets": 2, "closings-more": 4, "closings": 5, "contracts": 7,
		"extensions": 1, "number-heads": 2, "pairs": 4, "particle-counts": 2,
		"particles": 5, "short-closings": 6, "short-particles": 4, "short-slots": 4,
		"signs": 2, "top-domains": 1, "typo-closings": 2, "typo-words": 2,
		"unit-suffixes": 1, "units": 2,
	}
	entries, err := fs.ReadDir(held, "tables")
	if err != nil || len(entries) != len(cols) {
		t.Fatal(len(entries), len(cols), err)
	}
	for _, e := range entries {
		name := strings.TrimSuffix(e.Name(), ".tsv")
		n, ok := cols[name]
		if !ok {
			t.Errorf("%s: not in the list", e.Name())
			continue
		}
		if rows, err := Rows(name, n); err != nil || len(rows) == 0 {
			t.Errorf("%s: %d rows, %v", e.Name(), len(rows), err)
		}
	}
}
