//ff:func feature=token type=rule control=sequence
//ff:what TestAddLoneShut — the graph of a quotation mark that closes with none of its kind open is put under the sign after its other graphs, as a closing (닫기) with the sign after a colon in its name, and the graphs and edges are counted (three exceptions and two against two of them, four and two when the row bars the sign inside a word); the graph put is the lone closing's: it stands where the usual closing falls
package token

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAddLoneShut(t *testing.T) {
	ix := &index{pairs: map[rune][]*entry{}}
	ix.addPair('"', pairRows['"'], shutName)
	ix.addLoneShut('"', pairRows['"'])
	ix.addLoneShut('’', pairRows['’'])
	if got := entryNames(ix.pairs['"']); !reflect.DeepEqual(got, []string{"닫기:\"", "닫기:\""}) {
		t.Fatal(got)
	}
	lone := ix.pairs['’']
	if len(lone) != 1 || lone[0].name != "닫기:’" || lone[0].form != shutName || lone[0].g == nil {
		t.Fatal(lone)
	}
	if ix.graphs != 3 || ix.edges != 1+5+6 {
		t.Fatal(ix.graphs, ix.edges)
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyMark, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Run: 1})
	usual, _ := evaluate(ix.pairs['"'][0].g, ctx, false)
	added, _ := evaluate(ix.pairs['"'][1].g, ctx, false)
	if ix.pairs['"'][1].form != shutName || usual > 0 || added <= 0 {
		t.Fatal("the usual closing and the lone one with nothing open", usual, added)
	}
}
