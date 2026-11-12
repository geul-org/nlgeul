//ff:func feature=token type=helper control=sequence
//ff:what TestEojeolString — a word is printed as its parts joined with "+"; one part alone has no "+"; a word with no parts is nothing; a word of markers alone that attaches to the unit before begins with "+"
package token

import "testing"

func TestEojeolString(t *testing.T) {
	e := Eojeol{Parts: []Part{
		{Text: "마시", Stem: true, Name: Name{"열쇠", "동사"}},
		{Text: "다", Hidden: "ㄴ", Name: Name{"맺음", "ㄴ다"}},
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
	attach := Eojeol{Parts: []Part{{Text: "로", Name: Name{"부사격", "로"}}}, End: 1, Attach: true}
	if got := attach.String(); got != "+로[부사격:로]" {
		t.Fatalf("attached: %q", got)
	}
	attach.Parts = append(attach.Parts, Part{Text: "는", Name: Name{"보조사", "는"}})
	if got := attach.String(); got != "+로[부사격:로]+는[보조사:는]" {
		t.Fatalf("attached, two parts: %q", got)
	}
}
