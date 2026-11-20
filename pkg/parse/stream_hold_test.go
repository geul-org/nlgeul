//ff:func feature=parse type=reader control=sequence
//ff:what TestStreamHold — with no rule nothing is asked and the stream does not wait; with a rule that waits, a region whose word has a waiting candidate (학생이) makes the stream wait, a region of a predicate (되었다) ends the wait, and a region of both in one way (학생이되었다) ends with the wait closed; a region with no way changes nothing
package parse

import "testing"

func TestStreamHold(t *testing.T) {
	rgs := sentenceRegions(t, "학생이 되었다")
	bare := sentenceStream(t, "학생이", false)
	bare.hold(rgs[0])
	if bare.held {
		t.Fatal("no rule, no wait")
	}
	st := sentenceStream(t, "학생이", false, waitingComplementRule())
	st.hold(rgs[0])
	if !st.held {
		t.Fatal("학생이 waits")
	}
	st.hold(newRegion(Contest{Text: "가나"}, 9, nil))
	if !st.held {
		t.Fatal("a region with no way keeps the wait")
	}
	st.hold(rgs[1])
	if st.held {
		t.Fatal("되었다 ends the wait")
	}
	st.hold(sentenceRegions(t, "학생이되었다")[0])
	if st.held {
		t.Fatal("the predicate after 학생이 in one way ends the wait")
	}
}
