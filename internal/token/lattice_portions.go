//ff:func feature=token type=helper control=iteration dimension=2
//ff:what portions — the live states cut at the agreed places: one portion between two neighbouring places, with the states that lie in it (a state lies in the portion that holds where it begins; none lies across a place), in the order they were made; none when no way reads the unit
package token

import "sort"

func (l *lattice) portions() []portion {
	live := l.live()
	cuts := l.agreed(live)
	if len(cuts) < 2 {
		return nil
	}
	out := make([]portion, len(cuts)-1)
	for i := range out {
		out[i] = portion{from: cuts[i], to: cuts[i+1]}
	}
	for _, p := range live {
		i := sort.SearchInts(cuts, p.start+1) - 1
		out[i].states = append(out[i].states, p)
	}
	return out
}
