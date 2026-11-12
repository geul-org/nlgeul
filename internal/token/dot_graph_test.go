//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestDotGraph — each candidate of the full stop gets its exceptions: two for all but 줄임표, which has one
package token

import "testing"

func TestDotGraph(t *testing.T) {
	for _, name := range append(append([]string{}, joinNames...), stopNames...) {
		g, edges := dotGraph(name)
		if g == nil || (name == "줄임표") != (edges == 1) {
			t.Errorf("%s: %d", name, edges)
		}
	}
}
