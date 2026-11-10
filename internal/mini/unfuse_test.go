//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what TestUnfuse — 합 한 했 해 → 하, 됩 됐 돼 → 되, 켜 → 키, 갔 → 가; an open syllable stays
package mini

import "testing"

func TestUnfuse(t *testing.T) {
	cases := [][2]string{{"합", "하"}, {"한", "하"}, {"했", "하"}, {"해", "하"}, {"됩", "되"}, {"됐", "되"}, {"돼", "되"}, {"켜", "키"}, {"갔", "가"}, {"먹", "머"}, {"가", "가"}}
	for _, c := range cases {
		if got := string(unfuse([]rune(c[0])[0])); got != c[1] {
			t.Errorf("%s: %s", c[0], got)
		}
	}
}
