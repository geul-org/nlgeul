//ff:func feature=parse type=rule control=sequence
//ff:what TestDeterminerWord — 이 and 어떤 standing alone are determiners; 이것은 is not
package parse

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestDeterminerWord(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("determiners", determiners)}
	yes, ev := determinerWord(ctxFor(readWord("이", false, &sentence{}), tailFact{}), specs)
	also, _ := determinerWord(ctxFor(readWord("어떤", false, &sentence{}), tailFact{}), specs)
	no, _ := determinerWord(ctxFor(readWord("이것은", false, &sentence{}), tailFact{}), specs)
	if yes != 1 || ev != "이" || also != 1 || no != -1 {
		t.Fatal(yes, ev, also, no)
	}
}
