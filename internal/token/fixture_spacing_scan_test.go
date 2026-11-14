//ff:func feature=token type=helper control=sequence
//ff:what spacingScan — a unit under reading with the spacing words table of testdata and the given spacing rules on, its atoms and cues found
package token

func spacingScan(text string, spacing ...string) *scan {
	ix, err := newIndex(Options{Words: "testdata/words-spacing.tsv", Spacing: spacing})
	if err != nil {
		panic(err)
	}
	sc := &scan{ix: ix, s: []rune(text)}
	sc.prepare()
	return sc
}
