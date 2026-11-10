//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what collectTails — every closing of the table that is read at the end of s[:end]
package parse

// yo tells that s[end] is a closing 요: it is one more syllable and one
// more piece of the tail, and closings that take no 요 are skipped.
func collectTails(s []rune, end int, yo bool) []tailFact {
	var out []tailFact
	for _, e := range endings {
		if yo && e.flags&fNoYo != 0 {
			continue
		}
		t, ok := matchEnding(s, end, e, yo)
		if !ok {
			continue
		}
		if yo {
			t.size += 2
			t.marked += "+요"
		}
		out = append(out, t)
	}
	return out
}
