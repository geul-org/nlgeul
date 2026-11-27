//ff:func feature=parse type=helper control=sequence
//ff:what TestSuffixedNumber — a tail's stem is a number a suffix is read right after where a number ends where the tail starts, the tail's first piece is a suffix of numberSuffixes (사, 용) and the stem is named 수 or 차례 (3사, 제2사, 4용); not with no piece, where no number ends (3a사), with a suffix not in the table (률) or a stem of another name
package parse

import "testing"

func TestSuffixedNumber(t *testing.T) {
	ix, err := newIndex(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	sc := &scan{ix: ix, s: []rune("3사 3a사 3률 4용")}
	sc.prepare()
	named := func(kind string) step {
		return step{word: Eojeol{Parts: []Part{{Text: "3", Stem: true, Name: Name{Kind: kind}}}}}
	}
	sa := []piece{{text: "사", names: "접미:사"}}
	ryul := []piece{{text: "률", names: "접미:률"}}
	if !sc.suffixedNumber(tail{start: 1, pieces: sa}, named("수")) || !sc.suffixedNumber(tail{start: 1, pieces: sa}, named("차례")) {
		t.Fatal("a number and 사 right after it")
	}
	yong := []piece{{text: "용", names: "접미:용"}}
	if !sc.suffixedNumber(tail{start: 11, pieces: yong}, named("수")) {
		t.Fatal("a number and 용 right after it")
	}
	if sc.suffixedNumber(tail{start: 1}, named("수")) {
		t.Fatal("a tail with no piece")
	}
	if sc.suffixedNumber(tail{start: 6, pieces: sa}, named("수")) {
		t.Fatal("no number ends there")
	}
	if sc.suffixedNumber(tail{start: 1, pieces: ryul}, named("수")) {
		t.Fatal("률 is not read right after a number")
	}
	if sc.suffixedNumber(tail{start: 1, pieces: sa}, named("열쇠")) {
		t.Fatal("a stem that is no number")
	}
}
