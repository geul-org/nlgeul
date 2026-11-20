//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what waysLines — the segments each as its nodes in short (segmentLine)
package parse

func waysLines(segs []Segment) []string {
	var out []string
	for _, sg := range segs {
		out = append(out, segmentLine(sg))
	}
	return out
}
