//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPoliteSlotOf — a line of the old polite slots' table is read into a slot named with its form, in the polite slots' place, read only with a word table, with what must stand before it; a form that is no Hangul syllable, a form fused under a syllable, a form that ends in neither 오 nor 옵 and a name outside the sets are errors that name the column
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPoliteSlotOf(t *testing.T) {
	e, err := politeSlotOf(rowFor("오", "모음 뒤"))
	if err != nil || e.name != "선어말:오" || e.form != "오" || e.class != classSlot || e.slot != rule.SlotPolite || e.after != rule.AfterVowel || !e.wide {
		t.Fatal(e, err)
	}
	e, err = politeSlotOf(rowFor("사옵", "받침 뒤"))
	if err != nil || e.name != "선어말:사옵" || e.after != rule.AfterBatchim {
		t.Fatal(e, err)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-"}, []string{"1번째 칸(꼴)"}},
		{[]string{"ㄴ오", "-"}, []string{"1번째 칸(꼴)", "공손의 선어말"}},
		{[]string{"시", "-"}, []string{"1번째 칸(꼴)", "공손의 선어말"}},
		{[]string{"오", "앞"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름"}},
	} {
		_, err := politeSlotOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
