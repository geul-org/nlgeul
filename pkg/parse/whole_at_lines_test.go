//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestWholeAtLines — end to end, a unit 제 + a number joined by a hyphen with no marker after it (제1-1 before 서식, 제2-3 alone, 제1-2 after 별지) is not lost to the exception over the place where the script changes between 제 and the number: its atom's names 차례, 범위 and 빼기 stand, in the lines and in the first way of a stream, as they do with a particle after it (제1-1을)
package parse

import (
	"strings"
	"testing"
)

func TestWholeAtLines(t *testing.T) {
	r := streamReader(t, "words-order.tsv")
	cases := []struct {
		sentence string
		lines    []string
		stream   []string
	}{
		{"제1-1 서식", []string{"\n제1-1[차례]\n", "\n제1-1[범위]\n", "\n제1-1[빼기]\n", "\n서식[열쇠:체언]\n"}, []string{"{0 4} 제1-1[차례]\n", "{4 6} 서식[열쇠:체언]\n"}},
		{"제2-3", []string{"\n제2-3[차례]\n", "\n제2-3[범위]\n", "\n제2-3[빼기]\n"}, []string{"{0 4} 제2-3[차례]\n"}},
		{"별지 제1-2 서식에 따른다.", []string{"\n별지[열쇠:체언]\n", "\n제1-2[차례]\n", "\n제1-2[범위]\n", "서식[열쇠:체언]+에[부사격:에]"}, []string{"{2 6} 제1-2[차례]\n", "{6 9} 서식[열쇠:체언]+에[부사격:에]\n"}},
		{"제1-1을 본다", []string{"\n제1-1[차례]+을[목적격:을]\n", "\n제1-1[범위]+을[목적격:을]\n"}, []string{"{0 5} 제1-1[차례]+을[목적격:을]\n"}},
	}
	for _, c := range cases {
		checkLines(t, c.sentence, streamText(r, c.sentence), c.lines, nil)
		checkLines(t, c.sentence, strings.Join(flatNodes(r, c.sentence, true), "\n")+"\n", c.stream, nil)
	}
}
