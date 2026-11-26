//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestNameSuffixStream — read as a stream with words-name.tsv: the suffix 이 after a name that ends in a consonant is a candidate beside the subject and the complement particle where the name is called (right before a comma or an exclamation mark at the unit's end), and with a particle after it as before; with no such sign after it, before a full stop or a question mark, after an adverb and after a dependent noun it is not read
package parse

import "testing"

func TestNameSuffixStream(t *testing.T) {
	r := streamReader(t, "words-name.tsv")
	suffixed := "영숙[열쇠:체언]+이[접미:이]\n"
	for _, c := range []struct {
		sentence   string
		has, lacks []string
	}{
		{"영숙이, 가자", []string{suffixed, "영숙[열쇠:체언]+이[주격:이]\n", "영숙[열쇠:체언]+이[보격:이]\n"}, nil},
		{"영숙이! 가자", []string{suffixed, "영숙[열쇠:체언]+이[주격:이]\n"}, nil},
		{"\"길동이!\"", []string{"길동[열쇠:체언]+이[접미:이]\n", "길동[열쇠:체언]+이[주격:이]\n"}, nil},
		{"영숙이가 온다", []string{"영숙[열쇠:체언]+이[접미:이]+가[주격:가]\n"}, nil},
		{"영숙이 온다", []string{"영숙[열쇠:체언]+이[주격:이]\n"}, []string{suffixed}},
		{"영숙이. 가자", []string{"영숙[열쇠:체언]+이[주격:이]\n"}, []string{suffixed}},
		{"영숙이? 가자", []string{"영숙[열쇠:체언]+이[주격:이]\n"}, []string{suffixed}},
		{"꼬박이, 온다", nil, []string{"이[접미:이]"}},
		{"온 듯이, 가자", nil, []string{"이[접미:이]"}},
	} {
		checkLines(t, c.sentence, streamText(r, c.sentence), c.has, c.lacks)
	}
}
