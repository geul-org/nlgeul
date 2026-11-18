//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestPriorExcept — a claim on the determiner with the exceptions of priorExcept on it, evaluated by hand on rule.Around: 전 the determiner falls right after a -기 form, a number with its unit and a bare noun of an action (발병), and stands after such a noun with 의, after another noun, at the head, with a noun written onto it after a -기 form (전국민), and other letters (총) stand after a -기 form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPriorExcept(t *testing.T) {
	g, claim := sentenceClaim("시험", rule.Listed, listed("관형사", rule.ListSelfKind, []string{"열쇠:관형사"}))
	priorExcept(g, claim, []string{"발병"})
	r := sentenceRule{name: "시험", graph: g}
	det := rule.Near{Stem: "전", Kind: "열쇠:관형사"}
	other := rule.Near{Stem: "총", Kind: "열쇠:관형사"}
	noun := []rule.Near{{Stem: "사람", Kind: "열쇠:체언", Modifiable: true}}
	glued := []rule.Near{{Stem: "국민", Kind: "열쇠:체언", Modifiable: true, Glued: true}}
	before := func(stem, kind, last string) []rule.Near {
		return []rule.Near{{Stem: stem, Kind: kind, Last: last, Marks: last}}
	}
	for _, c := range []struct {
		name   string
		self   rule.Near
		before []rule.Near
		after  []rule.Near
		want   bool
	}{
		{"전 after -기", det, before("먹", "열쇠:동사", "이름:기"), noun, false},
		{"전 after a unit", det, before("3", "수", "단위:년"), noun, false},
		{"전 after a noun of an action", det, before("발병", "열쇠:체언", ""), noun, false},
		{"전 after a noun of an action with 의", det, before("발병", "열쇠:체언", "관형격:의"), noun, true},
		{"전 after another noun", det, before("사람", "열쇠:체언", ""), noun, true},
		{"전 at the head", det, nil, noun, true},
		{"전 with a noun written onto it after -기", det, before("먹", "열쇠:동사", "이름:기"), glued, true},
		{"전 with a noun written onto it after a noun of an action", det, before("발병", "열쇠:체언", ""), glued, true},
		{"총 after -기", other, before("먹", "열쇠:동사", "이름:기"), noun, true},
		{"총 after a noun of an action", other, before("발병", "열쇠:체언", ""), noun, true},
	} {
		if got := graphHolds(t, r, rule.Around{Self: c.self, Before: c.before, After: c.after}); got != c.want {
			t.Errorf("%s: %v, want %v", c.name, got, c.want)
		}
	}
}
