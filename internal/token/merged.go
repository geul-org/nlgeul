//ff:func feature=token type=reader control=iteration dimension=1
//ff:what merged — the candidates as states of the net: the candidates that end in the same word (stateKey) are one state, the first that came standing for them and keeping, in preds, each candidate before them once in the order they came (the root, whatever object it is, as the scan's root); a state met before (kept at a place again) stays as it is and its preds are not changed; the states in the order their words first came, each once; a word printed the same as another of other values is counted (twins)
package token

import "slices"

func (sc *scan) merged(cands []*path) []*path {
	reg := sc.registry()
	var out []*path
	seen := map[*path]bool{}
	for _, p := range cands {
		k := stateKeyOf(p)
		rep, ok := reg.states[k]
		switch {
		case !ok:
			rep = sc.registered(p, k)
		case rep != p && !slices.Contains(rep.preds, sc.stateBefore(p)):
			rep.preds = append(rep.preds, sc.stateBefore(p))
		}
		if !seen[rep] {
			seen[rep] = true
			out = append(out, rep)
		}
	}
	return out
}
