//ff:func feature=rule type=helper control=sequence
//ff:what markFacts — the yes/no facts of a sign as flags (Mark…), for MarkIs: those its sides and its run show, with those its reader tells (Facts)
package rule

func markFacts(m Mark) int {
	have := m.Facts
	if m.Run >= 2 {
		have |= MarkRepeated
	}
	if m.Left != 0 {
		have |= MarkAfterWord
	}
	if m.Left&SideDigit != 0 {
		have |= MarkAfterDigit
	}
	if m.Numeric {
		have |= MarkNumeric
	}
	return have
}
