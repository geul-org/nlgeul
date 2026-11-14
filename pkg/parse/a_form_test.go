//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAForm — 먹어, 해, 봐, 가 end in an 어/아 form; 먹, 기 and nothing do not
package parse

import "testing"

func TestAForm(t *testing.T) {
	for _, s := range []string{"먹어", "작성해", "봐", "가"} {
		if !aForm([]rune(s)) {
			t.Errorf("%s", s)
		}
	}
	for _, s := range []string{"먹", "기", ""} {
		if aForm([]rune(s)) {
			t.Errorf("%s", s)
		}
	}
}
