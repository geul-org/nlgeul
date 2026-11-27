//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestSignPartName — a sign's name is kept whole as the kind of the part, except the plain number name, which is made from the letters (수:십진 for -90, plain 수 for n)
package parse

import "testing"

func TestSignPartName(t *testing.T) {
	for _, c := range []struct {
		name, text string
		want       Name
	}{
		{"수", "-90", Name{"수", "십진"}}, {"수", "n", Name{"특수", "로마자"}}, {"수", "⑴", Name{"괄호", "수"}}, {"수", "ⓛ", Name{"수", ""}}, {"범위:~(단위 뒤)", "~", Name{Kind: "범위:~(단위 뒤)"}}, {"", "-", Name{}},
	} {
		if got := signPartName(c.name, c.text); got != c.want {
			t.Errorf("%s %s: %+v, want %+v", c.name, c.text, got, c.want)
		}
	}
}
