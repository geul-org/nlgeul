//ff:func feature=token type=helper control=sequence
//ff:what scanOf — a unit under reading with a fresh index and a trace
package token

func scanOf(text string) *scan {
	var notes []string
	return &scan{ix: newIndex(), s: []rune(text), trace: &notes}
}
