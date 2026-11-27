//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestValueEntries — a sign the sign table does not name is read through the one letter it writes by that letter's named candidates (？: 물음표; ＃: 제목, 샵, 태그; ＝ writes a sign with no name: none); a sign the sign table names itself (，), and a sign that writes more than one letter or a letter with no candidate (㎡, ㈜), have none
package parse

import (
	"reflect"
	"testing"
)

func TestValueEntries(t *testing.T) {
	ix := &index{}
	ix.addDots()
	ix.addSigns()
	for _, c := range []struct {
		f    form
		want []string
	}{
		{form{'？', "전각", "기호", "?", "", ""}, []string{"물음표"}},
		{form{'＃', "전각", "기호", "#", "", ""}, []string{"제목", "샵", "태그"}},
		{form{'＝', "전각", "기호", "=", "", ""}, nil},
		{form{'，', "전각", "기호", ",", "", ""}, nil},
		{form{'㎡', "네모", "기호+로마자", "m2", "", ""}, nil},
		{form{'㈜', "괄호", "한글", "주", "", ""}, nil},
	} {
		if got := entryNames(ix.valueEntries(c.f)); !reflect.DeepEqual(got, c.want) {
			t.Errorf("%c: %q, want %q", c.f.letter, got, c.want)
		}
	}
}
