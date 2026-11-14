//ff:func feature=token type=reader control=sequence
//ff:what TestStreamFind — the region sent with the span given, or nil when none was
package token

import "testing"

func TestStreamFind(t *testing.T) {
	a, b := &region{span: Span{0, 3}}, &region{span: Span{3, 5}}
	st := &Stream{sent: []*region{a, b}}
	if st.find(Span{3, 5}) != b || st.find(Span{0, 3}) != a || st.find(Span{0, 5}) != nil {
		t.Fatal("find")
	}
}
