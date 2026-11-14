//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestSearchBorn — in the candidates of a unit read the old way every word has the number of its birth: the same word (where it begins and ends, how it is printed) the same number whatever the words before it, another word another number, every number below the next to give
package token

import "testing"

func TestSearchBorn(t *testing.T) {
	sc := wordsScan("곰이간다", layerCost)
	paths := sc.search()
	seen := map[wordKey]int{}
	used := map[int]wordKey{}
	for _, p := range paths {
		for q := p; q.prev != nil; q = q.prev {
			k := wordKeyOf(q)
			if n, ok := seen[k]; ok && n != q.last.born {
				t.Fatal("one word, two numbers", k, n, q.last.born)
			}
			if other, ok := used[q.last.born]; ok && other != k {
				t.Fatal("two words, one number", k, other)
			}
			seen[k], used[q.last.born] = q.last.born, k
		}
	}
	if len(seen) < 3 || len(seen) != sc.registry().next {
		t.Fatal(len(seen), sc.registry().next)
	}
}
