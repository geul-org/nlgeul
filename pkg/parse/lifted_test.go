//ff:func feature=parse type=helper control=sequence
//ff:what TestLifted — a node gains the lifts of its own span: the candidate a lift names is raised and the candidates are put in order by lambda again (the tied 보격 이 goes before 주격 이), none is taken away, and the node it was made from keeps its candidates and their order; the lifts of another span leave the node as it is
package parse

import "testing"

func TestLifted(t *testing.T) {
	n := firstNode(t, "학생이")
	const subj, comp = "학생[열쇠:체언]+이[주격:이]", "학생[열쇠:체언]+이[보격:이]"
	if _, i := branchNamed(n, subj); i != 0 || n.Contest.Branches[1].String() != comp || n.Contest.Branches[0].Lambda() != n.Contest.Branches[1].Lambda() {
		t.Fatal("the tied 주격 is first before the lift", n.Contest.Branches)
	}
	was, _ := branchNamed(n, comp)
	other := lift{span: Span{3, 6}, word: comp, add: 9}
	if same := lifted(n, []lift{other}); &same.Contest.Branches[0] != &n.Contest.Branches[0] {
		t.Fatal("another span's lift touches the node")
	}
	up := lifted(n, []lift{other, {span: n.Span, word: comp, add: 2}})
	if len(up.Contest.Branches) != len(n.Contest.Branches) || up.Contest.Branches[0].String() != comp || up.Contest.Branches[0].Pref != was.Pref+2 || up.Contest.Branches[1].String() != subj {
		t.Fatal(up.Contest.Branches)
	}
	if n.Contest.Branches[0].String() != subj || n.Contest.Branches[1].Pref != was.Pref {
		t.Fatal("the node lifted from is touched")
	}
	if up.Span != n.Span || up.Contest.Text != n.Contest.Text {
		t.Fatal("the node's place")
	}
}
