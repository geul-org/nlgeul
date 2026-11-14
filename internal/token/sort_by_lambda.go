//ff:func feature=token type=helper control=sequence
//ff:what sortByLambda — candidates by lambda, the greatest first; those of the same lambda by the number of their birth, the smaller first (a branch with none, noBorn, after those with one), those of the same number as they came (a candidate of minus infinity is last among the lambdas, in the same order)
package token

import "sort"

func sortByLambda(bs []Branch) {
	sort.SliceStable(bs, func(i, j int) bool {
		li, lj := lambdaOf(bs[i]), lambdaOf(bs[j])
		return li > lj || (li == lj && bornOf(bs[i]) < bornOf(bs[j]))
	})
}
