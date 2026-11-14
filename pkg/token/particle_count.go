//ff:func feature=token type=reader control=sequence
//ff:what particleCount — the observed count of a particle's form as this layer keeps it (each form of a pair has its own); the count of the particle table when this layer keeps none
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func particleCount(m parse.Marker) int {
	if c, ok := particleCounts[m.Form]; ok {
		return c
	}
	return m.Count
}
