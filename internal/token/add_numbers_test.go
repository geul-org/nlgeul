//ff:func feature=token type=rule control=sequence
//ff:what TestAddNumbers — a number stands with verdict 1; each sign has its readings (the full stop keeps its own first); the front signs, the range sign after a unit and the pieces after a number are there
package token

import "testing"

func TestAddNumbers(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addDots()
	graphs := ix.graphs
	ix.addNumbers()
	if ix.number != 1 || len(ix.signs['-']) != 5 || ix.signs['/'][2].name != "경로:/" || len(ix.signs['.']) != len(joinNames)+2 || ix.signs['.'][0].name != "파일명:." || ix.signs[':'][0].name != "시각::" {
		t.Fatal(ix.number, ix.signs)
	}
	if len(ix.fronts) != 4 || ix.tilde == nil || len(ix.after['개']) != 1 || ix.after['월'][1].name != "단위:개월" || ix.after['째'][0].slot != placeOrder || ix.after['만'][0].name != "큰 수:만" || ix.afterPrefs[true] >= 0 || ix.afterPrefs[false] >= 0 {
		t.Fatal(ix.fronts, ix.after['월'])
	}
	if ix.graphs != graphs+1+12+len(nameSigns)+4+1+2+6+len(unitSuffixes)+len(unitRows) || len(ix.after['리']) != 2 {
		t.Fatal(ix.graphs - graphs)
	}
}
