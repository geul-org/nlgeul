//ff:func feature=parse type=helper control=sequence
//ff:what spanBefore — whether a span comes before another: by start, and for the same start by end
package parse

func spanBefore(a, b Span) bool {
	if a.Start != b.Start {
		return a.Start < b.Start
	}
	return a.End < b.End
}
