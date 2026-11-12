//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestAddTypoClosings — every row of the table of closings written wrong finds the right closing by its name among the closings of the tables and has an entry in the index, with the letters it is written with and the right closing as it is printed (in brackets for a final consonant); there are as many such entries as rows; a row whose name the tables do not have makes no entry
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddTypoClosings(t *testing.T) {
	ix, err := newIndex(Options{})
	if err != nil {
		t.Fatal(err)
	}
	rights, wrong := map[string]*entry{}, map[string]*entry{}
	for _, list := range ix.byKey {
		for _, e := range list {
			if e.wrote != "" {
				wrong[e.wrote+"|"+e.name] = e
			} else if e.class == classClosing && rights[e.name] == nil {
				rights[e.name] = e
			}
		}
	}
	if len(wrong) != len(typoClosings) {
		t.Fatal(len(wrong), len(typoClosings))
	}
	for _, row := range typoClosings {
		right, got := rights[row.Name], wrong[rule.LeadText(row.Lead)+row.Form+"|"+row.Name]
		if right == nil || got == nil {
			t.Errorf("%s: no entry for the right closing %s (%v, %v)", row.Form, row.Name, right, got)
			continue
		}
		want := right.form
		if right.lead != 0 {
			want = "(" + rule.LeadText(right.lead) + ")" + right.form
		}
		if got.right != want || got.g == nil || got.form != row.Form || got.lead != row.Lead || got.class != classClosing || got.count != right.count || got.head != right.head {
			t.Errorf("%s: %+v", row.Form, got)
		}
	}
	none := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	none.addTypoClosings(map[string]*entry{})
	if len(none.byKey) != 0 || none.graphs != 0 {
		t.Fatal("a row with no right closing makes no entry", none.graphs)
	}
}
