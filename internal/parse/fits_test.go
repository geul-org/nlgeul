//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFits — each condition against a vowel end, a ㄹ end, another consonant end and a Latin letter
package parse

import "testing"

func TestFits(t *testing.T) {
	cases := []struct {
		prev  string
		after int
		want  bool
	}{
		{"가", afterAny, true}, {"각", afterAny, true},
		{"각", afterBatchim, true}, {"가", afterBatchim, false}, {"L", afterBatchim, true},
		{"각", afterHard, true}, {"갈", afterHard, false}, {"가", afterHard, false}, {"L", afterHard, true},
		{"가", afterOpen, true}, {"갈", afterOpen, true}, {"각", afterOpen, false}, {"L", afterOpen, true},
		{"가", afterVowel, true}, {"갈", afterVowel, false}, {"L", afterVowel, true},
	}
	for _, c := range cases {
		if got := fits([]rune(c.prev)[0], c.after); got != c.want {
			t.Errorf("fits(%s, %d) = %v", c.prev, c.after, got)
		}
	}
}
