//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddForms — each sign of the tables of special letters gets a candidate for each name the tables give it, and the letters of one name share one graph (ⓐ and ⓩ: 원:로마자); a letter of the tables that is no sign (①, ㊀, Ⅲ) gets none, nor does a sign the sign table already names (，, ＋)
package parse

import (
	"reflect"
	"testing"
)

func TestAddForms(t *testing.T) {
	ix := &index{}
	ix.addDots()
	ix.addSigns()
	graphs := ix.graphs
	ix.addForms()
	for sign, want := range map[rune][]string{'ⓐ': {"원:로마자"}, '㉠': {"원:한글"}, '㈎': {"괄호:한글"}, '⒜': {"괄호:로마자"}, '？': {"전각:기호"}, '①': nil, '㊀': nil, 'Ⅲ': nil, '，': nil, '＋': nil, '○': nil} {
		if got := entryNames(ix.forms[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
	if ix.forms['ⓐ'][0] != ix.forms['ⓩ'][0] || ix.forms['ⓐ'][0] == ix.forms['㉠'][0] {
		t.Fatal("the letters of one name share one graph")
	}
	names := map[string]bool{}
	for _, list := range ix.forms {
		names[list[0].name] = true
	}
	if ix.graphs-graphs != len(names) || len(names) == 0 {
		t.Fatal(ix.graphs-graphs, len(names))
	}
}
