//ff:func feature=token type=rule control=sequence
//ff:what TestAddShort — a shortened closing is indexed under its name with the form it is short for; one the copula does not take is listed for its graph, another is not
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/koparse/internal/parse"
	"github.com/geul-org/koparse/pkg/rule"
)

func TestAddShort(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}, pred: &entry{name: "g"}}
	ix.addShort(short{parse.Marker{Form: "재"}, "자고 해", true, rule.TailVerb})
	ix.addShort(short{parse.Marker{Form: "래", Eu: true}, "라고 해", false, rule.TailVerb})
	e := ix.named('재', "맺음:재(자고 해)")
	if e.g == nil || !e.wide || e.gate != ix.pred || e.tail != rule.TailVerb || !slices.Contains(ix.barred, e.name) {
		t.Fatal(e)
	}
	if len(ix.byKey['래']) != 2 || len(ix.barred) != 1 {
		t.Fatal(ix.byKey['래'], ix.barred)
	}
}
