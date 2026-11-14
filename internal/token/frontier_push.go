//ff:func feature=token type=helper control=sequence
//ff:what push — put a walk in the frontier behind every walk whose f is not less than its own (so that of equal f the first put in is first out, also when f is minus infinity)
package token

import "sort"

func (f *frontier) push(w *walk) {
	at := sort.Search(len(f.items), func(i int) bool { return w.f > f.items[i].f })
	f.items = append(f.items, nil)
	copy(f.items[at+1:], f.items[at:])
	f.items[at] = w
	f.peak = max(f.peak, len(f.items))
}
