//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamOrder — the prefix 제 right before a number joins it at a word's head, after a sign and right after a number's end (a digit, a unit after a digit), named 차례 first and 수 second: in the first way of a stream it is 제N[차례] with the unit, the particle or the word after it, and 조 after it is the unit and no greater number; 제 after other letters (문제1, 같은법제3조) and a number with no 제 (36조) are as they were
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamOrder(t *testing.T) {
	r := streamReader(t, "words-order.tsv")
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"이 법 제3조제1항에 따른다.", []string{"제3[차례]+조[단위:조]", "제1[차례]+항[단위:항]+에[부사격:에]"}, []string{"큰 수"}},
		{"제36조에 따라", []string{"제36[차례]+조[단위:조]+에[부사격:에]"}, []string{"큰 수"}},
		{"제2편 제3장 제1절 제2관을 본다.", []string{"제2[차례]+편[단위:편]", "제3[차례]+장[단위:장]", "제1[차례]+절[단위:절]", "제2[차례]+관[단위:관]+을[목적격:을]"}, nil},
		{"제5조의2제3항에 해당하는 경우", []string{"제5[차례]+조[단위:조]+의[관형격:의]", "} 2[수]", "제3[차례]+항[단위:항]+에[부사격:에]"}, []string{"큰 수"}},
		{"갑 제4호증의 기재에", []string{"갑[열쇠:체언]", "제4[차례]+호[단위:호]", "증[열쇠:체언]+의[관형격:의]"}, nil},
		{"제1심 판결을 취소한다.", []string{"} 제1[차례]", "심[열쇠:체언]", "판결[열쇠:체언]+을[목적격:을]"}, nil},
		{"제2차 회의에서", []string{"제2[차례]+차[단위:차]"}, nil},
		{"별지 제3호서식에 따른다.", []string{"제3[차례]+호[단위:호]", "서식[열쇠:체언]+에[부사격:에]"}, nil},
		{"제1항ㆍ제2항 및 제4항에 따라", []string{"제1[차례]+항[단위:항]", "ㆍ[가운뎃점]", "제2[차례]+항[단위:항]", "제4[차례]+항[단위:항]+에[부사격:에]"}, nil},
		{"문제1을 풀었다.", []string{"문제[열쇠:체언]", "1[수]+을[목적격:을]"}, []string{"차례"}},
		{"같은법제3조에 따라", []string{"3+조[큰 수:조]+에[부사격:에]"}, []string{"차례"}},
		{"빚이 36조 원을 넘었다.", []string{"36[수]+조[큰 수:조]"}, []string{"차례"}},
	}
	for _, c := range cases {
		checkLines(t, c.sentence, strings.Join(flatNodes(r, c.sentence, true), "\n")+"\n", c.has, c.lacks)
	}
}
