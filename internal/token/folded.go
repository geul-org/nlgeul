//ff:func feature=token type=reader control=sequence
//ff:what folded — a key's entry with one more line of the word table folded in: the kind added, the greatest count kept, the kind's own count summed
package token

import "math/bits"

func folded(w word, kind, count int) word {
	w.by[bits.TrailingZeros(uint(kind))] += count
	return word{w.kinds | kind, max(w.count, count), w.by}
}
