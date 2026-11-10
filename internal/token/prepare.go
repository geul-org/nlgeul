//ff:func feature=token type=reader control=iteration dimension=1
//ff:what prepare — find the atoms of a unit before it is read: every run of Latin letters and digits, and the dotted names whose dots a joining reading stands for
package token

func (sc *scan) prepare() {
	sc.inside, sc.atoms, sc.numEnds, sc.adnEnds = make([]bool, len(sc.s)+1), map[int]atom{}, map[int]bool{}, map[int]bool{}
	runs := splitUnits(sc.s, runsOf(sc.s))
	for k := 0; k < len(runs); {
		chain := chainAt(sc.s, runs, k)
		sc.joinChain(chain)
		k += len(chain)
	}
}
