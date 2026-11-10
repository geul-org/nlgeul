//ff:func feature=token type=reader control=sequence
//ff:what TestWordOf — a new word with its marker: the low verdict, the stem and marker printed with the graph name; a closed-set word as stem has verdict 1
package token

import "testing"

func TestWordOf(t *testing.T) {
	sc := scanOf("곰이것")
	e := sc.wordOf(0, 2, tail{start: 1, pieces: []piece{{"이", "주격:이", 1, -1.2}}})
	if e.marked != "곰+이[주격:이]" || e.names != "새 낱말 + 주격:이" || e.verdict != sc.ix.unknown || e.end != 2 {
		t.Fatal(e)
	}
	if e = sc.wordOf(2, 3, tail{start: 3}); e.marked != "것[의존명사:것]" || e.verdict != 1 {
		t.Fatal(e)
	}
}
