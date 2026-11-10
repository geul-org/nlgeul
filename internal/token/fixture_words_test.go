//ff:func feature=token type=helper control=sequence
//ff:what wordsScan — a unit under reading with the small word table of testdata, used up to the given layer, its atoms found
package token

func wordsScan(text string, layers int) *scan {
	ix, err := newIndex(Options{Words: "testdata/words.tsv", Layers: layers})
	if err != nil {
		panic(err)
	}
	sc := &scan{ix: ix, s: []rune(text)}
	sc.prepare()
	return sc
}
