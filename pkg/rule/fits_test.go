//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestFits — each condition against a vowel end, a ㄹ end, another consonant end and a Latin letter or digit; a non-Hangul letter fits the four batchim conditions but neither AfterLiquid nor AfterHangul
package rule

import "testing"

func TestFits(t *testing.T) {
	cases := []struct {
		prev  rune
		after int
		want  bool
	}{
		{'가', AfterAny, true}, {'각', AfterAny, true}, {'L', AfterAny, true},
		{'각', AfterBatchim, true}, {'가', AfterBatchim, false}, {'L', AfterBatchim, true},
		{'각', AfterHard, true}, {'갈', AfterHard, false}, {'가', AfterHard, false}, {'L', AfterHard, true},
		{'가', AfterOpen, true}, {'갈', AfterOpen, true}, {'각', AfterOpen, false}, {'L', AfterOpen, true},
		{'가', AfterVowel, true}, {'갈', AfterVowel, false}, {'L', AfterVowel, true},
		{'갈', AfterLiquid, true}, {'각', AfterLiquid, false}, {'가', AfterLiquid, false}, {'L', AfterLiquid, false}, {'3', AfterLiquid, false},
		{'가', AfterHangul, true}, {'갈', AfterHangul, true}, {'각', AfterHangul, true}, {'L', AfterHangul, false}, {'3', AfterHangul, false},
	}
	for _, c := range cases {
		if got := Fits(c.prev, c.after); got != c.want {
			t.Errorf("Fits(%c, %d) = %v", c.prev, c.after, got)
		}
	}
}
