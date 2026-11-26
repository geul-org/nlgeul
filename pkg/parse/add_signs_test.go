//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddSigns — each named sign gets a graph for each of its candidates in the order they are printed; the full stop's graphs are the closing candidates of the stops; 33 candidates in all (the ellipsis …, the middle dot variants ‧ ․ •, the slash, the dashes ― —, the full-width comma ， and plus ＋, the backslash, the box-drawing line ─ as 줄표 and the hyphen as 줄표 and 빼기 among them), one graph more, 줄임표 for a run of middle dots, which every middle dot takes after its 가운뎃점, and one for each sign with no name that is read apart between letters (the equals sign: a candidate with no name)
package parse

import (
	"reflect"
	"testing"
)

func TestAddSigns(t *testing.T) {
	ix := &index{}
	ix.addDots()
	graphs := ix.graphs
	ix.addSigns()
	if ix.graphs-graphs != 35 || len(ix.marks) != len(signOrderNamed)+1+len([]rune(bareApart)) {
		t.Fatal(ix.graphs-graphs, len(ix.marks))
	}
	if len(ix.marks['.']) != 2 || ix.marks['.'][0].name != "마침표" || ix.marks['.'][1].name != "줄임표" {
		t.Fatal("the full stop", ix.marks['.'])
	}
	for sign, want := range map[rune][]string{',': {"쉼표"}, '#': {"제목", "샵", "태그"}, '@': {"골뱅이", "명칭"}, '~': {"물결표"}, '∼': {"물결표"}, '～': {"물결표"}, '〜': {"물결표"}, '…': {"줄임표"}, '·': {"가운뎃점", "줄임표"}, '‧': {"가운뎃점", "줄임표"}, '․': {"가운뎃점", "줄임표"}, '•': {"가운뎃점", "줄임표"}, '/': {"빗금"}, '―': {"줄표"}, '—': {"줄표"}, '─': {"줄표"}, '-': {"줄표", "빼기"}, '+': {"더하기", "연결"}, '＋': {"더하기", "연결"}, '\\': {bareSignName}, '=': {""}} {
		if got := entryNames(ix.marks[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
}
