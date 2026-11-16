//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderStreamEcho — right on a stem that ends in a syllable with no final consonant, 에 before 요 is rebutted (the letters are also 예요 written 에요) and stays: 고양이에요 and 결과에요 hold the candidate with 에 and 요 at 0.33 next to the closing 예요 written wrong (0.56); a stem with a final consonant (마당에요), 에 not before 요 (결과에는), another particle before 요 (고양이에서요), as it was; 에 after 접미 이 (물이에요) is rebutted too
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamEcho(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
	}{
		{"고양이에요", []string{"\n고양이[열쇠:체언]+에[부사격:에]+요[보조사:요]\t성립 0.33", "\n고양이[열쇠:체언]+예요[맺음:예요,오타정정:에요]\t성립 0.56"}},
		{"결과에요", []string{"\n결과[열쇠:체언]+에[부사격:에]+요[보조사:요]\t성립 0.33", "\n결과[열쇠:체언]+예요[맺음:예요,오타정정:에요]\t성립 0.56"}},
		{"마당에요", []string{"\n마당[열쇠:체언]+에[부사격:에]+요[보조사:요]\t성립 1.00"}},
		{"결과에는", []string{"\n결과[열쇠:체언]+에[부사격:에]+는[보조사:는]\t성립 1.00"}},
		{"고양이에서요", []string{"\n고양이[열쇠:체언]+에서[주격:에서]+요[보조사:요]\t성립 1.00"}},
		{"물이에요", []string{"\n물[열쇠:체언]+이[서술격:이]+에요[맺음:에요]\t성립 1.00", "\n물[열쇠:체언]+이[접미:이]+에[부사격:에]+요[보조사:요]\t성립 0.33"}},
	}
	for _, c := range cases {
		got := scoredText(r, c.sentence)
		for _, want := range c.has {
			if !strings.Contains(got, want) {
				t.Errorf("%s: missing %q in\n%s", c.sentence, want, got)
			}
		}
	}
}
