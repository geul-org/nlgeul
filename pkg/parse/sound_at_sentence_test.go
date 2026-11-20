//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestSoundAtSentence — end to end, by the lines of the stream (streamText): Hangul letters alone in a square or round bracket right after as many Han characters are their sound and read no marker (大棍[대곤] loses 대+곤 of 고는, 官治(관치) loses 관하+치, 傲氣[오기] loses 오+기) and keep their known and new readings; a count of characters that differs, a space before the bracket, a bracket closed by another sign (the round one is then a letter) and a bracket with no Han character before it keep the marker; a suffix and a unit after a number are not turned down so (保障率[보장률] keeps 보장+률, 三年[삼년] keeps 삼+년)
package parse

import (
	"strings"
	"testing"
)

func TestSoundAtSentence(t *testing.T) {
	r := streamReader(t, "words-gloss.tsv")
	for _, c := range []struct {
		sentence string
		gone     []string
		kept     []string
	}{
		{"그는 大棍[대곤]을 맞았다.", []string{"곤[맺음:곤(고는)]"}, []string{"\n대곤\n", "\n대[열쇠:체언] 곤[열쇠:체언]\n", "\n+을[목적격:을]\n"}},
		{"官治(관치)의 폐해가 크다.", []string{"치[맺음:지]"}, []string{"\n관치\n", "\n관[열쇠:체언] 치[열쇠:체언]\n", "\n+의[관형격:의]\n"}},
		{"傲氣[오기]를 부렸다.", []string{"기[이름:기]"}, []string{"\n오기[열쇠:체언]\n"}},
		{"大[대곤]을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"大棍 [대곤]을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"大棍[대곤)을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n", "\n)[닫기:1]\n"}},
		{"[대곤]을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"保障率[보장률]이 낮다.", nil, []string{"\n보장[열쇠:체언]+률[접미:률]\n"}},
		{"三年[삼년] 동안", nil, []string{"\n삼[수]+년[단위:년]\n"}},
	} {
		lines := streamText(r, c.sentence)
		for _, g := range c.gone {
			if strings.Contains(lines, g) {
				t.Errorf("%s: %s stands\n%s", c.sentence, g, lines)
			}
		}
		for _, k := range c.kept {
			if !strings.Contains(lines, k) {
				t.Errorf("%s: want %q\n%s", c.sentence, k, lines)
			}
		}
	}
}
