//ff:func feature=token type=helper control=sequence
//ff:what TestParticleAfterClosing — a particle after a closing among the parts after the stem; not a particle before it, nor a closing alone
package token

import "testing"

func TestParticleAfterClosing(t *testing.T) {
	if !particleAfterClosing(spacingWord("읽/열쇠:동사", "어/맺음:어", "도/보조사:도")) {
		t.Fatal("읽어도")
	}
	if particleAfterClosing(spacingWord("씨/열쇠:체언", "가/주격:가", "다/맺음:다")) || particleAfterClosing(spacingWord("읽/열쇠:동사", "어/맺음:어")) {
		t.Fatal("none")
	}
}
