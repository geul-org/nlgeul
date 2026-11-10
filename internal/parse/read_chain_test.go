//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReadChain — 에서+는 is adverbial or subject by 에서; 까지+의 adnominal; 만+을 object; 은 alone has no case; 대로 and 따라 adverbial; no particle no chain
package parse

import "testing"

func TestReadChain(t *testing.T) {
	cases := []struct {
		word string
		want chainFact
	}{
		{"학교에서는", chainFact{6, "에서+는", "는", 8160000, caseSubj | caseAdv, "학교+에서+는"}},
		{"어제까지의", chainFact{6, "까지+의", "의", 8140000, caseAdn, "어제+까지+의"}},
		{"일만을", chainFact{4, "만+을", "을", 9690000, caseObj, "일+만+을"}},
		{"이것은", chainFact{2, "은", "은", 8160000, 0, "이것+은"}},
		{"곰이", chainFact{2, "이", "이", 5530000, caseSubj | caseComp, "곰+이"}},
		{"철수야", chainFact{2, "야", "야", 7037, caseVoc, "철수+야"}},
		{"이대로", chainFact{4, "대로", "대로", 17000, caseAdv, "이+대로"}},
		{"오늘따라", chainFact{4, "따라", "따라", 17, caseAdv, "오늘+따라"}},
		{"매우", chainFact{}},
	}
	for _, c := range cases {
		if got := readChain([]rune(c.word)); got != c.want {
			t.Errorf("%s: %+v", c.word, got)
		}
	}
}
