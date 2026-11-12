//ff:func feature=combo type=helper control=sequence
//ff:what TestLine — the whole line: rank, verdict and preference with two decimals, and the candidate picked for each contest parted by " | "; a candidate of two words prints them with a space between
package combo

import (
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestLine(t *testing.T) {
	noun := token.Part{Text: "고양이", Name: token.Name{Kind: "열쇠", Form: "체언"}, Stem: true}
	subject := token.Part{Text: "가", Name: token.Name{Kind: "주격", Form: "가"}}
	unknown := token.Part{Text: "고양", Stem: true}
	suffix := token.Part{Text: "이", Name: token.Name{Kind: "접미", Form: "이"}}
	verb := token.Part{Text: "마시", Name: token.Name{Kind: "열쇠", Form: "동사"}, Stem: true}
	closing := token.Part{Text: "다", Hidden: "ㄴ", Name: token.Name{Kind: "맺음", Form: "ㄴ다"}}
	modifier := token.Part{Hidden: "ㄴ", Name: token.Name{Kind: "꾸밈", Form: "ㄴ"}}
	adverb := token.Part{Text: "다", Name: token.Name{Kind: "열쇠", Form: "부사"}, Stem: true}
	dot := token.Part{Text: ".", Name: token.Name{Kind: "마침표"}}
	choices := []token.Contest{
		{Text: "고양이가", Branches: []token.Branch{
			{Words: []token.Eojeol{{Parts: []token.Part{noun, subject}, End: 4}}},
			{Words: []token.Eojeol{{Parts: []token.Part{unknown, suffix, subject}, End: 4}}},
		}},
		{Text: "마신다", Branches: []token.Branch{
			{Words: []token.Eojeol{{Parts: []token.Part{verb, closing}, End: 3}}},
			{Words: []token.Eojeol{{Parts: []token.Part{verb, modifier}, End: 2}, {Parts: []token.Part{adverb}, End: 3}}},
		}},
		{Text: ".", Branches: []token.Branch{{Words: []token.Eojeol{{Parts: []token.Part{dot}, End: 1}}}}},
	}

	got := Line(3, Combination{Number: 1, Score: Score{Verdict: 0.5714, Pref: -17.958}}, choices)
	want := "3\t성립 0.57 선호 -17.96\t고양이[열쇠:체언]+가[주격:가] | 마시[열쇠:동사]+(ㄴ)[꾸밈:ㄴ] 다[열쇠:부사] | .[마침표]"
	if got != want {
		t.Fatalf("line\n%q\nwant\n%q", got, want)
	}

	got = Line(1, Combination{Number: 2, Score: Score{Verdict: 1, Pref: 0}}, choices)
	want = "1\t성립 1.00 선호 0.00\t고양+이[접미:이]+가[주격:가] | 마시[열쇠:동사]+(ㄴ)다[맺음:ㄴ다] | .[마침표]"
	if got != want {
		t.Fatalf("line\n%q\nwant\n%q", got, want)
	}
}
