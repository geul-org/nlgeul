//ff:func feature=token type=helper control=sequence
//ff:what TestReaches — a state that begins at the cut stays when a root is before it and not when none is; a state that begins after it stays when a state that stays is before it and not otherwise
package token

import "testing"

func TestReaches(t *testing.T) {
	a, b := &path{}, &path{}
	at := &path{start: 4, preds: []*path{a}}
	after := &path{start: 6, preds: []*path{b}}
	if !reaches(at, 4, map[*path]bool{a: true}, nil) || reaches(at, 4, map[*path]bool{b: true}, nil) {
		t.Fatal("a root before it")
	}
	if !reaches(after, 4, nil, map[*path]bool{b: true}) || reaches(after, 4, nil, map[*path]bool{a: true}) || reaches(after, 4, map[*path]bool{b: true}, nil) {
		t.Fatal("a state that stays before it")
	}
}
