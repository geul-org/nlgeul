//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestShortParticleOf — a line of the shortened particles' table is read into an entry named with its case, its form and the form it is short for, in the place of a topic particle; a form that is not whole Hangul syllables, a name outside the sets are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestShortParticleOf(t *testing.T) {
	e, err := shortParticleOf(rowFor("론", "모음·ㄹ 뒤", "부사격", "로는"))
	if err != nil || e.name != "부사격:론(로는)" || e.form != "론" || e.class != classParticle || e.after != parse.AfterNames["모음·ㄹ 뒤"] || e.slot != rule.LinkTopic || e.kase != parse.CaseNames["부사격"] {
		t.Fatal(e, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "부사격", "에서는"}, []string{"1번째 칸(꼴)", "한글 음절"}},
		{[]string{"ㄴ선", "-", "부사격", "에서는"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"", "-", "부사격", "에서는"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"선", "앞", "부사격", "에서는"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름"}},
		{[]string{"선", "-", "처격", "에서는"}, []string{"3번째 칸(갈래)", "관형격"}},
	} {
		_, err := shortParticleOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
