//ff:func feature=parse type=reader control=sequence
//ff:what TestParticleCount — a form this layer keeps a count for has that count (each form of a pair its own); another form keeps the count of the particle table
package parse

import (
	"testing"
)

func TestParticleCount(t *testing.T) {
	i, ga := particleCount(Marker{Form: "이", Count: 9}), particleCount(Marker{Form: "가", Count: 9})
	if i == 9 || ga == 9 || i == ga {
		t.Fatal(i, ga)
	}
	if got := particleCount(Marker{Form: "없는꼴", Count: 7}); got != 7 {
		t.Fatal(got)
	}
}
