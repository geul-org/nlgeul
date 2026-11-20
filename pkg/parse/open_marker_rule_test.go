//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestOpenMarkerRule — the rule is named as the trace prints it, is a rebuttal of the force forceUngrammatical that adds nothing and does not wait, and speaks of a candidate of markers alone whose first is a particle or a closing, not of a word with a stem, a punctuation mark or an empty one
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestOpenMarkerRule(t *testing.T) {
	r := openMarkerRule()
	if r.name != ruleOpenMarker || r.force != forceUngrammatical || r.bonus != 0 || r.graph == nil || r.waits != nil {
		t.Fatalf("%+v", r)
	}
	for _, c := range []struct {
		near rule.Near
		want bool
	}{
		{rule.Near{Last: "목적격:을", Marks: "목적격:을"}, true},
		{rule.Near{Last: "맺음:라고", Marks: "맺음:라고", Pred: true}, true},
		{rule.Near{Stem: "을", Kind: "열쇠:체언"}, false},
		{rule.Near{Last: "마침표", Marks: "마침표"}, false},
		{rule.Near{}, false},
	} {
		if got := r.about(c.near); got != c.want {
			t.Errorf("about(%+v) = %v, want %v", c.near, got, c.want)
		}
	}
}
