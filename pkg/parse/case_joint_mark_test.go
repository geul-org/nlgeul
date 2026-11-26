//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestCaseJointMark — the letters between the two numbers of a case number's shape are no case mark when they are the particle of a branch number alone (의) or a piece read after a number and then a joint (조의, 분의, 호의, 조제); the prefix of an order alone (제), a joint after letters no table reads after a number (다의), and a case mark (다, 회합) are not such letters
package parse

import "testing"

func TestCaseJointMark(t *testing.T) {
	cases := []struct {
		mark string
		want bool
	}{
		{"의", true},
		{"조의", true},
		{"분의", true},
		{"호의", true},
		{"조제", true},
		{"제", false},
		{"다의", false},
		{"다", false},
		{"회합", false},
	}
	sc := scanOf("")
	for _, c := range cases {
		if got := sc.caseJointMark(c.mark); got != c.want {
			t.Errorf("%s: %v, want %v", c.mark, got, c.want)
		}
	}
}
