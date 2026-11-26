//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what byVerdict — the readings of a sign of a pair put in order: those of a higher verdict first, the others in the order they came (the order of the ways the signs before it may stand open); given back are the candidates, each kind and depth once (its first, so its highest verdict), and the ways the signs may stand open after the sign, each once; a reading as the letter alone has no candidate and counts as a verdict of 1
package parse

import "slices"

// The first way after the sign is the way of its first candidate, so the
// signs after it are read on from the reading its place speaks for.
func byVerdict(reads []pairWay) ([]pairing, []string) {
	sorted := slices.Clone(reads)
	slices.SortStableFunc(sorted, func(a, b pairWay) int {
		switch {
		case a.read.verdict > b.read.verdict:
			return -1
		case a.read.verdict < b.read.verdict:
			return 1
		}
		return 0
	})
	var candidates []pairing
	var ways []string
	seen := map[pairing]bool{}
	for _, rd := range sorted {
		ways = append(ways, rd.way)
		key := pairing{name: rd.read.name, depth: rd.read.depth}
		if rd.read.name == "" || seen[key] {
			continue
		}
		seen[key] = true
		candidates = append(candidates, rd.read)
	}
	return candidates, distinct(ways)
}
