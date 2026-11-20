//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what writtenApart — does the new word from b before tail t (which ends at j) hold, inside its stem, a place where the script changes right after a Latin letter, such that the stem from b to that place is one atom and the letters from that place to j read only as a word the spacing articles write apart (apartWord: A씨, B군, PD님); a place right after a digit is not asked, nor is a tail that names its own stem or stands for a number
package parse

import "unicode"

// The stem ends where the tail starts, or one syllable before when a
// piece is fused into its last syllable.
func (sc *scan) writtenApart(b, j int, t tail) bool {
	if t.only != "" || t.as != "" {
		return false
	}
	end := t.start
	if t.base != 0 {
		end = t.start - 1
	}
	for k := b + 1; k < end; k++ {
		if sc.change(k) && unicode.Is(unicode.Latin, sc.s[k-1]) && len(sc.atomStem(b, k)) > 0 && sc.apartWord(k, j) {
			return true
		}
	}
	return false
}
