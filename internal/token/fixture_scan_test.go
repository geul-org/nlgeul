//ff:func feature=token type=helper control=sequence
//ff:what scanOf — a unit under reading with a fresh index (no word table) and a trace
package token

func scanOf(text string) *scan {
	var notes []string
	ix, _ := newIndex(Options{})
	return &scan{ix: ix, s: []rune(text), trace: &notes}
}
