//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesAngle — with no word table, < and > are read as a sign of a pair beside their letters: <명사>은 gives 열기:1 or the letter, and 닫기:1+보조사 beside + 보조사 and the whole; a > with nothing written onto it is 닫기:1 or the letter; a particle that the closing reading cannot take stays in the whole word
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesAngle(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"<명사>은", []string{"<\t경합(열기:1 / <)", "  <[열기:1]\t성립 1.00", "  <\t성립 1.00", "명사", ">은\t경합(닫기:1 + 보조사:은 / > + 보조사:은)", "  >[닫기:1]+은[보조사:은]\t성립 1.00", "  >+은[보조사:은]\t성립 1.00", "  >은\t성립 0.33"}},
		{"<명사>은과 큰 것", []string{"<\t경합(열기:1 / <)", "명사", ">은과", "  >은+과[부사격:과]", "  >은과", "  >[닫기:1]+은[보조사:은] 고+(아)[맺음:어(합침)]", "  >+은[보조사:은] 고+(아)[맺음:어(합침)]", "큰", "것"}},
		{"<알맞은 제목>.txt로 보관", []string{"<\t경합(열기:1 / <)", "알맞은", "제목", ">\t경합(닫기:1 / >)", "  >[닫기:1]\t성립 1.00", "  >\t성립 1.00", ".txt로", "  .txt[열:영숫자]+로[부사격:로]", "보관"}},
	} {
		lines := r.Lines(c.in, false)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
