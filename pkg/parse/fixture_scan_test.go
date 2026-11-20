//ff:func feature=parse type=helper control=sequence
//ff:what scanOf — a unit under reading with a fresh index (no word table), its atoms found, and a trace
package parse

func scanOf(text string) *scan {
	var notes []string
	ix, _ := newIndex(Options{})
	sc := &scan{ix: ix, s: []rune(text)}
	sc.prepare()
	sc.trace = &notes
	return sc
}
