//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddPickedMore — the graphs of the forms added to pickedClosings: 의 takes the linking 어서 (and its fused form) before it and turns down 고 and 다, and 든지 and 든가 take the closings 다 and ㄴ다 and turn down 지
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddPickedMore(t *testing.T) {
	ix := &index{}
	ix.addPicked()
	for _, c := range []struct {
		form, before string
		takes        bool
	}{
		{"의", "맺음:어서", true}, {"의", "맺음:어서(합침)", true}, {"의", "맺음:고", false}, {"의", "맺음:다", false},
		{"든지", "맺음:다", true}, {"든지", "맺음:ㄴ다", true}, {"든가", "맺음:ㄴ다", true}, {"든가", "맺음:지", false},
	} {
		v, _ := verdictOf(ix.pickedAux[c.form].g, rule.Piece{Form: c.form, Before: c.before, HasStem: true, Next: -1}, "", false)
		if (v > 0) != c.takes {
			t.Errorf("%s after %s: %v", c.form, c.before, v)
		}
	}
}
