//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestMatchA — 먹어서 and 했어 read a whole 어; 봐 and 해서 a fused one; 가 alone is not read but 가+요 is; 먹고서 and 서 are not
package parse

import "testing"

func TestMatchA(t *testing.T) {
	seo := ending{form: "서", lead: leadA}
	bare := ending{form: "", lead: leadA}
	cases := []struct {
		word string
		end  int
		e    ending
		yo   bool
		ok   bool
		stem string
		size int
		mark string
	}{
		{"먹어서", 3, seo, false, true, "먹", 4, "먹+어서"},
		{"했어", 2, bare, false, true, "하", 3, "했+어"},
		{"봐", 1, bare, false, true, "봐", 1, "봐"},
		{"해서", 2, seo, false, true, "하", 3, "해서"},
		{"돼요", 1, bare, true, true, "되", 1, "돼"},
		{"가", 1, bare, false, false, "", 0, ""},
		{"가요", 1, bare, true, true, "가", 1, "가"},
		{"먹고서", 3, seo, false, false, "", 0, ""},
		{"서", 1, seo, false, false, "", 0, ""},
		{"먹어도", 3, seo, false, false, "", 0, ""},
		{"지시해서", 4, seo, false, true, "지시하", 3, "지시+해서"},
		{"가셔서", 3, seo, false, true, "가셔", 3, "가+시+어서"},
	}
	for _, c := range cases {
		got, ok := matchA([]rune(c.word), c.end, c.e, c.yo)
		if ok != c.ok || got.stem != c.stem || got.size != c.size || got.marked != c.mark {
			t.Errorf("%s: %+v %v", c.word, got, ok)
		}
	}
}
