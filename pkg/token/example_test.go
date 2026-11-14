//ff:func feature=token type=reader control=iteration dimension=1
//ff:what Example — the stream as another module uses it: the word table read once into a reader, two sentences read as streams, NextHead for the best way of each stretch with at most two candidates a word, and AgainHead for the ways after it only where the best way leans on a new word (고양이가: no way of known words only; 방돌에: the second way is)
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
	//   아버지가: 아버지[열쇠:체언]+가[주격:가] (0.87 -3.61 -4.50) 아버지[열쇠:체언]+가[보격:가] (0.87 -3.61 -4.50)
	// [4-6] way 1 of 1
	//   방에: 방[열쇠:체언]+에[부사격:에] (0.80 -3.95 -5.41) 방에 (0.33 -9.10 -16.25)
	// [6-11] way 1 of 8
	//   들어가신다: 들어가[열쇠:동사]+시[선어말:시]+(ㄴ)다[맺음:ㄴ다] (0.83 -7.80 -9.03) 들어가신+다[맺음:다] (0.33 -18.28 -25.44)
	// [0-4] way 1 of 2
	//   고양이가: 고양이+가[주격:가] (0.33 -14.79 -21.95) 고양+이[접미:이]+가[주격:가] (0.33 -14.79 -21.95)
	// [0-4] way 2 of 2
	//   고양이: 고양+이[주격:이] (0.33 -10.58 -17.74) 고양+이[보격:이] (0.33 -10.58 -17.74)
	//   가: 가[열쇠:동사]+(아)[맺음:어(합침)] (0.56 -5.32 -9.05)
	//   none of known words only
	// [4-7] way 1 of 2
	//   방돌에: 방돌+에[부사격:에] (0.33 -10.37 -17.52) 방돌에 (0.33 -13.14 -20.30)
	// [4-7] way 2 of 2
	//   방: 방[열쇠:체언] (0.47 -2.68 -7.61)
	//   돌에: 돌[열쇠:체언]+에[부사격:에] (0.46 -5.06 -10.08) 돌에 (0.33 -10.10 -17.25)
}
