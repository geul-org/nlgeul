//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestShortSlotOf — a line of the shortened slots' table is read into an entry that stands where the past stands, named with its form and the form it is short for, with the lead its form has and the stems it follows; a form of 어/아, a form with no syllable and a name outside the sets are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestShortSlotOf(t *testing.T) {
	e, err := shortSlotOf(rowFor("댔", "-", "다고 했", "형용사"))
	if err != nil || e.name != "선어말:댔(다고 했)" || e.form != "댔" || e.lead != 0 || e.slot != rule.SlotPast || e.tail != rule.TailAdj || e.after != 0 {
		t.Fatal(e, err)
	}
	e, err = shortSlotOf(rowFor("ㄴ댔", "-", "ㄴ다고 했", "동사"))
	if err != nil || e.name != "선어말:ㄴ댔(ㄴ다고 했)" || e.form != "댔" || e.lead != rule.JongN || e.tail != rule.TailVerb {
		t.Fatal(e, err)
	}
	e, err = shortSlotOf(rowFor("는댔", "받침 뒤", "는다고 했", "-"))
	if err != nil || e.after != parse.AfterNames["받침 뒤"] || e.tail != 0 {
		t.Fatal(e, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "다고 했", "-"}, []string{"1번째 칸(꼴)"}},
		{[]string{"어/아댔", "-", "다고 했", "-"}, []string{"1번째 칸(꼴)", "선어말의 꼴"}},
		{[]string{"", "-", "다고 했", "-"}, []string{"1번째 칸(꼴)", "선어말의 꼴"}},
		{[]string{"댔", "앞", "다고 했", "-"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름"}},
		{[]string{"댔", "-", "다고 했", "명사"}, []string{"4번째 칸(조건)", "형용사"}},
	} {
		_, err := shortSlotOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
