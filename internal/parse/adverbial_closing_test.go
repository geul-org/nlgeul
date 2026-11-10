//ff:func feature=parse type=rule control=sequence
//ff:what TestAdverbialClosing — a closing in the criterion list makes an adverbial; another does not
package parse

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAdverbialClosing(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("adverbialClosings", adverbialClosings)}
	yes, _ := adverbialClosing(ctxFor(&word{}, tailFact{closing: "게", size: 2}), specs)
	no, _ := adverbialClosing(ctxFor(&word{}, tailFact{closing: "고", size: 2}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
