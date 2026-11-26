//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamQuoteSide — with no word table, a quotation mark is read by where it is written: one at a sentence's start before a marker that takes a quotation closes first, at depth 1, with the opening lowered behind it ("라고, '는), and the marks after it open and close in turn; one written onto a word before a space or a comma closes with none of its kind open (했다' 로, 몰",) and the next one after a space opens; a round mark drawn as a closing one opens where an opening mark is written (”두, ’거래) and one drawn as an opening one closes where a closing mark is written (사항‘이라고, 좋겠다“고), its reading by the shape lowered behind; a round opening mark written onto a word before another word stays an opening only
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamQuoteSide(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
		not  string
	}{
		{"\"라고 그는 말했다.", []string{"\"[닫기:1]\t성립 1.00", "\"[열기:1]\t성립 0.33", "+라고[부사격:라고]\t성립 1.00", "말했다"}, ""},
		{"'는 제목의 글을 \"곰\"이 썼다.", []string{"'[닫기:1]\t성립 1.00", "'\t성립 1.00", "'[열기:1]\t성립 0.33", "+는[보조사:는]", "\"[열기:1]\t성립 1.00", "\"[닫기:1]\t성립 1.00"}, ""},
		{"\"고 하고, 다음은 \"곰이 왔다.", []string{"\"[닫기:1]\t성립 1.00", "\"[열기:1]\t성립 0.33", "+고[맺음:고]", "\"[열기:1]\t성립 1.00", "\"[닫기:1]\t성립 0.33"}, ""},
		{"인사 했다' 로 고쳐줘", []string{"'[닫기:1]\t성립 1.00", "'\t성립 1.00", "'[열기:1]\t성립 0.33"}, ""},
		{"가게 몰\", \"곰이 왔다\"라는 글", []string{"\"[닫기:1]\t성립 1.00", "\"[열기:1]\t성립 0.33", ",[쉼표]", "\"[열기:1]\t성립 1.00", "\"[닫기:1]\t성립 0.33", "\"[닫기:1]\t성립 1.00", "\"[열기:1]\t성립 0.33", "+라는"}, ""},
		{"그는 ”두 곰이 왔다”고 했다.", []string{"”[열기:1]\t성립 1.00", "”[닫기:1]\t성립 1.00", "+고[맺음:고]"}, "닫기:0"},
		{"법의 ’거래 사항‘이라고 본다.", []string{"’[열기:1]\t성립 1.00", "‘[닫기:1]\t성립 1.00", "‘[열기:2]\t성립 0.33", "+이라고[부사격:이라고]"}, ""},
		{"그는 “곰이 좋겠다“고 했다.", []string{"“[열기:1]\t성립 1.00", "“[닫기:1]\t성립 1.00", "“[열기:2]\t성립 0.33", "+고[맺음:고]"}, ""},
		{"그것은‘가’라고 한다.", []string{"‘[열기:1]\t성립 1.00", "’[닫기:1]\t성립 1.00"}, "‘[닫기"},
	} {
		lines := scoredLines(r, c.in)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
		if c.not != "" && strings.Contains(strings.Join(lines, "\n"), c.not) {
			t.Errorf("%q: %q in\n%s", c.in, c.not, strings.Join(lines, "\n"))
		}
	}
}
