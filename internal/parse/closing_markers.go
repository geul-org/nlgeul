//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what ClosingMarkers — the closing table, row by row, for other packages
package parse

func ClosingMarkers() []Marker {
	var out []Marker
	for _, e := range endings {
		out = append(out, Marker{
			Form: e.form, Lead: e.lead, After: e.after, Kind: e.kind, Count: e.count,
			Eu: e.flags&fEu != 0, NoYo: e.flags&fNoYo != 0, Copula: e.flags&(fNeedIda|fSelfIda) != 0,
		})
	}
	return out
}
