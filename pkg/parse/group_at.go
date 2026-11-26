//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what groupAt — judge every entry of one class that ends at j and file the candidates that stand by their key (entryGroup: the same token sequence of the same place or kind, and the same name); two entries of one key keep the one with the higher verdict; the graphs are told whether the piece closes its unit, whether it stands right on its stem, and whether it ends the unit's letters where one is called (before a comma or an exclamation mark: Called)
package parse

// The pieces come back by their keys, with the entry behind each and
// where its form starts. next is the place of the piece that follows. closed
// tells the graphs that the piece ends its unit right on its stem and a
// sign that ends a sentence stands after it. onStem tells them that the
// piece stands right on its stem (a piece before the particle 요 then
// knows whether the stem ends in a vowel, or the suffix 이 stands left of it: YoOpen).
func (sc *scan) groupAt(class, j, next int, more, closed, onStem bool) (map[string]piece, map[string]*entry, map[string]int) {
	var pieces map[string]piece
	var from map[string]*entry
	var starts map[string]int
	for _, e := range sc.entriesAt(j) {
		start, p, ok := sc.place(e, j, more)
		if !ok || e.class != class {
			continue
		}
		p.Next, p.Closed = next, closed
		p.Called = sc.called && j == len(sc.s) && next < 0
		p.YoOpen = sc.yoOpen(p, start, onStem)
		if e.wide && !more {
			p.Right = sc.right(j)
		}
		v := sc.judge(e, start, j, p)
		key := entryGroup(e)
		if v <= 0 || (from[key] != nil && pieces[key].verdict >= v) {
			continue
		}
		if pieces == nil {
			pieces, from, starts = map[string]piece{}, map[string]*entry{}, map[string]int{}
		}
		pieces[key] = sc.entryPiece(e, string(sc.s[start:j]), v)
		from[key], starts[key] = e, start
	}
	return pieces, from, starts
}
