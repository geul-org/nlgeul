//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestWrittenApartLines — end to end, the result of the written-apart exception (writtenApart) run together with the exception over the place where the script changes (crossesChange, Q11): crossesChange also turns down every whole new word writtenApart turns down here and each expectation is one of TestCrossesChangeLines too, so these lines pass with writtenApart off and do not tell its own work; TestWrittenApart (writtenApart called itself) and TestUnlistedPiece tell it, and TestAddStems the exception it feeds; with titles and dependent nouns of the spacing articles as nouns of the word table: A씨가, B군은, PD님이, CEO님이, OO님께, B씨와, C씨가 and A씨 lose the whole new word over the Latin letters and the title, and the title stands as a word of its own; K팝을, e북을, N잡을, MC몽의 (no such word), OO내는 (내 a verb too), C양의 (양의 a word too) are not written apart, yet lose the whole new word to the exception over the place where the script changes (crossesChange); 3분이 (after a digit) keeps the number and its unit, K교수가 and T셔츠를 keep their reading cut in two
package parse

import (
	"strings"
	"testing"
)

func TestWrittenApartLines(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}
	for _, w := range []string{"씨", "군", "님", "양", "양의", "팝", "북", "잡", "몽", "교수", "셔츠", "분"} {
		r.ix.lex[w] = noun
	}
	r.ix.lex["내"] = word{kinds: kindNominal | kindVerb, count: 500, by: [6]int{250, 250}}
	for _, c := range []struct {
		in, want string
		gone     []string
	}{
		{"A씨가 왔다.", "\n씨[열쇠:체언]+가[주격:가]\n", []string{"A씨"}},
		{"B군은 학생이다.", "\n군[열쇠:체언]+은[보조사:은]\n", []string{"B군"}},
		{"PD님이 웃었다.", "\n님[열쇠:체언]+이[주격:이]\n", []string{"PD님"}},
		{"CEO님이 왔다.", "\n님[열쇠:체언]+이[주격:이]\n", []string{"CEO님"}},
		{"OO님께", "\n님[열쇠:체언]+께[부사격:께]\n", []string{"OO님"}},
		{"B씨와 C씨가 만났다.", "\n씨[열쇠:체언]+와[부사격:와]\n", []string{"B씨", "C씨"}},
		{"A씨", "\n씨[열쇠:체언]", []string{"A씨"}},
		{"K팝을 듣는다.", "\n팝[열쇠:체언]+을[목적격:을]\n", []string{"K팝"}},
		{"e북을 읽었다.", "\n북[열쇠:체언]+을[목적격:을]\n", []string{"e북"}},
		{"N잡을 한다.", "\n잡[열쇠:체언]+을[목적격:을]\n", []string{"N잡"}},
		{"MC몽의 노래", "\n몽[열쇠:체언]+의[관형격:의]\n", []string{"MC몽"}},
		{"OO내는 조용하다.", "\n내[열쇠:체언]+는[보조사:는]\n", []string{"OO내"}},
		{"C양의 가방", "\n양[열쇠:체언]+의[관형격:의]\n", []string{"C양"}},
		{"3분이 지났다.", "\n3[수:십진]+분[단위:분]+이[주격:이]\n", nil},
		{"K교수가 말했다.", "\n교수[열쇠:체언]+가[주격:가]\n", nil},
		{"T셔츠를 입고 갔다.", "\nT[열:영숫자] 셔츠[열쇠:체언]+를[목적격:를]\n", nil},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.want) {
			t.Errorf("%s: want %q\n%s", c.in, c.want, got)
		}
		for _, g := range c.gone {
			if strings.Contains(got, g) {
				t.Errorf("%s: the new word %s stands\n%s", c.in, g, got)
			}
		}
	}
}
