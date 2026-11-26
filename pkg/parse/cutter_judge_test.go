//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestCutterJudge — the graphs of a sign of a pair for one way the signs before it may stand open: an opening sign reads one deeper than what stands open and adds itself to the way; a closing sign reads as deep as the nearest sign of its kind that stands open and closes what was opened after it; a closing sign with none of its kind open has no candidate; a quotation mark opens when none of its kind stands open and closes when one does; a quotation mark written onto a word with Hangul right after it and none of its kind open in any way also closes, as 닫기:0, leaving the way as it was (beside the opening of a quotation mark written the same way at both ends); with one of its kind open in some way (MarkOpenedAny) it does not; with a space before it it closes only before a marker that takes a quotation ( ”고, at depth 1), with no letter or digit after it it closes at depth 1 (다”.), but not right after a digit (3".); each candidate holds its graph's verdict, lowered where the sign's place speaks against it (the opening of 다"고); a round quotation mark is read against its shape too where the other mark is written ( ”곰 opens, 다“고 closes); a closing bracket written onto a word with none of its kind open closes too, at loneBracketDepth, leaving the way as it was, and does not with one of its kind open in some way nor right after a digit; a trace keeps a line for each graph with its reason
package parse

import (
	"reflect"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCutterJudge(t *testing.T) {
	ix, _ := newIndex(Options{})
	force := forceQuoteSide
	low := 2/(1+force) - 1
	cases := []struct {
		text       string
		at         int
		way        string
		facts      int
		candidates []pairing
		after      []string
	}{
		{"(", 0, "", 0, []pairing{{openName, 1, 1}}, []string{"("}},
		{"(", 0, "{[", 0, []pairing{{openName, 3, 1}}, []string{"{[("}},
		{")", 0, "(", 0, []pairing{{shutName, 1, 1}}, []string{""}},
		{")", 0, "", 0, nil, nil},
		{")", 0, "[", 0, nil, nil},
		{")", 0, "{(\"[", 0, []pairing{{shutName, 2, 1}}, []string{"{"}},
		{"]", 0, "[(", 0, []pairing{{shutName, 1, 1}}, []string{""}},
		{"\"", 0, "", 0, []pairing{{openName, 1, 1}}, []string{"\""}},
		{"\"", 0, "(\"", 0, []pairing{{shutName, 2, 1}}, []string{"("}},
		{"\"", 0, "(", 0, []pairing{{openName, 2, 1}}, []string{"(\""}},
		{"”", 0, "“", 0, []pairing{{shutName, 1, 1}}, []string{""}},
		{"다\"고", 1, "", 0, []pairing{{openName, 1, low}, {shutName, 0, 1}}, []string{"\"", ""}},
		{"다\"고", 1, "(", 0, []pairing{{openName, 2, low}, {shutName, 0, 1}}, []string{"(\"", "("}},
		{"다”고", 1, "", 0, []pairing{{shutName, 0, 1}}, []string{""}},
		{"다’고", 1, "「", 0, []pairing{{shutName, 0, 1}}, []string{"「"}},
		{"다”고", 1, "“", rule.MarkOpenedAny, []pairing{{shutName, 1, 1}}, []string{""}},
		{"다”고", 1, "(", rule.MarkOpenedAny, nil, nil},
		{" ”고", 1, "", 0, []pairing{{shutName, 1, 1}}, []string{""}},
		{" ”곰", 1, "", 0, []pairing{{openName, 1, 1}}, []string{"“"}},
		{" ”곰", 1, "“", 0, []pairing{{shutName, 1, low}, {openName, 2, 1}}, []string{"", "““"}},
		{"다”.", 1, "", 0, []pairing{{shutName, 1, 1}}, []string{""}},
		{"다“고", 1, "", 0, []pairing{{openName, 1, low}, {shutName, 0, 1}}, []string{"“", ""}},
		{"다“고", 1, "“", 0, []pairing{{openName, 2, low}, {shutName, 1, 1}}, []string{"““", ""}},
		{"다“곰", 1, "", 0, []pairing{{openName, 1, 1}}, []string{"“"}},
		{"3\".", 1, "", 0, []pairing{{openName, 1, low}}, []string{"\""}},
		{"다)", 1, "", 0, []pairing{{shutName, loneBracketDepth, 1}}, []string{""}},
		{"다)", 1, "‘", 0, []pairing{{shutName, loneBracketDepth, 1}}, []string{"‘"}},
		{"다)", 1, "[", rule.MarkOpenedAny, nil, nil},
		{"2)", 1, "", 0, []pairing{{shutName, 0, 1}}, []string{""}},
		{"2)", 1, "(", rule.MarkOpenedAny, []pairing{{shutName, 1, 1}}, []string{""}},
		{"2]", 1, "", 0, nil, nil},
	}
	for _, c := range cases {
		cut := newCutter(ix, c.text, false)
		m := cut.facts(c.at)
		m.Facts |= c.facts
		candidates, after, notes := cut.judge(c.at, pairRows[cut.r[c.at]], m, c.way)
		if !reflect.DeepEqual(candidates, c.candidates) || !reflect.DeepEqual(after, c.after) || len(notes) != 0 {
			t.Errorf("%q with %q open: %v after %q notes %q", c.text, c.way, candidates, after, notes)
		}
	}
	cut := newCutter(ix, ")", true)
	_, _, notes := cut.judge(0, pairRows[')'], cut.facts(0), "[")
	if len(notes) != 2 || !strings.Contains(notes[0], "닫기:)") || !strings.Contains(notes[0], "기각: 같은 종류의 여는 기호가 열려 있지 않다") || !strings.Contains(notes[0], `열린 기호 "["`) || !strings.Contains(notes[1], "기각: "+whyNotOnWord) {
		t.Fatal("the trace of a defeated candidate", notes)
	}
	cut = newCutter(ix, "다”고", true)
	_, _, notes = cut.judge(1, pairRows['”'], cut.facts(1), "")
	if len(notes) != 3 || !strings.Contains(notes[0], "기각: 같은 종류의 여는 기호가 열려 있지 않다") || !strings.Contains(notes[1], "열기:”") || !strings.Contains(notes[1], "기각: "+whyNotOpenSide) || !strings.Contains(notes[2], "닫기:”") || !strings.Contains(notes[2], "통과") {
		t.Fatal("the trace of the usual closing and the lone one", notes)
	}
}
