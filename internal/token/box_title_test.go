//ff:func feature=token type=helper control=sequence
//ff:what TestBoxTitle — the place, the letters and the number of all the candidates of the contest
package token

import "testing"

func TestBoxTitle(t *testing.T) {
	n := boxNode(0, 3, "아버지", "아버지", "아버")
	if got := boxTitle(n); got != "[0-3] 아버지 ─ 후보경합 2" {
		t.Fatal(got)
	}
}
