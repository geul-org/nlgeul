//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNameKind — the part of a name before its colon (꾸밈 of 꾸밈:ㄹ, 열쇠 of 열쇠:체언, 보조 용언 of 보조 용언:있), the name itself with no colon (쉼표), and "" for ""
package token

import "testing"

func TestNameKind(t *testing.T) {
	for name, want := range map[string]string{"꾸밈:ㄹ": "꾸밈", "열쇠:체언": "열쇠", "보조 용언:있": "보조 용언", "열기:1": "열기", "쉼표": "쉼표", "": ""} {
		if got := nameKind(name); got != want {
			t.Errorf("nameKind(%q) = %q, want %q", name, got, want)
		}
	}
}
