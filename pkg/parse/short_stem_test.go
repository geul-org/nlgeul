//ff:func feature=parse type=helper control=sequence
//ff:what TestShortStem — a stem is too short where the tail's first piece is a suffix of longSuffixes (용, 제, 권) and the stem is one syllable (법+용); not with a stem of two syllables (사업+용), with no piece, or with a suffix that is not in the table (사, 별: it is read after one syllable too)
package parse

import "testing"

func TestShortStem(t *testing.T) {
	yong := tail{pieces: []piece{{text: "용", names: "접미:용"}}}
	sa := tail{pieces: []piece{{text: "사", names: "접미:사"}}}
	if !shortStem(yong, "법") {
		t.Fatal("a stem of one syllable")
	}
	if shortStem(yong, "사업") {
		t.Fatal("a stem of two syllables")
	}
	if !shortStem(tail{pieces: []piece{{text: "권", names: "접미:권"}}}, "법") || !shortStem(tail{pieces: []piece{{text: "제", names: "접미:제"}}}, "법") {
		t.Fatal("a stem of one syllable is too short for 권 and 제 too")
	}
	if shortStem(tail{}, "법") {
		t.Fatal("no piece")
	}
	if shortStem(sa, "법") || shortStem(tail{pieces: []piece{{text: "별", names: "접미:별"}}}, "시") {
		t.Fatal("a suffix that is not in the table")
	}
}
