//ff:func feature=mini type=reader control=sequence
//ff:what TestClosingName — the lead is written out: ㄴ다, ㄹ까, ㅂ니다, 어서; a whole-syllable closing is its form
package mini

import "testing"

func TestClosingName(t *testing.T) {
	if closingName(ending{form: "다", lead: jongN}) != "ㄴ다" || closingName(ending{form: "까", lead: jongL}) != "ㄹ까" || closingName(ending{form: "니다", lead: jongB}) != "ㅂ니다" || closingName(ending{form: "서", lead: leadA}) != "어서" || closingName(ending{form: "지만"}) != "지만" {
		t.Fatal("closingName")
	}
}
