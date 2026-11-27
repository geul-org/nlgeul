//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestPartName — the plain number name is made from the letters of the part (수:십진, 수:한글십진, plain 수 for a circled number), any other name is read as nameOf reads it
package parse

import "testing"

func TestPartName(t *testing.T) {
	for _, c := range []struct {
		name, text string
		want       Name
	}{
		{"수", "7", Name{"수", "십진"}}, {"수", "일곱", Name{"수", "한글십진"}}, {"수", "②", Name{"원", "수"}}, {"수", "n", Name{"특수", "로마자"}}, {"수", "몇", Name{"수", ""}}, {plainNumName, "사", Name{"수", ""}},
		{"수:2진법", "101", Name{"수", "2진법"}}, {"단위:개", "개", Name{"단위", "개"}}, {"", "x", Name{}},
	} {
		if got := partName(c.name, c.text); got != c.want {
			t.Errorf("%s %s: %+v, want %+v", c.name, c.text, got, c.want)
		}
	}
}
