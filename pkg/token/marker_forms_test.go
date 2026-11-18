//ff:func feature=token type=helper control=sequence
//ff:what TestMarkerForms — the forms of the markers given make the set (a form given twice once), and no markers make an empty set
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
)

func TestMarkerForms(t *testing.T) {
	got := markerForms([]parse.Marker{{Form: "이"}, {Form: "가"}, {Form: "이"}})
	if len(got) != 2 || !got["이"] || !got["가"] || got["을"] {
		t.Fatal(got)
	}
	if len(markerForms(nil)) != 0 {
		t.Fatal("no markers, no forms")
	}
	if !particleForms["는"] || !particleForms["를"] || !particleForms["에게"] || particleForms["다"] {
		t.Fatal("the particle tables' forms", len(particleForms))
	}
}
