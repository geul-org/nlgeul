//ff:func feature=parse type=reader control=sequence
//ff:what TestTailFacts — every fact of a tail reading is carried over
package parse

import "testing"

func TestTailFacts(t *testing.T) {
	got := tailFacts(tailFact{"다", kindFinal, "방식이", 4, true, 7, "방식+이+다"})
	if got.Closing != "다" || got.Kind != kindFinal || got.Stem != "방식이" || got.Size != 4 || !got.Copula || got.Count != 7 || got.Marked != "방식+이+다" {
		t.Fatalf("%+v", got)
	}
}
