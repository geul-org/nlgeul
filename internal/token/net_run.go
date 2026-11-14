//ff:type feature=token type=model
//ff:what netRun — the net of a unit worked out one place at a time as the stream asks for stretches: the scan, the unit, the states kept at each place, the unit's length, the places worked out (j), the place of the last cut by the window (c) and the roots after it (seeds, those the cut kept among its states, and keptSeeds, those among the candidates kept there at a sign; none before the first cut), whether the first stretch may touch the head (a unit read by itself), how many stretches were made, the words planted at each place (the letters after a sign of a pair), where the numbers of birth begin, whether the window cuts, the stretches made and not yet sent, whether the unit's end was reached, and what is counted (stats); the cuts are tried before each place is worked out (next)
package token

type netRun struct {
	sc        *scan
	u         unit
	best      [][]*path
	n         int
	j         int
	c         int
	seeds     []*path
	keptSeeds []*path
	first     bool
	made      int
	plant     func(j int) []*path
	base      int
	window    bool
	rest      []unitStretch
	done      bool
	stats     *netStats
	lam       map[Span]spanLambda
}
