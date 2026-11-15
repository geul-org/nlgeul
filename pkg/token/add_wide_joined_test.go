//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddWideJoined — the particles shortened with 는 are not read on a stem they make a word of the word table with: 나이론 gives the noun and no 나이+론, 표선, the noun and no 표+선; where the two make no word they stand as before (여기선: 여기+선, 학교론: 학교+론)
package token

import "testing"

func TestAddWideJoined(t *testing.T) {
	r := streamReader(t, "words-joined.tsv")
	for _, c := range []struct {
		text, has, lack string
	}{
		{"나이론", "나이론[열쇠:체언]", "나이[열쇠:체언]+론[부사격:론(로는)]"},
		{"표선,", "표선[열쇠:체언]", "표[열쇠:체언]+선[부사격:선(에서는)]"},
		{"여기선", "여기[열쇠:체언]+선[부사격:선(에서는)]", ""},
		{"학교론", "학교[열쇠:체언]+론[부사격:론(로는)]", ""},
	} {
		lines := branchLines(r, c.text)
		if !lines[c.has] || (c.lack != "" && lines[c.lack]) {
			t.Errorf("%s: %v", c.text, lines)
		}
	}
}
