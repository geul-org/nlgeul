//ff:func feature=token type=helper control=sequence
//ff:what lambdaOf — the place of a candidate in the sorting formula (decision 014-35): its preference plus Lambda times log10 of its verdict; minus infinity when the verdict is 0 or less (the same as internal/combo's order)
package token

import "math"

func lambdaOf(br Branch) float64 {
	if br.Verdict <= 0 {
		return math.Inf(-1)
	}
	return br.Pref + float64(Lambda*math.Log10(br.Verdict))
}
