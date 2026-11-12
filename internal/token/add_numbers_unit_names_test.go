//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddNumbersUnitNames — every line of the unit table is in the index after a number under the name 단위: and the name of its line, found by the last letter of the unit; a letter on two lines has two entries
package token

import "testing"

func TestAddNumbersUnitNames(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addDots()
	ix.addNumbers()
	for _, row := range unitRows {
		letters := []rune(row.Cells[0])
		found := false
		for _, e := range ix.after[letters[len(letters)-1]] {
			found = found || (e.name == "단위:"+row.Cells[1] && e.form == row.Cells[0] && e.slot == placeUnit)
		}
		if !found {
			t.Errorf("%s: no entry named 단위:%s", row.Cells[0], row.Cells[1])
		}
	}
	names := map[string]bool{}
	for _, e := range ix.after['m'] {
		if e.form == "m" {
			names[e.name] = true
		}
	}
	if len(names) != 2 || !names["단위:미터"] || !names["단위:분"] {
		t.Error(names)
	}
}
