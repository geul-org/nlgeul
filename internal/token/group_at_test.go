//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestGroupAt — told that the piece closes its unit before a sign that ends a sentence, the auxiliary 야 and 이야 do not stand and the vocative 야 does; not told so, all stand; a particle that is not of the leading ones stands either way
package token

import "testing"

func TestGroupAt(t *testing.T) {
	sc := scanOf("학교야")
	open, _, _ := sc.groupAt(classParticle, 3, -1, false, false, false)
	shut, _, starts := sc.groupAt(classParticle, 3, -1, false, true, false)
	if len(open) != 2 || len(shut) != 1 {
		t.Fatal(open, shut)
	}
	for key, pc := range shut {
		if pc.names != vocName || starts[key] != 2 {
			t.Fatal(pc, starts)
		}
	}
	sc = scanOf("마당이야")
	open, _, _ = sc.groupAt(classParticle, 4, -1, false, false, false)
	shut, _, _ = sc.groupAt(classParticle, 4, -1, false, true, false)
	if _, ok := shut["·이야·7·보조사:이야"]; ok || open["·이야·7·보조사:이야"].names != "보조사:이야" {
		t.Fatal(open, shut)
	}
	sc = scanOf("마당도")
	open, _, _ = sc.groupAt(classParticle, 3, -1, false, false, false)
	shut, _, _ = sc.groupAt(classParticle, 3, -1, false, true, false)
	if len(open) != 1 || len(shut) != 1 {
		t.Fatal(open, shut)
	}
}
