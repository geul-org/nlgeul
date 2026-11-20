//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLabelOf — the kind before the colon (주격 of 주격:이, 꾸밈 of 꾸밈:ㄹ, 열쇠 of 열쇠:체언, 보조 용언 of 보조 용언:있), the name itself with no colon (없음, 쉼표), and "" for ""
package parse

import "testing"

func TestLabelOf(t *testing.T) {
	for name, want := range map[string]string{"주격:이": "주격", "꾸밈:ㄹ": "꾸밈", "열쇠:체언": "열쇠", "보조 용언:있": "보조 용언", "열기:1": "열기", "없음": "없음", "쉼표": "쉼표", "": ""} {
		if got := labelOf(name); got != want {
			t.Errorf("labelOf(%q) = %q, want %q", name, got, want)
		}
	}
}
