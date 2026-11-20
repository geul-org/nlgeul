//ff:func feature=parse type=helper control=sequence
//ff:what almost — two sums of preferences are the same but for the rounding of adding in another order (within 1e-9)
package parse

import "math"

func almost(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
