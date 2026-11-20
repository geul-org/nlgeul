//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMoreClosingsOld — the old-fashioned and other endings the dictionary lists that were added at the end of the closing table (ㄴ즉, 은즉, 리오, 노라니, 노라니까, 지마는, 느니, 려는가, 라지) are closings of the table, and a word that ends in one is read first as such (본즉, 어기리오, 어렵지마는); ㄹ새 is read after a predicate's stem of the word table (지낼새) but is no row of the closing table, nor are 리, 리니 and 나마; ㄴ즉 does not stand after the copula (뜻인즉 has the particle 인즉)
package parse

import (
	"strings"
	"testing"
)

func TestMoreClosingsOld(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["보"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["어기"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["어렵"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	r.ix.lex["지내"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["뜻"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for _, form := range []string{"ㄴ즉", "은즉", "리오", "으리오", "노라니", "노라니까", "지마는", "느니", "려는가", "으려는가", "라지", "으라지"} {
		if !r.ix.closingSet[form] {
			t.Errorf("%s is no closing", form)
		}
	}
	for _, form := range []string{"리", "ㄹ새", "리니", "나마"} {
		if r.ix.closingSet[form] {
			t.Errorf("%s is a row of the closing table", form)
		}
	}
	for in, want := range map[string]string{
		"본즉":    "보[열쇠:동사]+(ㄴ)즉[맺음:ㄴ즉]",
		"어기리오":  "어기[열쇠:동사]+리오[맺음:리오]",
		"어렵지마는": "어렵[열쇠:형용사]+지마는[맺음:지마는]",
		"지낼새":   "지내[열쇠:동사]+(ㄹ)새[맺음:ㄹ새]",
	} {
		if got := flatNodes(r, in, true); len(got) != 1 || !strings.HasSuffix(got[0], "} "+want) {
			t.Errorf("%s: want %s first\n%v", in, want, got)
		}
	}
	if got := streamText(r, "뜻인즉"); strings.Contains(got, "이[서술격:이]+(ㄴ)즉") || !strings.Contains(got, "인즉[보조사:인즉]") {
		t.Errorf("뜻인즉\n%s", got)
	}
}
