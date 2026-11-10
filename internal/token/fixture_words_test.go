//ff:func feature=token type=helper control=sequence
//ff:what wordsScan — a unit under reading with the small word table of testdata, used up to the given layer
package token

func wordsScan(text string, layers int) *scan {
	ix, err := newIndex(Options{Words: "testdata/words.tsv", Layers: layers})
	if err != nil {
		panic(err)
	}
	return &scan{ix: ix, s: []rune(text)}
}
