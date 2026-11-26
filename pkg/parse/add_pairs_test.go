//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestAddPairs — each sign of a pair gets a graph for each of its candidates: an opening bracket one, a closing bracket one, a quotation mark written the same way at both ends two (열기 first), a quotation mark that may close with none of its kind open (" ' ” ’) one more closing last, and so does a closing bracket that may (loneBrackets: ) } ] > 」 』 》 〉 ｣ 】 〕 ） ］ ｝); 35 signs give 55 graphs (the backtick only opens); the last closing of such a sign is the lone closing, which stands with nothing open where the usual one falls
package parse

import (
	"reflect"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAddPairs(t *testing.T) {
	ix := &index{}
	ix.addPairs()
	if len(ix.pairs) != len(pairOrder) || ix.graphs != 55 {
		t.Fatal(len(ix.pairs), ix.graphs)
	}
	for sign, want := range map[rune][]string{'(': {"열기:("}, '`': {"열기:`"}, ')': {"닫기:)", "닫기:)"}, '"': {"열기:\"", "닫기:\"", "닫기:\""}, '\'': {"열기:'", "닫기:'", "닫기:'"}, '“': {"열기:“"}, '”': {"닫기:”", "닫기:”"}, '’': {"닫기:’", "닫기:’"}, '》': {"닫기:》", "닫기:》"}, '｢': {"열기:｢"}, '｣': {"닫기:｣", "닫기:｣"}} {
		if got := entryNames(ix.pairs[sign]); !reflect.DeepEqual(got, want) {
			t.Errorf("%c: %q, want %q", sign, got, want)
		}
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyMark, rule.Mark{Left: rule.SideHangul, Right: rule.SideHangul, Run: 1})
	for _, sign := range loneShuts + loneBrackets {
		es := ix.pairs[sign]
		usual, _ := evaluate(es[len(es)-2].g, ctx, false)
		lone, _ := evaluate(es[len(es)-1].g, ctx, false)
		if es[len(es)-1].form != shutName || usual > 0 || lone <= 0 {
			t.Errorf("%c: the usual closing %.2f, the last %s %.2f", sign, usual, es[len(es)-1].form, lone)
		}
	}
}
