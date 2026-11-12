//ff:func feature=token type=helper control=sequence
//ff:what TestJoinBare — a bundle of letters with no name right after another such bundle is joined to it (the text of the bundle and of its part both); any other bundle is added; the first bundle is added
package token

import "testing"

func TestJoinBare(t *testing.T) {
	bare := func(text string) Bundle {
		return Bundle{Text: text, Branches: []Branch{signBranch(Part{Text: text, Stem: true}, 0)}}
	}
	named := Bundle{Text: "?", Branches: []Branch{signBranch(Part{Text: "?", Name: Name{Kind: "물음표"}}, 0)}}
	out := joinBare(nil, bare("-"))
	out = joinBare(out, bare(">"))
	if len(out) != 1 || out[0].Text != "->" || out[0].Branches[0].Words[0].Parts[0].Text != "->" {
		t.Fatal("joined", out)
	}
	out = joinBare(out, named)
	out = joinBare(out, bare("="))
	if len(out) != 3 || out[1].Text != "?" || out[2].Text != "=" {
		t.Fatal("a named sign stands between", out)
	}
	out = joinBare(out, bare("="))
	if len(out) != 3 || out[2].Text != "==" {
		t.Fatal(out)
	}
}
