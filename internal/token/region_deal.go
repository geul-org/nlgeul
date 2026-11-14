//ff:func feature=token type=helper control=sequence
//ff:what deal — the next way of the region to send: the one after the last sent, made when it has not been made yet (only one at a time), false when there is none
package token

func (rg *region) deal() (Segment, bool) {
	if rg.sent == len(rg.out) {
		seg, ok := rg.more()
		if !ok {
			return Segment{}, false
		}
		rg.out = append(rg.out, seg)
	}
	rg.sent++
	return rg.out[rg.sent-1], true
}
