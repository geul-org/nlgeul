//ff:func feature=token type=reader control=iteration dimension=1
//ff:what findMid — the signs between letters of a unit that are also read as a word of their own (apartSigns): the first of a run of one sign, not at the unit's head (nor right after a sign of a pair there), with letters after it, not inside an atom and not where an atom begins; each with the candidates its graphs let stand (none, no sign)
package token

import "strings"

// The place after a unit's head is left out: the letters after a
// closing sign are read on their own, with no sign at their head.
func (sc *scan) findMid() {
	sc.mid = map[int]midSign{}
	for q := sc.head + 1; q < len(sc.s); q++ {
		if !strings.ContainsRune(apartSigns, sc.s[q]) || sc.s[q-1] == sc.s[q] || sc.inside[q] {
			continue
		}
		if _, ok := sc.atoms[q]; ok {
			continue
		}
		end := q + 1
		for end < len(sc.s) && sc.s[end] == sc.s[q] {
			end++
		}
		if end >= len(sc.s) {
			continue
		}
		text := string(sc.s[q:end])
		marks := sc.judgeMark(sc.ix.marks[sc.s[q]], sc.midFacts(q, end), text)
		if len(marks) == 0 {
			continue
		}
		var reads []Part
		for _, m := range marks {
			reads = append(reads, Part{Text: text, Name: Name{Kind: m.name}})
		}
		sc.mid[q] = midSign{reads: reads}
	}
}
