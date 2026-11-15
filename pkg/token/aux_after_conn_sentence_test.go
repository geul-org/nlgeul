//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestAuxAfterConnSentence — end to end, over the small word table with the auxiliaries and the main verbs of the sentences as verbs seen enough, and its total made a hundred times as large (so that a known word weighs against a new word about as in a table of real size): 해줘, 해봐, 해놓고, 가보면, 와주었으면, 해둔, 꽂아둔, 읽어봤다 and 만들어줘 are read first as the main verb's 어/아 form and the auxiliary, whose node stands fully (the rebuttal that it is one syllable long does not hold), where 해줘 … 해둔 and 가보는 were read whole as a new word before; 해주에, 가보를, 들어가다 and 내려가는 keep their reading, and in 먹고있다 (고, not 어/아) the node of 있 stays below 1; 가보는 (the known harm) is read first as 가+(아) and 보+는, and the new word 가보 before the particle 는 stays a candidate
package token

import (
	"strings"
	"testing"
)

func TestAuxAfterConnSentence(t *testing.T) {
	r := streamReader(t, "words.tsv")
	verb := word{kinds: kindVerb, count: seenEnough, by: [6]int{0, seenEnough}}
	for _, w := range []string{"주", "보", "놓", "있", "오", "꽂", "읽", "만들", "들어가", "내려가"} {
		r.ix.lex[w] = verb
	}
	r.ix.lex["두"] = word{kinds: kindNominal | kindVerb, count: 300 + seenEnough, by: [6]int{300, seenEnough}}
	r.ix.lex["해주"] = word{kinds: kindNominal, count: 500, by: [6]int{500}}
	r.ix.lexTotal *= 100
	for _, c := range []struct {
		in, want, aux string
		lifted        bool
	}{
		{"이것 좀 해줘.", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 주[보조 용언:주]+(어)[맺음:어(합침)]", "주[보조 용언:주]+(어)[맺음:어(합침)]", true},
		{"다시 해봐.", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 보[보조 용언:보]+(아)[맺음:어(합침)]", "보[보조 용언:보]+(아)[맺음:어(합침)]", true},
		{"일을 해놓고 갔다.", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 놓[보조 용언:놓]+고[맺음:고]", "놓[보조 용언:놓]+고[맺음:고]", true},
		{"가보면 안다.", "가[열쇠:동사]+(아)[맺음:어(합침)] │ 보[보조 용언:보]+면[맺음:면]", "보[보조 용언:보]+면[맺음:면]", true},
		{"와주었으면 좋겠다.", "오[열쇠:동사]+(아)[맺음:어(합침)] │ 주[보조 용언:주]+었[선어말:었]+으면[맺음:으면]", "주[보조 용언:주]+었[선어말:었]+으면[맺음:으면]", true},
		{"해둔 일이 많다.", "하[열쇠:동사]+(여)[맺음:어(합침)] │ 두[보조 용언:두]+(ㄴ)[꾸밈:ㄴ]", "두[보조 용언:두]+(ㄴ)[꾸밈:ㄴ]", true},
		{"꽂아둔", "꽂[열쇠:동사]+아[맺음:아] │ 두[보조 용언:두]+(ㄴ)[꾸밈:ㄴ]", "두[보조 용언:두]+(ㄴ)[꾸밈:ㄴ]", true},
		{"그 책을 읽어봤다.", "읽[열쇠:동사]+어[맺음:어] │ 보[보조 용언:보]+(았)[선어말:ㅆ]+다[맺음:다]", "보[보조 용언:보]+(았)[선어말:ㅆ]+다[맺음:다]", true},
		{"만들어줘.", "만들[열쇠:동사]+어[맺음:어] │ 주[보조 용언:주]+(어)[맺음:어(합침)]", "주[보조 용언:주]+(어)[맺음:어(합침)]", true},
		{"해주에 갔다.", "해주[열쇠:체언]+에[부사격:에]", "", false},
		{"가보를 팔았다.", "가보+를[목적격:를]", "", false},
		{"들어가다.", "들어가[열쇠:동사]+다[맺음:다]", "", false},
		{"내려가는 길", "내려가[열쇠:동사]+는[꾸밈:는]", "", false},
		{"먹고있다.", "먹[열쇠:동사]+고[맺음:고] │ 있[보조 용언:있]+다[맺음:다]", "있[보조 용언:있]+다[맺음:다]", false},
		{"가보는 귀하다.", "가[열쇠:동사]+(아)[맺음:어(합침)] │ 보[보조 용언:보]+는[꾸밈:는]", "보[보조 용언:보]+는[꾸밈:는]", true},
	} {
		var words []string
		aux := -1.0
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
			if b.String() == c.aux {
				aux = b.Verdict
			}
		}
		got := " │ " + strings.Join(words, " │ ") + " │ "
		if !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
		if c.aux != "" && (aux <= 0 || (aux == 1) != c.lifted) {
			t.Errorf("%s: the node %s stands at %v", c.in, c.aux, aux)
		}
	}
	if !netHas(r, "가보는 귀하다.", "가보+는[보조사:는]") {
		t.Fatal("가보는: the new word before the particle is gone")
	}
}
