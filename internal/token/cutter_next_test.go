//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutterNext — the unit at a place by what stands there: plain letters up to the next sign of a pair, an opening sign alone (the ways of standing open move past it), a span of code with what is written onto it, a closing sign with the particle written onto it (the ways move back), a sign read as nothing but a letter in its word, a sign that is also its letter; the first unit of an input line and a space right after a unit are told
package token

import (
	"reflect"
	"testing"
)

func TestCutterNext(t *testing.T) {
	ix, _ := newIndex(Options{})
	c := newCutter(ix, "곰(`a`를)이\n'가' x < y", false)
	steps := []struct {
		at   int
		want unit
		end  int
		open []string
	}{
		{0, unit{text: "곰", first: true}, 1, []string{""}},
		{1, unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}}}, 2, []string{"("}},
		{2, unit{text: "`a`를", head: 3, name: codeName}, 6, []string{"("}},
		{6, unit{text: ")이", head: 1, name: pairName, candidates: []pairing{{shutName, 1}}, spaced: true}, 8, []string{""}},
		{9, unit{text: "'", head: 1, name: pairName, candidates: []pairing{{openName, 1}}, bare: true, first: true}, 10, []string{"'", ""}},
		{10, unit{text: "가", spaced: false}, 11, []string{"'", ""}},
		{11, unit{text: "'", head: 1, name: pairName, candidates: []pairing{{openName, 1}, {shutName, 1}}, bare: true, spaced: true}, 12, []string{"", "'"}},
		{13, unit{text: "x", spaced: true}, 14, []string{"", "'"}},
		{15, unit{text: "<", spaced: true}, 16, []string{"", "'"}},
	}
	for _, s := range steps {
		got, end := c.next(s.at)
		if !reflect.DeepEqual(got, s.want) || end != s.end || !reflect.DeepEqual(c.open, s.open) {
			t.Errorf("next(%d): %+v end %d open %q\nwant %+v end %d open %q", s.at, got, end, c.open, s.want, s.end, s.open)
		}
	}
}
