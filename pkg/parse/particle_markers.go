//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what ParticleMarkers — the particle table, row by row, for other packages
package parse

func ParticleMarkers() []Marker {
	var out []Marker
	for _, p := range particles {
		out = append(out, Marker{Form: p.form, After: p.after, Slot: p.slot, Kase: p.kase, Count: p.count})
	}
	return out
}
