//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestTildeShapes — the tilde and each of its variants join two runs of digits as a range named with the sign, beside the shapes given (the comma's is kept); the shapes of the numbers' signs have them
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTildeShapes(t *testing.T) {
	comma := rule.Shape{Of: "자릿수:,", Min: 2}
	got := tildeShapes(map[rune][]rule.Shape{',': {comma}})
	if len(got) != 1+len([]rune(tildeSigns)) || len(got[',']) != 1 || got[','][0].Of != comma.Of {
		t.Fatal(got)
	}
	for _, r := range tildeSigns {
		want := rule.Shape{Of: "범위:" + string(r), Min: 2, Max: 2}
		if s := got[r]; len(s) != 1 || s[0].Of != want.Of || s[0].Min != 2 || s[0].Max != 2 {
			t.Fatal(string(r), s)
		}
		if s := signShapes[r]; len(s) != 1 || s[0].Of != want.Of {
			t.Fatal("signShapes", string(r), s)
		}
	}
}
