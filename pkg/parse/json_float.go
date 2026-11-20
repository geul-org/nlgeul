//ff:func feature=parse type=helper control=sequence
//ff:what jsonFloat — a number for JSON: null for infinity or not a number (JSON has neither)
package parse

import "math"

func jsonFloat(x float64) *float64 {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return nil
	}
	return &x
}
