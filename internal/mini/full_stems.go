//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what fullStems — the stems under a word's longest tail readings
package mini

func fullStems(w *word) []string {
	var out []string
	for _, t := range w.tails {
		if t.size == w.longest {
			out = append(out, t.stem)
		}
	}
	return out
}
