//ff:func feature=token type=helper control=sequence
//ff:what TestEojeolString — a word is printed as its parts joined with "+"; one part alone has no "+"; a word with no parts is nothing
package token

import "testing"

func TestEojeolString(t *testing.T) {
	e := Eojeol{Parts: []Part{
		{Text: "마시", Stem: true, Names: []Name{{"열쇠", "동사"}}},
		{Text: "다", Hidden: "ㄴ", Names: []Name{{"맺음", "ㄴ다"}}},
	}, End: 3}
	if got := e.String(); got != "마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다]" {
		t.Fatalf("got %q", got)
	}
	if got := (Eojeol{Parts: []Part{{Text: "마신", Stem: true}}}).String(); got != "마신" {
		t.Fatalf("one part: %q", got)
	}
	if got := (Eojeol{}).String(); got != "" {
		t.Fatalf("no parts: %q", got)
	}
}
