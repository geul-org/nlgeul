//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddForms — each sign of the tables of special letters gets its candidates in the order they are printed: a unit of the unit table written with the letter first (㎡), then each name the tables give it, then the named candidates of the letter it writes when the sign table does not name the sign itself (？: 물음표 after 전각:기호; ， and ＋, which the sign table names, get the shape alone); the letters of one name share one graph (ⓐ and ⓩ); a sign of the hand-written table gets each of its names in the order of its lines (○, ●), and a sign of no other table the kind it falls into (→, 😅); a letter of the tables that is no sign (①, ㊀, Ⅲ) gets none, nor does a sign a hand-written table reads get the kind it would fall into (a comma, a bracket, %, the hyphen); one graph for each name made here and one more for the verdict of a letter of a run read alone by its shape (loneForm)
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
	for sign, want := range map[rune][]string{'ⓐ': {"원:로마자"}, '㉠': {"원:한글"}, '㈎': {"괄호:한글"}, '⒜': {"괄호:로마자"}, '？': {"전각:기호", "물음표"}, '，': {"전각:기호"}, '＋': {"전각:기호"}, '㎡': {"단위:제곱미터", "네모:기호+로마자"}, '（': {"전각:기호"}, '○': {"가림:원", "기호:원"}, '●': {"글머리:검은 원", "가림:검은 원", "기호:도형"}, '→': {"화살표:오른쪽"}, '😅': {"기호:이모지"}, '①': nil, '㊀': nil, 'Ⅲ': nil, ',': nil, '(': nil, '%': nil, '-': nil} {
		if got := entryNames(ix.forms[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
	if ix.forms['ⓐ'][0] != ix.forms['ⓩ'][0] || ix.forms['ⓐ'][0] == ix.forms['㉠'][0] {
		t.Fatal("the letters of one name share one graph")
	}
	if ix.graphs-graphs != len(ix.formEntries)+1 || len(ix.formEntries) == 0 || ix.loneForm <= 0 {
		t.Fatal(ix.graphs-graphs, len(ix.formEntries))
	}
	saved := formList
	defer func() { formList = saved }()
	formList = []form{{'？', "전각", "기호", "?", "", ""}, {'？', "원", "기호", "?", "", ""}}
	jx := &index{}
	jx.addForms()
	if got := entryNames(jx.forms['？']); !reflect.DeepEqual(got, []string{"전각:기호", "원:기호"}) {
		t.Fatalf("a second line of the same sign adds its name after the first: %q", got)
	}
}
