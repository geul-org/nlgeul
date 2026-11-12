//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestCutterCandidates — the candidates of a sign at a place for every way the signs before it may stand open: an opening sign opens one deeper, a closing sign closes the nearest of its kind, a quotation mark may open or close, an unsure sign is also its letter, a sign that no candidate stands for in its word, a letter, a sign in a span of code and a closing sign with nothing open give nothing; the cutter is not moved; a trace keeps the lines of the graphs
package token

import (
	"reflect"
	"strings"
	"testing"
)

func TestCutterCandidates(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text       string
		open       []string
		at         int
		candidates []pairing
		bare       bool
		after      []string
	}{
		{"(가)", nil, 0, []pairing{{openName, 1}}, false, []string{"("}},
		{"(가)", []string{"("}, 2, []pairing{{shutName, 1}}, false, []string{""}},
		{"(가)", nil, 2, nil, false, nil},
		{"(가)", nil, 1, nil, false, nil},
		{"(\"가\")", []string{"(\""}, 3, []pairing{{shutName, 2}}, false, []string{"("}},
		{"(\"가\")", []string{"(\""}, 4, []pairing{{shutName, 1}}, false, nil},
		{"\"가\"", nil, 0, []pairing{{openName, 1}}, false, []string{"\""}},
		{"\"가\"", []string{"\""}, 2, []pairing{{shutName, 1}}, false, []string{""}},
		{"'가'", nil, 0, []pairing{{openName, 1}}, true, []string{"'", ""}},
		{"'가'", []string{"'", ""}, 2, []pairing{{openName, 1}, {shutName, 1}}, true, []string{"", "'"}},
		{"it's", nil, 2, nil, false, nil},
		{"x < y", nil, 2, nil, false, nil},
		{"<가>", nil, 0, []pairing{{openName, 1}}, true, []string{"<", ""}},
		{"`(`", nil, 1, nil, false, nil},
	}
	for _, c := range cases {
		cut := newCutter(ix, c.text, false)
		if c.open != nil {
			cut.open = c.open
		}
		before := append([]string{}, cut.open...)
		candidates, bare, after, notes := cut.candidates(c.at)
		if !reflect.DeepEqual(candidates, c.candidates) || bare != c.bare || (c.after != nil && !reflect.DeepEqual(after, c.after)) || len(notes) != 0 || !reflect.DeepEqual(cut.open, before) {
			t.Errorf("%q at %d with %q: %v bare %v open after %q notes %q", c.text, c.at, before, candidates, bare, after, notes)
		}
	}
	cut := newCutter(ix, "(가)", true)
	cut.open = []string{"("}
	if _, _, _, notes := cut.candidates(2); len(notes) != 1 || !strings.Contains(notes[0], "닫기:)") || !strings.Contains(notes[0], "통과") {
		t.Fatal("the trace of a closing sign", notes)
	}
}
