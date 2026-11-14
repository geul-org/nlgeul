//ff:func feature=token type=helper control=sequence
//ff:what typoScan — a unit under reading with the small word table of testdata that has a misspelt form (words-typo.tsv), used up to the given layer, its atoms found
package token

func typoScan(text string, layers int) *scan {
	ix, err := newIndex(Options{Words: "testdata/words-typo.tsv", Layers: layers})
	if err != nil {
		panic(err)
	}
	sc := &scan{ix: ix, s: []rune(text)}
	sc.prepare()
	return sc
}
