//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestShortOf — a line of the shortened closings' table is read into the closing, the form it is short for and the stems it follows; the mark that the copula does not take it moves to the row's own flag; an error of the closing's columns and a stem kind outside the names are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestShortOf(t *testing.T) {
	s, err := shortOf(rowFor("대", "-", "종결", "서술격 안 됨", "다고 해", "형용사"))
	want := short{parse.Marker{Form: "대", Kind: rule.KindFinal}, "다고 해", true, rule.TailAdj}
	if err != nil || s != want {
		t.Fatal(s, err)
	}
	s, err = shortOf(rowFor("ㄴ랄까", "-", "종결", "으 붙음", "ㄴ다고 할까", "동사"))
	want = short{parse.Marker{Form: "랄까", Lead: rule.JongN, Kind: rule.KindFinal, Eu: true}, "ㄴ다고 할까", false, rule.TailVerb}
	if err != nil || s != want {
		t.Fatal(s, err)
	}
	s, err = shortOf(rowFor("래", "-", "종결", "-", "라고 해", "-"))
	if err != nil || s.tail != 0 || s.bars {
		t.Fatal(s, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "종결", "-", "다고 해", "-"}, []string{"1번째 칸(꼴)"}},
		{[]string{"대", "-", "종결", "-", "다고 해", "명사"}, []string{"6번째 칸(여섯째)", "쓸 수 있는 이름", "동사"}},
	} {
		_, err := shortOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
