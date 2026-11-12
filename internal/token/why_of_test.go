//ff:func feature=token type=rule control=sequence
//ff:what TestWhyOf — the name of the Claim among a node's criteria, wherever it stands and the first when there are two; "" when none is there
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWhyOf(t *testing.T) {
	if got := whyOf(nil); got != "" {
		t.Fatal("no criteria", got)
	}
	if got := whyOf(toulmin.Specs{fact(rule.PieceStem), not}); got != "" {
		t.Fatal("no Claim", got)
	}
	if got := whyOf(toulmin.Specs{fact(rule.PieceStem), not, &rule.Claim{Name: whyNoStem, Concludes: exceptLabel}}); got != whyNoStem {
		t.Fatal("the Claim after the criteria", got)
	}
	if got := whyOf(toulmin.Specs{&rule.Claim{Name: whyLast, Concludes: exceptLabel}, &rule.Claim{Name: whyNoStem, Concludes: exceptLabel}}); got != whyLast {
		t.Fatal("the first of two", got)
	}
}
