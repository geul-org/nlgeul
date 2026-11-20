//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what markerForms — the set of the forms of markers (a particle table's forms, for pronounParticled)
package parse

func markerForms(markers []Marker) map[string]bool {
	out := map[string]bool{}
	for _, m := range markers {
		out[m.Form] = true
	}
	return out
}
