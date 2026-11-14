//ff:func feature=token type=helper control=sequence
//ff:what peek — the best way of the region, made when it has not been made yet, without taking it as sent (deal gives it after); false when there is none
package token

func (rg *region) peek() (Segment, bool) {
	if len(rg.out) == 0 {
		seg, ok := rg.more()
		if !ok {
			return Segment{}, false
		}
		rg.out = append(rg.out, seg)
	}
	return rg.out[0], true
}
