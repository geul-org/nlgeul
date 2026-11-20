//ff:func feature=parse type=rule control=sequence
//ff:what TestAddNumbers — a number stands with verdict 1; each sign has its candidates (the full stop its shapes of a number first, then the shapes of a date of two numbers with a full stop after them, then its own); the front signs, the graph of 제 before a number, the range sign after a unit, the pieces after a number and the graph of a quotation mark after a digit read as a unit (measure) are there
package parse

import "testing"

func TestAddNumbers(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addDots()
	graphs := ix.graphs
	ix.addNumbers()
	if ix.number != 1 || len(ix.signs['-']) != 5 || ix.signs['/'][2].name != "경로:/" || len(ix.signs['.']) != len(joinNames)+2+len(stopShapes) || ix.signs['.'][0].name != "날짜:." || ix.signs['.'][2].name != "날짜:." || ix.signs['.'][3].name != "날짜:." || ix.signs['.'][4].name != "파일명:." || ix.signs[':'][0].name != "시각::" || len(ix.signs['∼']) != 1 || len(ix.signs['～']) != 1 || len(ix.signs['〜']) != 1 || len(ix.signs['–']) != 4 || ix.signs['–'][1].name != "빼기:–" {
		t.Fatal(ix.number, ix.signs)
	}
	if len(ix.fronts) != 4 || ix.tilde == nil || len(ix.after['개']) != 1 || ix.after['월'][1].name != "단위:개월" || ix.after['째'][0].slot != placeOrder || ix.after['만'][0].name != "큰 수:만" || ix.afterPrefs[true] >= 0 || ix.afterPrefs[false] >= 0 {
		t.Fatal(ix.fronts, ix.after['월'])
	}
	if ix.order == nil || ix.order.name != orderHeadName || ix.measure == nil {
		t.Fatal(ix.order)
	}
	if ix.graphs != graphs+1+19+len(stopShapes)+len(nameSigns)+4+1+1+2+6+len(unitSuffixes)+len(unitRows)+1 || len(ix.after['리']) != 2 {
		t.Fatal(ix.graphs - graphs)
	}
}
