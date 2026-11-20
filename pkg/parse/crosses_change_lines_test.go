//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestCrossesChangeLines — end to end, with the words after the Latin letters as words of the word table: A씨가, B군은, PD님이, CEO님이, OO님께, B씨와, C씨가 and A씨 lose the whole new word over the place where the script changes and the word after it stands as a word of its own; so do K팝을, e북을, N잡을 and MC몽의 (any known word after the run), OO내는 (내 a verb too, read both ways) and C양의 (양의 a word too, read both ways); 3분이 keeps the number and its unit, K교수가 and T셔츠를 keep their reading cut in two; git했다 and push했다 lose the whole new word git하, push하 over the place right before 했, the stem's last syllable a piece is fused into, and 하+ㅆ+다 stands after the run
package parse

import (
	"strings"
	"testing"
)

func TestCrossesChangeLines(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	noun := word{kinds: kindNominal, count: 500, by: [6]int{500}}
	for _, w := range []string{"씨", "군", "님", "양", "양의", "팝", "북", "잡", "몽", "교수", "셔츠", "분"} {
		r.ix.lex[w] = noun
	}
	r.ix.lex["내"] = word{kinds: kindNominal | kindVerb, count: 500, by: [6]int{250, 250}}
	for _, c := range []struct {
		in   string
		want []string
		gone []string
	}{
		{"A씨가 왔다.", []string{"\n씨[열쇠:체언]+가[주격:가]\n"}, []string{"A씨"}},
		{"B군은 학생이다.", []string{"\n군[열쇠:체언]+은[보조사:은]\n"}, []string{"B군"}},
		{"PD님이 웃었다.", []string{"\n님[열쇠:체언]+이[주격:이]\n"}, []string{"PD님"}},
		{"CEO님이 왔다.", []string{"\n님[열쇠:체언]+이[주격:이]\n"}, []string{"CEO님"}},
		{"OO님께", []string{"\n님[열쇠:체언]+께[부사격:께]\n"}, []string{"OO님"}},
		{"B씨와 C씨가 만났다.", []string{"\n씨[열쇠:체언]+와[부사격:와]\n"}, []string{"B씨", "C씨"}},
		{"A씨", []string{"\n씨[열쇠:체언]"}, []string{"A씨"}},
		{"K팝을 듣는다.", []string{"\n팝[열쇠:체언]+을[목적격:을]\n"}, []string{"K팝"}},
		{"e북을 읽었다.", []string{"\n북[열쇠:체언]+을[목적격:을]\n"}, []string{"e북"}},
		{"N잡을 한다.", []string{"\n잡[열쇠:체언]+을[목적격:을]\n"}, []string{"N잡"}},
		{"MC몽의 노래", []string{"\n몽[열쇠:체언]+의[관형격:의]\n"}, []string{"MC몽"}},
		{"OO내는 조용하다.", []string{"\n내[열쇠:체언]+는[보조사:는]\n", "\n내[열쇠:동사]+는[꾸밈:는]\n"}, []string{"OO내"}},
		{"C양의 가방", []string{"\n양의[열쇠:체언]\n", "\n양[열쇠:체언]+의[관형격:의]\n"}, []string{"C양"}},
		{"3분이 지났다.", []string{"\n3[수]+분[단위:분]+이[주격:이]\n"}, nil},
		{"K교수가 말했다.", []string{"\n교수[열쇠:체언]+가[주격:가]\n"}, []string{"K교수"}},
		{"T셔츠를 입고 갔다.", []string{"\nT[열:영숫자] 셔츠[열쇠:체언]+를[목적격:를]\n"}, []string{"T셔츠+"}},
		{"파일을 git했다.", []string{"\n하[열쇠:동사]+(였)[선어말:ㅆ]+다[맺음:다]\n"}, []string{"git하"}},
		{"코드를 push했다.", []string{"\n하[열쇠:동사]+(였)[선어말:ㅆ]+다[맺음:다]\n"}, []string{"push하"}},
	} {
		got := streamText(r, c.in)
		for _, w := range c.want {
			if !strings.Contains(got, w) {
				t.Errorf("%s: want %q\n%s", c.in, w, got)
			}
		}
		for _, g := range c.gone {
			if strings.Contains(got, g) {
				t.Errorf("%s: the new word %s stands\n%s", c.in, g, got)
			}
		}
	}
}
