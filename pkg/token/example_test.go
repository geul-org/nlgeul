//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Example — the stream as another module uses it: the word table read once into a reader, two sentences read as streams, NextHead for the best way of each stretch with at most two candidates a word, and AgainHead for the ways after it only where the best way leans on a new word (고양이가: no way of known words only; 방돌에: the best way is of known words, 방 and 돌에, and nothing is asked after it)
package token_test

import "github.com/geul-org/nlgeul/pkg/token"

func Example() {
	r, err := token.New(token.Options{Words: "testdata/words-stream.tsv"})
	if err != nil {
		panic(err)
	}
	for _, s := range []string{"아버지가 방에 들어가신다", "고양이가 방돌에"} {
		readSentence(r.Stream(s), 2)
	}
	// Output:
	// [0-4] way 1 of 2
	//   아버지가: 아버지[열쇠:체언]+가[주격:가] (1.00 -3.43 -3.43) 아버지[열쇠:체언]+가[보격:가] (1.00 -3.43 -3.43)
	// [4-6] way 1 of 1
	//   방에: 방[열쇠:체언]+에[부사격:에] (1.00 -3.77 -3.77) 방에 (0.41 -8.92 -14.75)
	// [6-11] way 1 of 8
	//   들어가신다: 들어가[열쇠:동사]+시[선어말:시]+(ㄴ)다[맺음:ㄴ다] (1.00 -7.62 -7.62) 들어가신+다[맺음:다] (0.41 -18.10 -23.94)
	// [0-4] way 1 of 2
	//   고양이가: 고양이+가[주격:가] (0.41 -14.61 -20.45) 고양+이[접미:이]+가[주격:가] (0.41 -14.61 -20.45)
	// [0-4] way 2 of 2
	//   고양이: 고양+이[주격:이] (0.41 -10.40 -16.24) 고양+이[보격:이] (0.41 -10.40 -16.24)
	//   가: 가[열쇠:동사]+(아)[맺음:어(합침)] (0.65 -5.11 -7.89)
	//   none of known words only
	// [4-7] way 1 of 2
	//   방: 방[열쇠:체언] (0.65 -2.50 -5.28)
	//   돌에: 돌[열쇠:체언]+에[부사격:에] (0.65 -1.37 -4.14) 돌에 (0.41 -9.74 -15.57)
}
