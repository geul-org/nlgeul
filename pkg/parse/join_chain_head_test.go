//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinChainHead — the sign kept at the head of a unit's first run ($ .) is no letter of the run's side: a hidden name and a name joined by a slash are a path (.cache/db, $HOME/bin), a hidden name before a dot or a hyphen joins nothing, as before (.config.json, .well-known); a hidden name alone stays a run
package parse

import "testing"

func TestJoinChainHead(t *testing.T) {
	sc := scanOf(".cache/db")
	if a := sc.atoms[0]; a.end != 9 || atomNames(a) != "경로" || !sc.inside[1] || !sc.inside[7] {
		t.Fatal(sc.atoms, sc.inside)
	}
	sc = scanOf("$HOME/bin을")
	if a := sc.atoms[0]; a.end != 9 || atomNames(a) != "경로" || sc.inside[9] {
		t.Fatal(sc.atoms, sc.inside)
	}
	sc = scanOf(".config.json")
	if a := sc.atoms[0]; a.end != 7 || atomNames(a) != "" {
		t.Fatal(sc.atoms)
	}
	sc = scanOf(".well-known")
	if a := sc.atoms[0]; a.end != 5 || atomNames(a) != "" {
		t.Fatal(sc.atoms)
	}
	sc = scanOf(".env")
	if a := sc.atoms[0]; a.end != 4 || atomNames(a) != "" {
		t.Fatal(sc.atoms)
	}
}
