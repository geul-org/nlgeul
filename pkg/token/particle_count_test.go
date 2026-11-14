//ff:func feature=token type=reader control=sequence
//ff:what TestParticleCount — a form this layer keeps a count for has that count (each form of a pair its own); another form keeps the count of the particle table
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
)

func TestParticleCount(t *testing.T) {
	i, ga := particleCount(parse.Marker{Form: "이", Count: 9}), particleCount(parse.Marker{Form: "가", Count: 9})
	if i == 9 || ga == 9 || i == ga {
		t.Fatal(i, ga)
	}
	if got := particleCount(parse.Marker{Form: "없는꼴", Count: 7}); got != 7 {
		t.Fatal(got)
	}
}
