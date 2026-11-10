//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestUnfuse — 합 한 했 해 → 하, 됩 됐 돼 → 되, 켜 → 키, 봐 → 보, 줘 준 → 주, 져 → 지, 갔 → 가; an open syllable stays
package parse

import "testing"

func TestUnfuse(t *testing.T) {
	cases := [][2]string{{"합", "하"}, {"한", "하"}, {"했", "하"}, {"해", "하"}, {"됩", "되"}, {"됐", "되"}, {"돼", "되"}, {"켜", "키"}, {"봐", "보"}, {"줘", "주"}, {"준", "주"}, {"져", "지"}, {"갔", "가"}, {"먹", "머"}, {"가", "가"}}
	for _, c := range cases {
		if got := string(unfuse([]rune(c[0])[0])); got != c[1] {
			t.Errorf("%s: %s", c[0], got)
		}
	}
}
