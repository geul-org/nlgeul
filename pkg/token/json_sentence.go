//ff:type feature=token type=model
//ff:what JSONSentence — one sentence as nlgeul token --format json writes it (one line): its number, its text, the stretches in the order Stream.Next sends them, and the sum of lambdas of the best way of all of them
package token

type JSONSentence struct {
	Sentence   int           `json:"sentence"`
	Text       string        `json:"text"`
	Segments   []JSONSegment `json:"segments"`
	BestLambda *float64      `json:"best_lambda"`
}
