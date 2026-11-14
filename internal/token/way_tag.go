//ff:func feature=token type=helper control=sequence
//ff:what wayTag — the tag on the first line of a way: its rank in a circle and the sum of the lambdas of its nodes ("① λ 합 -11.98")
package token

import "fmt"

func wayTag(seg Segment) string {
	return fmt.Sprintf("%s λ 합 %.2f", circled(seg.Rank), seg.Lambda)
}
