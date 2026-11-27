//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamOrderLetter — a letter of the order (가, 나, …) before 목 is a number at a word's head and right after a number's end: in the first way of a stream 가[차례:한글]+목[단위:목] stands after 제1[차례]+호 and alone; where the letter and 목 are a word of the table (나목, 거목) that word stands first and the number stays a candidate; with letters running on (다목적) it is no number
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamOrderLetter(t *testing.T) {
	r := streamReader(t, "words-order.tsv")
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"제7조제2항제1호가목에서 정한 기준", []string{"제7[차례]+조[단위:조]", "제2[차례]+항[단위:항]", "제1[차례]+호[단위:호]", "가[차례:한글]+목[단위:목]+에서"}, []string{"[주격:가]"}},
		{"가목에서 정한", []string{"} 가[차례:한글]+목[단위:목]+에서"}, nil},
		{"나목에 따른 금액", []string{"나목[열쇠:체언]+에[부사격:에]"}, []string{"나[수]"}},
		{"거목이 쓰러졌다.", []string{"거목[열쇠:체언]+이"}, []string{"[수:한글십진]"}},
		{"다목적 댐을 지었다.", []string{"다목적[열쇠:체언]"}, []string{"[수:한글십진]", "단위:목"}},
	}
	for _, c := range cases {
		checkLines(t, c.sentence, strings.Join(flatNodes(r, c.sentence, true), "\n")+"\n", c.has, c.lacks)
	}
	checkLines(t, "나목에", streamText(r, "나목에")+"\n", []string{"\n나[차례:한글]+목[단위:목]+에[부사격:에]\n"}, nil)
	checkLines(t, "사목에", streamText(r, "사목에")+"\n", []string{"\n사[차례:한글]+목[단위:목]+에[부사격:에]\n"}, []string{"사[수:한글십진]+목"})
}
