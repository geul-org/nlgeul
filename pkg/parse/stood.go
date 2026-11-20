//ff:func feature=parse type=reader control=sequence
//ff:what stood — may the letters from b up to j stand as a word of no markers inside the unit although no table lists them (a cue of the spacing rules made them one)
package parse

func (sc *scan) stood(b, j int) bool {
	return j < len(sc.stands) && sc.stands[j][b]
}
