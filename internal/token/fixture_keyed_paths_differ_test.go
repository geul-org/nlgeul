//ff:func feature=token type=helper control=iteration dimension=2
//ff:what keyedPathsDiffer — two words of some candidates that have the same state key but are not the same word ("" when there are none)
package token

import (
	"fmt"
	"reflect"
)

func keyedPathsDiffer(paths []*path) string {
	seen := map[stateKey]Eojeol{}
	for _, p := range paths {
		for q := p; q.prev != nil; q = q.prev {
			k := stateKeyOf(q)
			if w, ok := seen[k]; ok && !reflect.DeepEqual(w, q.last.word) {
				return fmt.Sprint(k, w, q.last.word)
			}
			seen[k] = q.last.word
		}
	}
	return ""
}
