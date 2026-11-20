//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what hold — after a region is read, whether the stream must read on before the rules of a sentence judge what is queued: each word of the region's best way in order closes the wait when it is a predicate and opens it when one of its candidates waits for the predicate after it (heldAfter). Nothing is asked with no rule
package parse

func (s *Stream) hold(rg *region) {
	if len(s.rules) == 0 {
		return
	}
	seg, _ := rg.peek()
	for _, n := range seg.Nodes {
		s.held = heldAfter(s.held, n, s.rules)
	}
}
