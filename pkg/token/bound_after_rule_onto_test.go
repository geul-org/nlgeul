//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestBoundAfterRuleOnto — the rule "a dependent noun stands right after a modifying form" (boundAfterRule, bonusBoundAfter) with the dependent noun 거 written onto the modifying form or apart from it, end to end, over the small word table words-adn-bound.tsv (the verb 걸, a stem that loses its ㄹ before ㄴ and ㅂ, the nouns 거지 and 건지 and the words of the sentences, with the counts of the word table) and its total made 23 times as large (about that of the word table): right after a modifying form, apart or written onto it, 겁니다, 건가요, 거지 and 건지 are read first as the dependent noun 거 with the copula dropped and the closing (할 겁니다, 갈 건가요, 먹는거지, 있는건지), which without the rules of a sentence they are not; after a word that is no modifying form 겁니다 is read as it is without the rules (돈을 겁니다: nothing is lifted)
package token

import (
	"strings"
	"testing"
)

func TestBoundAfterRuleOnto(t *testing.T) {
	r := streamReader(t, "words-adn-bound.tsv")
	r.ix.lexTotal *= 23
	cases := []struct{ in, want string }{
		{"내일 할 겁니다.", "거[의존명사:거]+(이)[서술격:이(빠짐)]+(ㅂ)니다[맺음:ㅂ니다]"},
		{"언제 갈 건가요?", "거[의존명사:거]+(이)[서술격:이(빠짐)]+(ㄴ)가[맺음:ㄴ가]+요[보조사:요]"},
		{"먹는거지.", "거[의존명사:거]+(이)[서술격:이(빠짐)]+지[맺음:지]"},
		{"있는건지 몰라.", "거[의존명사:거]+(이)[서술격:이(빠짐)]+(ㄴ)지[맺음:ㄴ지]"},
	}
	for _, c := range cases {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	after := bestNodes(r, "돈을 겁니다.")
	saved := r.ix.sentence
	r.ix.sentence = nil
	defer func() { r.ix.sentence = saved }()
	if bare := bestNodes(r, "돈을 겁니다."); len(after) != len(bare) || after[1].String() != bare[1].String() || after[1].Pref != bare[1].Pref {
		t.Fatal("after a word that is no modifying form nothing is lifted", after, bare)
	}
	for _, c := range cases {
		for _, b := range bestNodes(r, c.in) {
			if b.String() == c.want {
				t.Errorf("%s: %s is first without the rules of a sentence", c.in, c.want)
			}
		}
	}
}
