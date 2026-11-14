//ff:func feature=token type=reader control=iteration dimension=1
//ff:what findCues — where a word of the spacing words may end a word that no table lists, the places the rules of the spacing articles that are on show: for each place of the unit (not inside an atom) the words that link and close lists (45), the titles and offices (48), the dependent nouns after a nominal (42), the units after a Hangul numeral (43), and the groups of a Hangul numeral parted at 만, 억 and 조 (44); stands is made at every length so that a scan with no such rule on reads as it did
package token

// The cues are a permission, each by itself: no candidate is made
// here, and the search makes the ones its graphs let stand.
func (sc *scan) findCues() {
	sc.stands, sc.shows = make([]map[int]bool, len(sc.s)+1), map[int]map[int]bool{}
	for p := 0; p < len(sc.s); p++ {
		if sc.inside[p] {
			continue
		}
		sc.cueList(p, "45", linkWords)
		sc.cueList(p, "48", titleWords)
		sc.cueList(p, "42", afterNounWords)
		if sc.ix.spacing["43"] && sc.numeralBefore(p) {
			sc.cueList(p, "43", hangulUnits)
		}
		if sc.ix.spacing["44"] {
			sc.cueGroups(p)
		}
	}
}
