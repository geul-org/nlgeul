//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMatchJamo — 분해한다, 갑니다, 된다 read a fused consonant; 학생입니다 and 학생인 read the copula; 보이는 꼴 without it; 큰 and words without the consonant give nothing
package parse

import "testing"

func TestMatchJamo(t *testing.T) {
	nda := ending{form: "다", lead: jongN, flags: fNoIda}
	bnida := ending{form: "니다", lead: jongB}
	bare := ending{form: "", lead: jongN, count: 5}
	cases := []struct {
		word   string
		e      ending
		ok     bool
		stem   string
		size   int
		copula bool
		marked string
	}{
		{"분해한다", nda, true, "분해하", 3, false, "분해+한+다"},
		{"된다", nda, true, "되", 3, false, "된+다"},
		{"보인다", nda, true, "보이", 3, false, "보인+다"},
		{"갑니다", bnida, true, "가", 5, false, "갑+니다"},
		{"학생입니다", bnida, true, "학생이", 6, true, "학생+입+니다"},
		{"입니다", bnida, true, "이", 5, false, "입+니다"},
		{"학생인", bare, true, "학생이", 2, true, "학생+인"},
		{"분해한", bare, true, "분해하", 1, false, "분해+한"},
		{"큰", bare, false, "", 0, false, ""},
		{"좋았니다", bnida, false, "", 0, false, ""},
		{"니다", bnida, false, "", 0, false, ""},
		{"예쁘다", bnida, false, "", 0, false, ""},
	}
	for _, c := range cases {
		s := []rune(c.word)
		got, ok := matchJamo(s, len(s), c.e)
		if ok != c.ok || got.stem != c.stem || got.size != c.size || got.copula != c.copula || got.marked != c.marked {
			t.Errorf("%s: %+v %v", c.word, got, ok)
		}
	}
}
