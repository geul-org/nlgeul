//ff:func feature=token type=helper control=sequence
//ff:what oddWeave — a portion made up over four letters with two ways of cutting: (0,1) (1,2) (2,3) (3,4) and (0,2) (2,4); the words at (0,2) and (2,3) stand side by side but no word at (2,3) comes after the one at (0,2); the weave of it (made with a word table's index), and its states by name
package token

func oddWeave() (*weave, map[string]*path) {
	sc := &scan{ix: scanOf("x").ix}
	root := sc.registry().root
	s := map[string]*path{}
	s["c"] = addState(sc, []*path{root}, wordStep(1, "c", 1, -1))
	s["x"] = addState(sc, []*path{root}, wordStep(2, "x", 1, -2))
	s["z"] = addState(sc, []*path{s["c"]}, wordStep(2, "z", 1, -1))
	s["y"] = addState(sc, []*path{s["z"]}, wordStep(3, "y", 1, -1))
	s["q"] = addState(sc, []*path{s["y"]}, wordStep(4, "q", 1, -1))
	s["t"] = addState(sc, []*path{s["x"]}, wordStep(4, "t", 1, -3))
	pt := portion{from: 0, to: 4, states: sc.registry().all}
	return newWeave(sc.ix, pt), s
}
