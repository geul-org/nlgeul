//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestBoundAfterAdnSentence — end to end, over the small word table with the predicates of the sentences and the verb 걸 (a stem that loses its ㄹ before ㄴ) and the noun 건지 added, and its total made a hundred times as large: the dependent noun 거 or 것 written right onto a modifying closing is read first as the modifying form and 거 (것) with its particle or the copula, and its node stands fully (the rebuttal that it is one syllable long does not hold): 먹는게, 큰건, 갈거라, 읽는건가, 할것이다; 거 at a unit's head (거기에) and 수 after a modifying closing (할수) keep their reading
package parse

import (
	"strings"
	"testing"
)

func TestBoundAfterAdnSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	enough := 10000
	verb := word{kinds: kindVerb, count: enough, by: [6]int{0, enough}}
	adj := word{kinds: kindAdj, count: enough, by: [6]int{0, 0, enough}}
	for _, w := range []string{"읽", "보", "오"} {
		r.ix.lex[w] = verb
	}
	for _, w := range []string{"크", "좋"} {
		r.ix.lex[w] = adj
	}
	r.ix.lex["걸"] = word{kinds: kindVerb, count: 5000, by: [6]int{0, 5000}, irr: irrL}
	r.ix.lex["건지"] = word{kinds: kindNominal, count: 70, by: [6]int{70}}
	r.ix.lex["거기"] = word{kinds: kindNominal, count: 9000, by: [6]int{9000}}
	r.ix.lexTotal *= 100
	for _, c := range []struct{ in, want, bound string }{
		{"먹는게 좋다.", "먹[열쇠:동사]+는[꾸밈:는] │ 거[의존명사:거]+(이)[주격:이]", "거[의존명사:거]+(이)[주격:이]"},
		{"작은건 좋다.", "작[열쇠:형용사]+은[꾸밈:은] │ 거[의존명사:거]+(는)[보조사:는]", "거[의존명사:거]+(는)[보조사:는]"},
		{"좋을거라 본다.", "좋[열쇠:형용사]+을[꾸밈:을] │ 거[의존명사:거]+(이)[서술격:이(빠짐)]+라[맺음:라]", "거[의존명사:거]+(이)[서술격:이(빠짐)]+라[맺음:라]"},
		{"읽는건가?", "읽[열쇠:동사]+는[꾸밈:는] │ 거[의존명사:거]+(이)[서술격:이(빠짐)]+(ㄴ)가[맺음:ㄴ가]", "거[의존명사:거]+(이)[서술격:이(빠짐)]+(ㄴ)가[맺음:ㄴ가]"},
		{"할것이다.", "하[열쇠:동사]+(ㄹ)[꾸밈:ㄹ] │ 것[의존명사:것]+이[서술격:이]+다[맺음:다]", "것[의존명사:것]+이[서술격:이]+다[맺음:다]"},
		{"거기에 갔다.", "거기[열쇠:체언]+에[부사격:에]", ""},
	} {
		var words []string
		bound := -1.0
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
			if b.String() == c.bound {
				bound = b.Verdict
			}
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
		if c.bound != "" && bound != 1 {
			t.Errorf("%s: the node %s stands at %v", c.in, c.bound, bound)
		}
	}
}
