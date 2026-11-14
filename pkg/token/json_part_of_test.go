//ff:func feature=token type=helper control=sequence
//ff:what TestJSONPartOf — a part of a table has its name and is known, a stem no table knows has no name and is not known, a sign with no name is known, a part written wrong has the right name with the form written, a part with no trace is restored
package token

import "testing"

func TestJSONPartOf(t *testing.T) {
	p := jsonPartOf(Part{Text: "만", Name: Name{Kind: "보조사", Form: "만"}})
	if p.Name == nil || *p.Name != "보조사:만" || !p.Known || p.Restored || p.Stem {
		t.Fatal(p)
	}
	p = jsonPartOf(Part{Text: "쪼간", Stem: true})
	if p.Name != nil || p.Known || !p.Stem {
		t.Fatal(p)
	}
	p = jsonPartOf(Part{Text: "->", Stem: true})
	if p.Name != nil || !p.Known {
		t.Fatal(p)
	}
	p = jsonPartOf(Part{Text: "게", Hidden: "ㄹ", Name: Name{Kind: "맺음", Form: "게"}, Written: "ㄹ께"})
	if p.Name == nil || *p.Name != "맺음:게,오타정정:ㄹ께" || p.Hidden != "ㄹ" {
		t.Fatal(p.Name)
	}
	if p = jsonPartOf(Part{Text: "이", Name: Name{Kind: "서술격", Form: "이"}, Unseen: true}); !p.Restored {
		t.Fatal(p)
	}
}
