//ff:func feature=token type=reader control=iteration dimension=1
//ff:what prepare — find the atoms of a unit before it is read (and then the signs between letters that are also read as a word of their own): an address, a mail address, a path or an option that the unit begins with (one atom, not read inside); then every run of Latin letters and digits (the first with the sign at the unit's head, when one was kept), and the names whose dots, slashes and hyphens a joining candidate stands for, and then the cues of the spacing rules that are on (findCues), and last what the places of the unit are to a stem (how many tokens lie up to each, the letters as one string, the words of a bound stem) and from where a word that has its stem end at a place may begin (beginnings)
package token

func (sc *scan) prepare() {
	sc.inside, sc.atoms, sc.numEnds, sc.adnEnds, sc.detEnds = make([]bool, len(sc.s)+1), map[int]atom{}, map[int]bool{}, map[int]bool{}, map[int]bool{}
	sc.marks = make([]map[int]bool, len(sc.s)+1)
	whole := runsOf(sc.s)
	sc.wide = wideRuns(sc.s, whole)
	runs := sc.headRun(sc.after(splitUnits(sc.s, whole), sc.wholeAtom()))
	for k := 0; k < len(runs); {
		chain := chainAt(sc.s, runs, k)
		sc.joinChain(chain)
		k += len(chain)
	}
	sc.kept = map[int][]*path{}
	sc.registry()
	sc.findMid()
	sc.findCues()
	sc.countTokens()
	sc.placeLetters()
	sc.partStarts()
	sc.beginnings()
}
