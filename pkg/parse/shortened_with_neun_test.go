//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestShortenedWithNeun — 이젠 is 이제 with 는, 으론 is 으로 with 는 after a noun that ends in a final consonant other than ㄹ, and 곤 is 고 with 는 after a verb stem; a name that ends in 곤 and a word with 곤 inside keep their readings as nouns
package parse

import "testing"

func TestShortenedWithNeun(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-ijen.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ sentence, word, want string }{
		{"이젠 늦었다.", "이젠", "이제[열쇠:체언]+(는)[보조사:는]"},
		{"앞으론 조심해라.", "앞으론", "앞[열쇠:체언]+으론[부사격:으론(으로는)]"},
		{"그것으론 안 된다.", "그것으론", "그것[열쇠:체언]+으론[부사격:으론(으로는)]"},
		{"그는 자주 놀러 가곤 했다.", "가곤", "가[열쇠:동사]+곤[맺음:곤(고는)]"},
		{"밥을 먹곤 한다.", "먹곤", "먹[열쇠:동사]+곤[맺음:곤(고는)]"},
		{"재곤에게 주었다.", "재곤에게", "재곤[열쇠:체언]+에게[부사격:에게]"},
		{"그는 피곤하다.", "피곤하다", "피곤[열쇠:체언] 하+다[맺음:다]"},
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
