//ff:func feature=token type=helper control=sequence
//ff:what wordStep — a word made up for the net tests: where it ends, how it is printed (its one part), and a verdict and preference
package token

func wordStep(end int, sig string, verdict, pref float64) step {
	return step{end: end, sig: sig, word: Eojeol{Parts: []Part{{Text: sig, Stem: true}}, End: end}, verdict: verdict, pref: pref}
}
