//ff:func feature=token type=reader control=sequence
//ff:what TestEntriesOf — the entries that may end with a token are those under the token and then those under its key as a fused token (간: under ㄴ), each in the order of the index; a token no entry ends with has none
package token

import (
	"slices"
	"testing"
)

func TestEntriesOf(t *testing.T) {
	ix, _ := newIndex(Options{})
	if got, want := ix.entriesOf('간'), append(slices.Clone(ix.byKey['간']), ix.byKey['ㄴ']...); len(want) == 0 || !slices.Equal(got, want) {
		t.Fatal(len(got), len(want))
	}
	if got, want := ix.entriesOf('가'), append(slices.Clone(ix.byKey['가']), ix.byKey[keyVowel]...); len(ix.byKey['가']) == 0 || !slices.Equal(got, want) {
		t.Fatal(len(got), len(want))
	}
	if got := ix.entriesOf('a'); got != nil {
		t.Fatal(got)
	}
}
