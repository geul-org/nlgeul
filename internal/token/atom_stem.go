//ff:func feature=token type=reader control=iteration dimension=1
//ff:what atomStem — when the stem from b to end is exactly one atom: a reading for each name that joins it (or the run's own name) with the atom's verdict; or a whole run of digits and a unit that was cut in two (64GB): the run's own name and verdict
package token

func (sc *scan) atomStem(b, end int) []markRead {
	a, ok := sc.atoms[b]
	if (!ok || a.end != end) && sc.wide[b] == end {
		return []markRead{{runName, sc.ix.run}}
	}
	if !ok || a.end != end {
		return nil
	}
	if len(a.names) == 0 {
		return []markRead{{runName, sc.ix.run}}
	}
	v := 1.0
	if a.num {
		v = sc.ix.number
	}
	var out []markRead
	for _, n := range a.names {
		out = append(out, markRead{n, v})
	}
	return out
}
