//ff:func feature=token type=helper control=sequence
//ff:what TestMidOf — a word of one part named 기호 stands for the sign at its end less one; a word of other parts or names, or at a place with no sign, does not
package token

import "testing"

func TestMidOf(t *testing.T) {
	sc := scanOf("가·나")
	word := Eojeol{Parts: []Part{{Text: "·", Names: []Name{{Kind: midName}}}}, End: 2}
	if m, ok := sc.midOf(word); !ok || len(m.reads) != 1 {
		t.Fatal(m, ok)
	}
	word.End = 3
	if _, ok := sc.midOf(word); ok {
		t.Fatal("no sign there")
	}
	word.End = 2
	word.Parts[0].Names[0].Kind = "열쇠"
	if _, ok := sc.midOf(word); ok {
		t.Fatal("other name")
	}
	if _, ok := sc.midOf(Eojeol{Parts: []Part{{}, {}}, End: 2}); ok {
		t.Fatal("two parts")
	}
}
