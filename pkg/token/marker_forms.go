//ff:func feature=token type=helper control=iteration dimension=1
//ff:what markerForms — the set of the forms of markers (a particle table's forms, for pronounParticled)
package token

import "github.com/geul-org/nlgeul/pkg/parse"

func markerForms(markers []parse.Marker) map[string]bool {
	out := map[string]bool{}
	for _, m := range markers {
		out[m.Form] = true
	}
	return out
}
