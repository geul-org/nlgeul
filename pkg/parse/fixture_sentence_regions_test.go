//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what sentenceRegions — every stretch of a sentence read with words-sentence.tsv and no rule of a sentence, as the regions a stream queues, in order (nothing judged, nothing lifted)
package parse

import "testing"

func sentenceRegions(t *testing.T, s string) []*region {
	t.Helper()
	st := sentenceStream(t, s, false)
	for st.cur != nil || st.read < len(st.units) {
		st.pull()
	}
	return st.queue
}
