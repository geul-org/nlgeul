//ff:func feature=rule type=helper control=sequence
//ff:what markFacts — the yes/no facts of a sign as flags (Mark…), for MarkIs
package rule

func markFacts(m Mark) int {
	have := 0
	if m.Run >= 2 {
		have |= MarkRepeated
	}
	if m.Left != 0 {
		have |= MarkAfterWord
	}
	if m.Numeric {
		have |= MarkNumeric
	}
	return have
}
