//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTypoClosingOf — a line of the table of closings written wrong is read into the letters it is written with (whole syllables or under a final consonant) and the name of the right closing; a form of 어/아, a form with no syllable and a right closing that is not an index name are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTypoClosingOf(t *testing.T) {
	c, err := typoClosingOf(rowFor("ㄹ께", "맺음:ㄹ게"))
	if err != nil || c != (typoClosing{"께", rule.JongL, "맺음:ㄹ게"}) {
		t.Fatal(c, err)
	}
	c, err = typoClosingOf(rowFor("을께", "맺음:을게"))
	if err != nil || c != (typoClosing{"을께", 0, "맺음:을게"}) {
		t.Fatal(c, err)
	}
	for _, r := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "맺음:ㄹ게"}, []string{"1번째 칸(꼴)"}},
		{[]string{"어/아께", "맺음:ㄹ게"}, []string{"1번째 칸(꼴)", "틀린 꼴"}},
		{[]string{"", "맺음:ㄹ게"}, []string{"1번째 칸(꼴)", "틀린 꼴"}},
		{[]string{"ㄹ께", "ㄹ게"}, []string{"2번째 칸(앞말)", "색인의 이름"}},
	} {
		_, err := typoClosingOf(rowFor(r.row...))
		errHas(t, err, r.parts...)
	}
}
