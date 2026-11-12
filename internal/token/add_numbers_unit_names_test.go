//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddNumbersUnitNames — every unit of the unit table is in the index after a number under the name 단위: and the name its column gives, found by the last letter of the unit
package token

import "testing"

func TestAddNumbersUnitNames(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addDots()
	ix.addNumbers()
	for _, u := range unitWords {
		letters := []rune(u)
		found := false
		for _, e := range ix.after[letters[len(letters)-1]] {
			found = found || (e.name == "단위:"+unitNames[u] && e.form == u && e.slot == placeUnit)
		}
		if !found {
			t.Errorf("%s: no entry named 단위:%s", u, unitNames[u])
		}
	}
}
