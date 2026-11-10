//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestHelperStart — an auxiliary verb's stem after an 어/아 form: 해주, 작성해준 (fused), 먹어버리; none in 배우, 주, 먹주
package parse

import "testing"

func TestHelperStart(t *testing.T) {
	cases := []struct {
		stem  string
		fused bool
		want  int
	}{
		{"해주", false, 1}, {"작성해준", true, 3}, {"먹어버리", false, 2}, {"해봐", true, 1},
		{"배우", false, -1}, {"주", false, -1}, {"먹주", false, -1},
	}
	for _, c := range cases {
		if got := helperStart([]rune(c.stem), c.fused); got != c.want {
			t.Errorf("%s: %d", c.stem, got)
		}
	}
}
