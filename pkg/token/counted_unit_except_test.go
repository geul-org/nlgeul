//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestCountedUnitExcept — a claim with the exception "the word right after it is a unit a numeral determiner counts" on it, evaluated by hand on rule.Around: 하다 with ㄴ written as 한 falls before a unit (번, 명), and stands before a noun that is no unit and at the end; a candidate whose letters are no numeral determiner (쓴) stands before a unit; 한 written onto a noun stands before a unit, and written onto an opening sign it falls; countedUnits keeps the units that are no dependent noun a modifying form stands before (번 kept, 자·줄 left out)
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestCountedUnitExcept(t *testing.T) {
	g, claim := sentenceClaim("시험", rule.Listed, listed("꾸밈 ㄴ", rule.ListSelfLast, []string{haAdn}))
	countedUnitExcept(g, claim, numDeterminers, []string{"번", "명"})
	r := sentenceRule{name: "시험", graph: g}
	one := rule.Near{Text: "한", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true}
	glued := one
	glued.Glued = true
	wrote := rule.Near{Text: "쓴", Stem: "쓰", Kind: "열쇠:동사", Last: "꾸밈:ㄴ", Pred: true}
	obj := []rule.Near{{Stem: "일", Kind: "열쇠:체언", Last: "목적격:을", Marks: "목적격:을"}}
	next := func(stem string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: "열쇠:체언", Modifiable: true}}
	}
	for _, c := range []struct {
		name   string
		around rule.Around
		want   bool
	}{
		{"한 before 번", rule.Around{Self: one, Before: obj, After: next("번")}, false},
		{"한 before 명", rule.Around{Self: one, After: next("명")}, false},
		{"한 before a noun", rule.Around{Self: one, Before: obj, After: next("사람")}, true},
		{"한 at the end", rule.Around{Self: one, Before: obj}, true},
		{"쓴 before 번", rule.Around{Self: wrote, Before: obj, After: next("번")}, true},
		{"한 onto a noun before 번", rule.Around{Self: glued, Before: []rule.Near{{Stem: "창백", Kind: "열쇠:체언"}}, After: next("번")}, true},
		{"한 onto an opening sign before 번", rule.Around{Self: glued, Before: []rule.Near{{Stem: "“", Kind: "열기:1"}}, After: next("번")}, false},
	} {
		if got := graphHolds(t, r, c.around); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
	if !slices.Contains(countedUnits, "번") || slices.Contains(countedUnits, "자") || slices.Contains(countedUnits, "줄") {
		t.Errorf("countedUnits: 번 %v, 자 %v, 줄 %v", slices.Contains(countedUnits, "번"), slices.Contains(countedUnits, "자"), slices.Contains(countedUnits, "줄"))
	}
}
