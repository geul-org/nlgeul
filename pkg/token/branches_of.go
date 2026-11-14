//ff:func feature=token type=helper control=iteration dimension=1
//ff:what branchesOf — the candidates of the states that stand at one span, one for each way a word is printed there, in the order the words were made: the state's word as a branch of its own with its verdict and preference, and with them the number of its birth; of two states that print the same word with other values the one that comes first by the order candidates have (stepBefore, the first made when equal) stands, the other is left out, as the old reading keeps the better of two candidates printed the same
package token

func branchesOf(ix *index, states []*path) []Branch {
	var out []Branch
	at := map[string]int{}
	var kept []step
	for _, p := range states {
		e := p.last
		br := Branch{Words: []Eojeol{e.word}, Verdict: e.verdict, Pref: e.pref, scores: []wordScore{{verdict: e.verdict, pref: e.pref, born: e.born}}}
		i, dup := at[e.sig]
		switch {
		case !dup:
			at[e.sig] = len(out)
			out, kept = append(out, br), append(kept, e)
		case ix.stepBefore(e, kept[i]):
			out[i], kept[i] = br, e
		}
	}
	return out
}
