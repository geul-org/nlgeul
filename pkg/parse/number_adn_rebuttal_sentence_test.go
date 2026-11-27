//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestNumberAdnRebuttalSentence — end to end over a small word table with 인 as a noun: 1인 and 300인 right after a number are read first as the number and its unit 인 (1+인, the unit table's 단위:인; before the unit was in the table, as the number and the noun 인), where the copula's modifying form came first before, and the copula stays a candidate; 84.2%인 (a unit between) and 3이다 (another closing) keep the copula where it stands
package parse

import (
	"strings"
	"testing"
)

func TestNumberAdnRebuttalSentence(t *testing.T) {
	r := streamReader(t, "words-number-adn.tsv")
	for _, c := range []struct{ in, want string }{
		{"그 1인 회사의 주주", "1[수:십진]+인[단위:인]"},
		{"300인 이상", "300[수:십진]+인[단위:인]"},
		{"84.2%인 비율", "84.2[소수점]+%[단위:퍼센트]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]"},
		{"주주는 3이다.", "3[수:십진]+이[서술격:이]+다[맺음:다]"},
	} {
		var words []string
		for _, b := range bestNodes(r, c.in) {
			words = append(words, b.String())
		}
		if got := " │ " + strings.Join(words, " │ ") + " │ "; !strings.Contains(got, " │ "+c.want+" │ ") {
			t.Errorf("%s: want %s first\n%s", c.in, c.want, got)
		}
	}
	if !netHas(r, "300인 이상", "300[수:십진]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]") {
		t.Fatal("the copula right after 300 is gone")
	}
}
