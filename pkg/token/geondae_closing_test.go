//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGeondaeClosing — the linking closing 건대 (살피건대, 바라건대, 듣건대) is read after a verb stem as one word, and a name 건대 that stands alone or takes a particle keeps its reading as a noun
package token

import "testing"

func TestGeondaeClosing(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-geondae.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ sentence, word, want string }{
		{"살피건대 원심의 판단은 옳다.", "살피건대", "살피[열쇠:동사]+건대[맺음:건대]"},
		{"바라건대 무사히 돌아오라.", "바라건대", "바라[열쇠:동사]+건대[맺음:건대]"},
		{"내가 듣건대 그가 왔다.", "듣건대", "듣[열쇠:동사]+건대[맺음:건대]"},
		{"그는 건대에 다닌다.", "건대에", "건대[열쇠:체언]+에[부사격:에]"},
		{"건대가 이겼다.", "건대가", "건대[열쇠:체언]+가[주격:가]"},
	} {
		var first []string
		for _, sg := range streamFirsts(r.Stream(c.sentence)) {
			first = append(first, firstsOf(sg, c.word)...)
		}
		if len(first) == 0 {
			t.Errorf("%s: no word %s", c.sentence, c.word)
		} else if first[0] != c.want {
			t.Errorf("%s: first is %q, want %s", c.sentence, first[0], c.want)
		}
	}
}
