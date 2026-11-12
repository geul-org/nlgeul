//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitNames — the unit table gives every line a name; unitWords holds each letter once, and a letter on two lines (m, ′) keeps both lines in unitRows with different names; a Hangul unit is named by its own letters
package token

import (
	"slices"
	"testing"
)

func TestUnitNames(t *testing.T) {
	if len(unitWords) == 0 || len(unitRows) <= len(unitWords) {
		t.Fatal(len(unitWords), len(unitRows))
	}
	names := map[string][]string{}
	for _, row := range unitRows {
		if row.Cells[1] == "" || !slices.Contains(unitWords, row.Cells[0]) {
			t.Errorf("%s: name %q", row.Cells[0], row.Cells[1])
		}
		names[row.Cells[0]] = append(names[row.Cells[0]], row.Cells[1])
	}
	if len(names) != len(unitWords) || !slices.Equal(names["m"], []string{"미터", "분"}) || !slices.Equal(names["개"], []string{"개"}) || !slices.Equal(names["GB"], []string{"기가바이트"}) {
		t.Error(len(names), len(unitWords), names["m"], names["GB"])
	}
}
