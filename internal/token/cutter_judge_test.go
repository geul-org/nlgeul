//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCutterJudge — the graphs of a sign of a pair for one way the signs before it may stand open: an opening sign reads one deeper than what stands open and adds itself to the way; a closing sign reads as deep as the nearest sign of its kind that stands open and closes what was opened after it; a closing sign with none of its kind open has no candidate; a quotation mark opens when none of its kind stands open and closes when one does; a trace keeps a line for each graph with its reason
package token

import (
	"reflect"
	"strings"
	"testing"
)

func TestCutterJudge(t *testing.T) {
	ix, _ := newIndex(Options{})
	cases := []struct {
		text       string
		at         int
		way        string
		candidates []pairing
		after      []string
	}{
		{"(", 0, "", []pairing{{openName, 1}}, []string{"("}},
		{"(", 0, "{[", []pairing{{openName, 3}}, []string{"{[("}},
		{")", 0, "(", []pairing{{shutName, 1}}, []string{""}},
		{")", 0, "", nil, nil},
		{")", 0, "[", nil, nil},
		{")", 0, "{(\"[", []pairing{{shutName, 2}}, []string{"{"}},
		{"]", 0, "[(", []pairing{{shutName, 1}}, []string{""}},
		{"\"", 0, "", []pairing{{openName, 1}}, []string{"\""}},
		{"\"", 0, "(\"", []pairing{{shutName, 2}}, []string{"("}},
		{"\"", 0, "(", []pairing{{openName, 2}}, []string{"(\""}},
		{"”", 0, "“", []pairing{{shutName, 1}}, []string{""}},
	}
	for _, c := range cases {
		cut := newCutter(ix, c.text, false)
		candidates, after, notes := cut.judge(c.at, pairRows[cut.r[c.at]], cut.facts(c.at), c.way)
		if !reflect.DeepEqual(candidates, c.candidates) || !reflect.DeepEqual(after, c.after) || len(notes) != 0 {
			t.Errorf("%q with %q open: %v after %q notes %q", c.text, c.way, candidates, after, notes)
		}
	}
	cut := newCutter(ix, ")", true)
	_, _, notes := cut.judge(0, pairRows[')'], cut.facts(0), "[")
	if len(notes) != 1 || !strings.Contains(notes[0], "닫기:)") || !strings.Contains(notes[0], "기각: 같은 종류의 여는 기호가 열려 있지 않다") || !strings.Contains(notes[0], `열린 기호 "["`) {
		t.Fatal("the trace of a defeated candidate", notes)
	}
}
