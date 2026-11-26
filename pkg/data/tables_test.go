//ff:func feature=data type=reader control=iteration dimension=1
//ff:what TestTables — every table held in the program is read and has at least one row (the case exception table may have none: its rows are put in case by case), and the tests know every table held: no table is held that the list lacks, none is listed that is not held
package data

import (
	"io/fs"
	"strings"
	"testing"
)

func TestTables(t *testing.T) {
	cols := map[string]int{
		"case-exceptions": 6, "spacing-words": 2, "closed-words": 2, "bound-nouns": 5, "closing-counts": 2, "closing-pieces": 2,
		"closing-sets": 2, "closings": 5, "contracts": 7, "gated-closings": 4,
		"extensions": 1, "noun-suffixes": 3, "number-heads": 2, "pairs": 4, "particle-counts": 2,
		"particles": 5, "short-closings": 6, "short-particles": 4, "short-slots": 4,
		"signs": 2, "spoken-markers": 3, "top-domains": 1, "typo-closings": 2, "typo-words": 2,
		"unit-suffixes": 2, "units": 2,
		"after-unit-words": 1, "sino-units": 1, "closing-chains": 2, "number-suffixes": 1,
		"polite-slots":      2,
		"closing-particles": 2,
		"quote-markers":     1,
		"plate-letters":     1,
		"plate-regions":     1,
	}
	empty := map[string]bool{"case-exceptions": true}
	entries, err := fs.ReadDir(held, ".")
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
		if rows, err := Rows(name, n); err != nil || (len(rows) == 0 && !empty[name]) {
			t.Errorf("%s: %d rows, %v", e.Name(), len(rows), err)
		}
	}
}
