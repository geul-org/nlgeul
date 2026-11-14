//ff:func feature=token type=helper control=iteration dimension=1
//ff:what segmentKeys — the ways of cutting the segments are, each as its spans written by lineKey (once for each way)
package token

func segmentKeys(segs []Segment) map[string]int {
	out := map[string]int{}
	for _, sg := range segs {
		var spans []Span
		for _, n := range sg.Nodes {
			spans = append(spans, n.Span)
		}
		out[lineKey(spans)]++
	}
	return out
}
