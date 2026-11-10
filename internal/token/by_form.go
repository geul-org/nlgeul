//ff:func feature=token type=reader control=iteration dimension=1
//ff:what byForm — the particle readings that stand at one place, gathered by form: one piece for each form with its names joined, and the kinds of its readings (a suffix apart from particles of the same form)
package token

import (
	"sort"

	"github.com/geul-org/koparse/pkg/rule"
)

// The keys of pieces and from are those of readGroup. The forms come
// back in order.
func byForm(pieces map[string]piece, from map[string]*entry) ([]string, map[string]piece, map[string][]*entry) {
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var forms []string
	merged, kinds := map[string]piece{}, map[string][]*entry{}
	for _, k := range keys {
		f := from[k].form
		if from[k].slot == rule.LinkSuffix {
			f += "·접미"
		}
		if _, ok := merged[f]; !ok {
			forms = append(forms, f)
		}
		merged[f] = joinPieces(merged[f], pieces[k])
		kinds[f] = append(kinds[f], from[k])
	}
	return forms, merged, kinds
}
