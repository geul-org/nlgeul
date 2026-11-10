//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestStemOf — 갔 → 가, 했 → 하, 됐 and 돼 → 되, 분해해 → 분해하; an unfused stem stays
package mini

import "testing"

func TestStemOf(t *testing.T) {
	cases := [][2]string{{"갔", "가"}, {"했", "하"}, {"됐", "되"}, {"돼", "되"}, {"분해해", "분해하"}, {"먹", "먹"}, {"아니", "아니"}}
	for _, c := range cases {
		if got := stemOf([]rune(c[0])); got != c[1] {
			t.Errorf("%s: %s", c[0], got)
		}
	}
}
