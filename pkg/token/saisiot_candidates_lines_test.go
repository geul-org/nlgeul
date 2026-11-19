//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestSaisiotCandidatesLines — end to end with the small word table of 사이시옷 (rows of the words in the sentences): 전셋값, 장맛비, 아랫목, 전깃불, 배춧잎 (ㄴㄴ added), 처갓집 and 머릿결 come first as one nominal bound by 사이시옷 (전깃불+이, not 전깃불이 whole, which stays a candidate below), with the new word of the same letters still a candidate (전셋값+이); 짜릿함, 오붓한 (ㅎ), 아랫층 (ㅊ, a new word first), 할것이라고 (하+(ㄹ) 것+이+라고: 것 뒤의 인용 조사는 thingQuoteRule이 깎는다), 이것밖엔 (이것 밖+엔), 오얏나무 (오얏 나무), 나뭇가지 (the table's word), 귓밥 and 헛기침 (a front of one syllable) have no reading of 사이시옷
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestSaisiotCandidatesLines(t *testing.T) {
	r := streamReader(t, "words-saisiot.tsv")
	cases := []struct {
		text  string
		first []string // the best of a word each; "a|b" when either is (the two of the same lambda)
		also  []string // candidates that are among the lines
	}{
		{"전셋값이 올랐다.", []string{"전셋값[사이시옷:ㅅ]+이[주격:이]|전셋값[사이시옷:ㅅ]+이[보격:이]"}, []string{"전셋값+이[주격:이]"}},
		{"장맛비가 내렸다.", []string{"장맛비[사이시옷:ㅅ]+가[주격:가]|장맛비[사이시옷:ㅅ]+가[보격:가]"}, nil},
		{"아랫목에 앉았다.", []string{"아랫목[사이시옷:ㅅ]+에[부사격:에]"}, nil},
		{"전깃불이 꺼졌다.", []string{"전깃불[사이시옷:ㅅ]+이[주격:이]|전깃불[사이시옷:ㅅ]+이[보격:이]"}, []string{"전깃불이[사이시옷:ㅅ]", "전깃불+이[주격:이]"}},
		{"배춧잎을 씻었다.", []string{"배춧잎[사이시옷:ㅅ]+을[목적격:을]"}, nil},
		{"처갓집에 갔다.", []string{"처갓집[사이시옷:ㅅ]+에[부사격:에]"}, nil},
		{"머릿결이 곱다.", []string{"머릿결[사이시옷:ㅅ]+이[주격:이]|머릿결[사이시옷:ㅅ]+이[보격:이]"}, nil},
		{"짜릿함을 느꼈다.", nil, nil},
		{"오붓한 시간", nil, nil},
		{"아랫층에 산다.", []string{"아랫층+에[부사격:에]"}, nil},
		{"할것이라고 했다.", []string{"하[열쇠:동사]+(ㄹ)[꾸밈:ㄹ]", "것[의존명사:것]+이[서술격:이]+라고[맺음:라고]"}, nil},
		{"이것밖엔 없다.", []string{"이것[열쇠:체언]", "밖[열쇠:체언]+엔[부사격:엔]"}, nil},
		{"오얏나무 아래", []string{"오얏[열쇠:체언]", "나무[열쇠:체언]"}, nil},
		{"나뭇가지가 흔들렸다.", []string{"나뭇가지[열쇠:체언]+가[주격:가]"}, nil},
		{"귓밥을 팠다.", nil, nil},
		{"헛기침을 했다.", nil, nil},
	}
	for _, c := range cases {
		nodes := flatNodes(r, c.text, true)
		lines := streamText(r, c.text)
		best := func(w string) bool {
			return slices.ContainsFunc(strings.Split(w, "|"), func(alt string) bool {
				return slices.ContainsFunc(nodes, func(n string) bool { return strings.HasSuffix(n, "} "+alt) })
			})
		}
		if i := slices.IndexFunc(c.first, func(w string) bool { return !best(w) }); i >= 0 {
			t.Errorf("%s: %s not first\n%v", c.text, c.first[i], nodes)
		}
		if i := slices.IndexFunc(c.also, func(w string) bool { return !strings.Contains(lines, "\n"+w+"\n") }); i >= 0 {
			t.Errorf("%s: no %s\n%s", c.text, c.also[i], lines)
		}
		if want := len(c.first) > 0 && strings.Contains(c.first[0], saisiotName); strings.Contains(lines, saisiotName) != want {
			t.Errorf("%s: a reading of 사이시옷 should be there: %v\n%s", c.text, want, lines)
		}
	}
}
