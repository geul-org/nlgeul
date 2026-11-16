//ff:func feature=parse type=reader control=sequence
//ff:what TestTailFacts — the facts of a tail reading the rules read are carried over (closing, kind, length, count)
package parse

import "testing"

func TestTailFacts(t *testing.T) {
	got := tailFacts(tailFact{"다", kindFinal, "방식이", 4, true, 7, "방식+이+다"})
	if got.Closing != "다" || got.Kind != kindFinal || got.Size != 4 || got.Count != 7 {
		t.Fatalf("%+v", got)
	}
}
