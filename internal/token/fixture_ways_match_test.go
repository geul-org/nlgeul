//ff:func feature=token type=helper control=iteration dimension=1
//ff:what waysMatch — what a stream gives for each stretch of a sentence, all of it (Next, then Again to the end), against the contests Reader.Contests gives: ranks in order with lambdas that never rise, each way once, the ways of cutting exactly those the branches of the contest make, Ways their number, every way tiling its stretch; the names of what does not hold
package token

import "math/big"

func waysMatch(st *Stream, contests []Contest, letters []rune) []string {
	var bad []string
	at := 0
	for _, c := range contests {
		first, ok := st.Next()
		if !ok {
			return append(bad, "missing "+c.Text)
		}
		ways := streamWays(st, first)
		forms := formsOf(cutsOf(c, at))
		keys := segmentKeys(ways)
		if !ranked(ways) || len(keys) != len(ways) || len(ways) != len(forms) || first.Ways.Cmp(big.NewInt(int64(len(forms)))) != 0 {
			bad = append(bad, "ways "+c.Text)
		}
		bad = append(bad, missingForms(forms, keys)...)
		bad = append(bad, waysCheck(ways, letters)...)
		at += len([]rune(c.Text))
	}
	if _, ok := st.Next(); ok {
		bad = append(bad, "more than the contests")
	}
	return bad
}
