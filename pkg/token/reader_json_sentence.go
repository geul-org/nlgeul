//ff:func feature=token type=reader control=iteration dimension=1
//ff:what JSONSentence — a sentence read as a stream and set down as data: each stretch Next sends (its best way with every candidate of every word, nothing cut), in the order it sends them, and the sum of their lambdas; index is the number of the sentence
package token

func (r *Reader) JSONSentence(sentence string, index int) JSONSentence {
	st := r.Stream(sentence)
	letters := lettersOf(sentence)
	out := JSONSentence{Sentence: index, Text: sentence, Segments: []JSONSegment{}}
	total := 0.0
	for seg, ok := st.Next(); ok; seg, ok = st.Next() {
		out.Segments = append(out.Segments, jsonSegmentOf(seg, letters))
		total += seg.Lambda
	}
	out.BestLambda = jsonFloat(total)
	return out
}
