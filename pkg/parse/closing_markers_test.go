//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestClosingMarkers — every closing row comes out; the flags are told apart (으면, 던, 예요)
package parse

import "testing"

func TestClosingMarkers(t *testing.T) {
	got := ClosingMarkers()
	if len(got) != len(endings) {
		t.Fatal(len(got))
	}
	seen := map[string]Marker{}
	for _, m := range got {
		if m.Lead == leadNone {
			seen[m.Form] = m
		}
	}
	if !seen["면"].Eu || !seen["던"].NoYo || seen["던"].Kind != kindAdn || !seen["예요"].SelfIda || !seen["에요"].NeedIda || !seen["는"].NoIda || seen["다"].NoIda {
		t.Fatal(seen["면"], seen["던"], seen["예요"])
	}
}
