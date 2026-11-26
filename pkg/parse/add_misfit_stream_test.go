//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAddMisfitStream — a particle written against the final consonant of a nominal of the word table is read as that particle, by its own name (마당+는[보조사:는], 학교+은[보조사:은], 사과+을[목적격:을], 마당+를[목적격:를], 학교+이[주격:이], 마당+와[부사격:와], 마당+로[부사격:로], 마당+야[보조사:야]), beside the readings there were (the whole letters as a new word); a particle written by the norm is read as before and has no second reading; a word the table does not hold takes no such particle (나무은), nor does a naming closing (먹음야) or a number's unit (3번가)
package parse

import (
	"strings"
	"testing"
)

func TestAddMisfitStream(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, has, not string
	}{
		{"마당는 넓다.", "\n마당[열쇠:체언]+는[보조사:는]\n", ""},
		{"학교은 멀다.", "\n학교[열쇠:체언]+은[보조사:은]\n", ""},
		{"사과을 먹었다.", "\n사과[열쇠:체언]+을[목적격:을]\n", ""},
		{"마당를 쓸었다.", "\n마당[열쇠:체언]+를[목적격:를]\n", ""},
		{"학교이 크다.", "\n학교[열쇠:체언]+이[주격:이]\n", ""},
		{"마당와 학교", "\n마당[열쇠:체언]+와[부사격:와]\n", ""},
		{"마당로 갔다.", "\n마당[열쇠:체언]+로[부사격:로]\n", ""},
		{"마당야 넓지.", "\n마당[열쇠:체언]+야[보조사:야]\n", ""},
		{"마당는 넓다.", "\n마당는\n", ""},
		{"마당은 넓다.", "\n마당[열쇠:체언]+은[보조사:은]\n", "+는["},
		{"학교는 멀다.", "\n학교[열쇠:체언]+는[보조사:는]\n", "+은["},
		{"나무은 크다.", "\n나무은\n", "나무+은["},
		{"먹음야 좋다.", "먹", "+야[보조사:야]"},
		{"3번가 길다.", "3", "+가[주격:가]"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || (c.not != "" && strings.Contains(got, c.not)) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
