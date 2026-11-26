//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCutterCandidates — the candidates of a sign at a place for every way the signs before it may stand open: an opening sign opens one deeper, a closing sign closes the nearest of its kind, a quotation mark may open or close, an unsure sign is also its letter, a backtick outside the spans of code opens beside its letter, a < and a > each written apart open and close as a pair when both are and give nothing alone, a sign that no candidate stands for in its word, a letter, a sign in a span of code, a closing bracket with nothing written right before it and one other than a round bracket right after a digit give nothing, a round bracket right after a digit with none open is a list number (닫기:0) beside its letter and the usual closing with one open, and one written onto a word with nothing open closes at loneBracketDepth with the way left as it was; a quotation mark written onto a word with Hangul right after it closes as 닫기:0 with the ways left as they were when none of its kind stands open in any way, and not in a way with none open when one stands open in another way; the cutter is not moved; a trace keeps the lines of the graphs (the usual closing's and the lone one's)
package parse

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
		{"(가)", nil, 2, []pairing{{shutName, loneBracketDepth}}, false, []string{""}},
		{"(1)", nil, 2, []pairing{{shutName, 0}}, true, []string{""}},
		{"(1)", []string{"("}, 2, []pairing{{shutName, 1}}, false, []string{""}},
		{"[1]", nil, 2, nil, false, nil},
		{"( )", nil, 2, nil, false, nil},
		{"(가)", nil, 1, nil, false, nil},
		{"(\"가\")", []string{"(\""}, 3, []pairing{{shutName, 2}}, false, []string{"("}},
		{"(\"가\")", []string{"(\""}, 4, []pairing{{shutName, 1}}, false, nil},
		{"\"가\"", nil, 0, []pairing{{openName, 1}}, false, []string{"\""}},
		{"\"가\"", []string{"\""}, 2, []pairing{{shutName, 1}}, false, []string{""}},
		{"'가'", nil, 0, []pairing{{openName, 1}}, true, []string{"'", ""}},
		{"'가'", []string{"'", ""}, 2, []pairing{{shutName, 1}, {openName, 1}}, true, []string{"", "'"}},
		{"it's", nil, 2, nil, false, nil},
		{"x < y", nil, 2, nil, false, nil},
		{"<가>", nil, 0, []pairing{{openName, 1}}, true, []string{"<", ""}},
		{"`(`", nil, 1, nil, false, nil},
		{"`가", nil, 0, []pairing{{openName, 1}}, true, []string{"`", ""}},
		{"가 `", []string{"("}, 2, []pairing{{openName, 2}}, true, []string{"(`", "("}},
		{"< 가 >", nil, 0, []pairing{{openName, 1}}, true, []string{"<", ""}},
		{"< 가 >", []string{"<", ""}, 4, []pairing{{shutName, 1}}, true, []string{"", "<"}},
		{"가 > 나", nil, 2, nil, false, nil},
		{"좋다\"고", nil, 2, []pairing{{openName, 1}, {shutName, 0}}, false, []string{"\"", ""}},
		{"“좋다\"고", []string{"“"}, 3, []pairing{{openName, 2}, {shutName, 0}}, false, []string{"“\"", "“"}},
		{"\"좋다”고", []string{"\""}, 3, []pairing{{shutName, 0}}, false, []string{"\""}},
		{"“가”고", []string{"“", ""}, 2, []pairing{{shutName, 1}}, true, []string{""}},
		{"“가”고", []string{"(", "(“"}, 2, []pairing{{shutName, 2}}, true, []string{"("}},
		{"'좋다'고", []string{"'", ""}, 3, []pairing{{shutName, 1}, {openName, 1}}, true, []string{"", "'"}},
		{"좋다’고", nil, 2, []pairing{{shutName, 0}}, false, []string{""}},
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
	if _, _, _, notes := cut.candidates(2); len(notes) != 2 || !strings.Contains(notes[0], "닫기:)") || !strings.Contains(notes[0], "통과") || !strings.Contains(notes[1], whyOpenedAny) {
		t.Fatal("the trace of a closing sign", notes)
	}
}
