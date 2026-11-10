//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestReadChain — 에서+는 is adverbial or subject by 에서; 까지+의 adnominal; 만+을 object; 은 alone has no case; no particle no chain
package mini

import "testing"

func TestReadChain(t *testing.T) {
	cases := []struct {
		word string
		want chainFact
	}{
		{"학교에서는", chainFact{6, "에서+는", "는", 8160000, caseSubj | caseAdv}},
		{"어제까지의", chainFact{6, "까지+의", "의", 8140000, caseAdn}},
		{"일만을", chainFact{4, "만+을", "을", 9690000, caseObj}},
		{"이것은", chainFact{2, "은", "은", 8160000, 0}},
		{"곰이", chainFact{2, "이", "이", 5530000, caseSubj | caseComp}},
		{"철수야", chainFact{2, "야", "야", 7037, caseVoc}},
		{"매우", chainFact{}},
	}
	for _, c := range cases {
		if got := readChain([]rune(c.word)); got != c.want {
			t.Errorf("%s: %+v", c.word, got)
		}
	}
}
