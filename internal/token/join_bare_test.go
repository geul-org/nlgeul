//ff:func feature=token type=helper control=sequence
//ff:what TestJoinBare — a contest of letters with no name right after another such contest is joined to it (the text of the contest and of its part both); any other contest is added; the first contest is added
package token

import "testing"

func TestJoinBare(t *testing.T) {
	bare := func(text string) Contest {
		return Contest{Text: text, Branches: []Branch{signBranch(Part{Text: text, Stem: true}, 0)}}
	}
	named := Contest{Text: "?", Branches: []Branch{signBranch(Part{Text: "?", Name: Name{Kind: "물음표"}}, 0)}}
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
