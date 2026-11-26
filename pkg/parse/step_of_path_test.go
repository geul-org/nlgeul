//ff:func feature=parse type=reader control=sequence
//ff:what TestStepOfPath — a path read with the Hangul after its last slash as its name costs one token, as the path up to its slash does, and less than the new word of the same letters, which costs its tokens
package parse

import "testing"

func TestStepOfPath(t *testing.T) {
	sc := scanOf("~/a/b/회의록")
	sc.search()
	named := sc.stepOf(0, 9, tail{start: 9}, stemCandidate{kindCandidate{0, pathName, sc.ix.pathNamed}, "~/a/b/회의록", ""})
	atom := sc.stepOf(0, 6, tail{start: 6}, stemCandidate{kindCandidate{0, pathName, 1}, "~/a/b/", ""})
	plain := sc.stepOf(0, 9, tail{start: 9}, stemCandidate{kindCandidate{}, "~/a/b/회의록", ""})
	if named.pref != atom.pref || named.verdict != 1 || named.word.String() != "~/a/b/회의록[경로]" {
		t.Fatal(named.pref, atom.pref, named.verdict, named.word.String())
	}
	if plain.pref >= named.pref || plain.verdict >= named.verdict {
		t.Fatal(plain.pref, plain.verdict)
	}
}
