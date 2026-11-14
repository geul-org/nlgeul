//ff:func feature=token type=reader control=sequence
//ff:what TestRunsOf — the runs of Latin letters, digits and underscores, also at the unit's end; none in Hangul
package token

import "testing"

func TestRunsOf(t *testing.T) {
	got := runsOf([]rune("task.py에 my_var2"))
	if len(got) != 3 || got[0] != [2]int{0, 4} || got[1] != [2]int{5, 7} || got[2] != [2]int{9, 16} || runsOf([]rune("한글")) != nil {
		t.Fatal(got)
	}
}
