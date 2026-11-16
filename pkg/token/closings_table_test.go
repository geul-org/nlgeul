//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestMoreClosingsQuote — the quoting endings the dictionary lists in their shortened forms (냐는, 느냐는, 자는, 냐며, 냐면, 냐면서, 냐니까, 자면, 자니, 자니까, 다니까, 라니까, 다더라 and the verbs' ㄴ/는 forms) are among the closings of the closing table, and a word that ends in one is one node read first as such (왔느냐는, 누구냐면, 말하자면, 좋다더라, 간다니까, 가자니, 좋으냐는); a noun that ends in 자 stays a noun first before 는 (기자는)
package token

import (
	"strings"
	"testing"
)

func TestMoreClosingsQuote(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	r.ix.lex["오"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["좋"] = word{kinds: kindAdj, count: 50, by: [6]int{0, 0, 50}}
	r.ix.lex["말하"] = word{kinds: kindVerb, count: 50, by: [6]int{0, 50}}
	r.ix.lex["누구"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	r.ix.lex["기자"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for _, form := range []string{"냐는", "느냐는", "자는", "냐며", "느냐며", "냐면", "느냐면", "냐면서", "느냐면서", "냐니까", "느냐니까", "자면", "자니", "자니까", "다니까", "ㄴ다니까", "는다니까", "라니까", "다더라", "ㄴ다더라", "는다더라"} {
		if !r.ix.closingSet[form] {
			t.Errorf("%s is no closing", form)
		}
	}
	for in, want := range map[string]string{
		"왔느냐는": "오[열쇠:동사]+(았)[선어말:ㅆ]+느냐는[꾸밈:느냐는]",
		"누구냐면": "누구[열쇠:체언]+(이)[서술격:이(빠짐)]+냐면[맺음:냐면]",
		"말하자면": "말하[열쇠:동사]+자면[맺음:자면]",
		"좋다더라": "좋[열쇠:형용사]+다더라[맺음:다더라]",
		"간다니까": "가[열쇠:동사]+(ㄴ)다니까[맺음:ㄴ다니까]",
		"가자니":  "가[열쇠:동사]+자니[맺음:자니]",
		"좋으냐는": "좋[열쇠:형용사]+으[이음:으]+냐는[꾸밈:냐는]",
		"기자는":  "기자[열쇠:체언]+는[보조사:는]",
	} {
		if got := flatNodes(r, in, true); len(got) != 1 || !strings.HasSuffix(got[0], "} "+want) {
			t.Errorf("%s: want %s first\n%v", in, want, got)
		}
	}
}
