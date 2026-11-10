//ff:func feature=token type=reader control=iteration dimension=1
//ff:what prepare — find the atoms of a unit before it is read: an address, a mail address, a path or an option that the unit begins with (one atom, not read inside); then every run of Latin letters and digits (the first with the sign at the unit's head, when one was kept), and the names whose dots, slashes and hyphens a joining reading stands for
package token

func (sc *scan) prepare() {
	sc.inside, sc.atoms, sc.numEnds, sc.adnEnds = make([]bool, len(sc.s)+1), map[int]atom{}, map[int]bool{}, map[int]bool{}
	sc.marks = make([]map[int]bool, len(sc.s)+1)
	runs := sc.headRun(sc.after(splitUnits(sc.s, runsOf(sc.s)), sc.wholeAtom()))
	for k := 0; k < len(runs); {
		chain := chainAt(sc.s, runs, k)
		sc.joinChain(chain)
		k += len(chain)
	}
}
