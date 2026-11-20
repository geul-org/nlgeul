//ff:func feature=parse type=reader control=sequence
//ff:what TestAddAux — the closed set of the tables gives its stems but 지, which is always written onto the main verb (주, 보, 버리, 놓, 두, 있); with a word table 가, 내, 드리, 오 and 달 follow them; of a made list the one-syllable helpers (지, 하) are left out and a longer stem that begins with one of their letters is kept
package parse

import (
	"slices"
	"testing"
)

func TestAddAux(t *testing.T) {
	closed := ClosedWords()[auxSet]
	bare := &index{}
	bare.addAux(closed)
	want := []string{"주", "보", "버리", "놓", "두", "있"}
	if !slices.Contains(closed, "지") || !slices.Equal(bare.aux, want) {
		t.Fatal(closed, bare.aux)
	}
	lexed := &index{lexed: true}
	lexed.addAux(closed)
	if !slices.Equal(lexed.aux, slices.Concat(want, []string{"가", "내", "드리", "오", "달"})) {
		t.Fatal("with a word table", lexed.aux)
	}
	made := &index{}
	made.addAux([]string{"하", "지우", "지", "버리"})
	if !slices.Equal(made.aux, []string{"지우", "버리"}) {
		t.Fatal("a made list", made.aux)
	}
}
