//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestAddSpoken — with a word table every row of the table of markers written as they are spoken has an entry under the key of its written form, with the standard marker's name and class, the written form, no count, a graph of its own, and the gate of its class (a predicate's stem for a closing, a nominal for a particle), and the standard marker stays as it is; the total of the counts is not changed; without a word table there are none; a row whose standard marker the index does not have, or has with another lead, makes none
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddSpoken(t *testing.T) {
	bare, _ := newIndex(Options{})
	if e := bare.named('구', "맺음:고"); e.g != nil {
		t.Fatal("no word table, no written forms", e)
	}
	ix, err := newIndex(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range spokenMarkers {
		f := []rune(row.Form)
		e := &entry{}
		for _, c := range ix.byKey[f[len(f)-1]] {
			if c.name == row.Name && c.form == row.Form {
				e = c
			}
		}
		if e.g == nil || e.form != row.Form || e.lead != row.Lead || e.count != 0 || e.needs != (row.Copula || row.Name == "맺음:올시다") {
			t.Errorf("%s %s: %+v", row.Form, row.Name, e)
			continue
		}
		if want := map[int]*entry{classClosing: ix.pred, classParticle: ix.nominal}[e.class]; want == nil || e.gate != want {
			t.Errorf("%s %s: gate %v", row.Form, row.Name, e.gate)
		}
	}
	if e := ix.named('고', "맺음:고"); e.form != "고" || e.gate != nil {
		t.Fatal("the standard marker is changed", e)
	}
	old := spokenMarkers
	t.Cleanup(func() { spokenMarkers = old })
	spokenMarkers = []spoken{{"구", 0, "맺음:고", false}, {"두", 0, "보조사:도", false}, {"구", 0, "맺음:없는맺음", false}, {"다구", rule.JongL, "맺음:다고", false}}
	one := &index{byKey: map[rune][]*entry{'고': {ix.named('고', "맺음:고"), ix.named('고', "맺음:다고")}, '도': {ix.named('도', "보조사:도")}}, words: map[string]*entry{}, lexed: true, pred: ix.pred, nominal: ix.nominal, total: 5}
	one.addSpoken()
	gu, du := one.named('구', "맺음:고"), one.named('두', "보조사:도")
	if len(one.byKey['구']) != 1 || len(one.byKey['두']) != 1 || one.graphs != 2 || one.edges < 2 || one.total != 5 {
		t.Fatal("two rows with a standard marker, two entries", one.byKey, one.graphs, one.edges, one.total)
	}
	if gu.gate != ix.pred || du.gate != ix.nominal || gu.g == nil || gu.g == ix.named('고', "맺음:고").g {
		t.Fatal("a closing after a predicate's stem, a particle after a nominal, each with its own graph", gu, du)
	}
}
