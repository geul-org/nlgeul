//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what romanNumerals — the Roman numeral written in the Latin capitals I, V and X that ends at a place before a unit: a whole run of Latin letters and digits (an atom) that is such a numeral and is not left open by its letters (romanSure: XIV before 세, III before 호), named 수:로마십진; with the prefix 제 right before it, when that prefix's graph stands, the numeral begins at 제 and is named with it too, as a run of digits is (orderFront: 제II before 장)
package parse

import "math"

// The numeral is taken as one of the numerals its letters can write.
func (sc *scan) romanNumerals(end int) []numeral {
	var out []numeral
	for start := end - 1; start >= 0 && romanWritten(string(sc.s[start:end])); start-- {
		if a, ok := sc.atoms[start]; !ok || a.end != end || !romanSure(string(sc.s[start:end])) {
			continue
		}
		from, names := sc.orderFront(start, []string{romanNumName})
		for _, name := range names {
			out = append(out, numeral{string(sc.s[from:end]), math.Log10(1.0 / romanCount), name})
		}
	}
	return out
}
