//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamLoneShut — with no word table, a quotation mark written onto a word with Hangul right after it and none of its kind open closes as 닫기:0 (beside 열기 for " and '), and the particle or ending on it is read as one (+고, +며, +에, +라고), whichever kind of quotation mark opened before it (“…", "…”); a quotation mark written onto the word before an opening one gives 열기:1 and 닫기:0 alike; a quotation mark right after a digit gets 닫기:0 beside its units; a pair that closes as it did, a quotation mark after a full stop, an apostrophe inside a word and single quotation marks give no 닫기:0
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamLoneShut(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		lone bool
		want []string
	}{
		{"그는 좋다\"고 말했다.", true, []string{"그는", "좋다", "\"[열기:1]\t성립 1.00", "\"[닫기:0]\t성립 1.00", "+고[맺음:고]\t성립 1.00", "말했다"}},
		{"그는 있다”며 웃었다.", true, []string{"그는", "있[보조 용언:있]+다[맺음:다]", "”[닫기:0]", "+며[맺음:며]\t성립 1.00", "웃었다"}},
		{"그는 좋다’고 했다.", true, []string{"좋다", "’[닫기:0]", "+고[맺음:고]\t성립 1.00", "했다"}},
		{"경우”에 따른다.", true, []string{"경우", "”[닫기:0]", "+에[부사격:에]\t성립 1.00", "따른다"}},
		{"그는 “좋다\"고 했다.", true, []string{"“[열기:1]", "좋다", "\"[열기:2]", "\"[닫기:0]", "+고[맺음:고]\t성립 1.00"}},
		{"그는 \"좋다”고 했다.", true, []string{"\"[열기:1]", "좋다", "”[닫기:0]", "+고[맺음:고]\t성립 1.00"}},
		{"그는 좋다\"라고 했다.", true, []string{"좋다", "\"[열기:1]", "\"[닫기:0]", "+라고[맺음:라고]\t성립 1.00"}},
		{"그는\"좋다\"고 했다.", true, []string{"그는", "좋다", "\"[닫기:1]", "\"[열기:1]"}},
		{"키가 5'10\"이다.", true, []string{"5'10", "\"[단위:인치]", "\"[단위:초]", "\"[열기:1]", "\"[닫기:0]", "+이[서술격:이]+다[맺음:다]\t성립 1.00"}},
		{"그는 \"좋다\"고 했다.", false, []string{"그는", "\"[열기:1]", "좋다", "\"[닫기:1]", "+고[맺음:고]\t성립 1.00", "했다"}},
		{"\"좋다\"고 했다.", false, []string{"\"[열기:1]", "좋다", "\"[닫기:1]", "+고[맺음:고]"}},
		{"그는 좋다.\" 하고 나갔다.", false, []string{"좋다", ".[마침표]", "\"[열기:1]", "하고", "나갔다"}},
		{"it's 좋다.", false, []string{"it's", "좋다"}},
		{"'좋다'고 했다.", false, []string{"좋다", "'[닫기:1]", "'[열기:1]", "'\t성립 1.00", "+고[맺음:고]"}},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		if !c.lone && strings.Contains(strings.Join(lines, "\n"), "닫기:0") {
			t.Errorf("%q: a quotation mark closes with none of its kind open, in\n%s", c.in, strings.Join(lines, "\n"))
		}
	}
}
