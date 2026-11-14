//ff:func feature=token type=helper control=iteration dimension=1
//ff:what groupKeys — the keys of the candidates groupAt gave at j, in order: by their place or kind (groupBase), and the candidates of one place or kind in the order their entries are in the index
package token

import (
	"slices"
	"sort"
)

func (sc *scan) groupKeys(pieces map[string]piece, from map[string]*entry, j int) []string {
	if len(pieces) == 0 {
		return nil
	}
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	at := sc.entriesAt(j)
	sort.Slice(keys, func(a, b int) bool {
		ea, eb := from[keys[a]], from[keys[b]]
		if ba, bb := groupBase(ea), groupBase(eb); ba != bb {
			return ba < bb
		}
		return slices.Index(at, ea) < slices.Index(at, eb)
	})
	return keys
}
