//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddHa — with a word table the six written forms of clause 40 (치, 케, 타, 토록, 코자, 컨대) are kept by their last letter under the plain closings' names, read only with a word table (wide) and with no count, each with a graph of its own that is counted, and none of them under the index's keys; without a word table there are none; a plain closing the index does not have is passed over
package token

import "testing"

func TestAddHa(t *testing.T) {
	bare, _ := newIndex(Options{})
	if bare.haEnds != nil {
		t.Fatal("no word table, no written forms of clause 40", bare.haEnds)
	}
	ix, _ := newIndex(Options{Words: "testdata/words.tsv"})
	want := map[rune][2]string{'치': {"맺음:지", "치"}, '케': {"맺음:게", "케"}, '타': {"맺음:다", "타"}, '록': {"맺음:도록", "토록"}, '자': {"맺음:고자", "코자"}, '대': {"맺음:건대", "컨대"}}
	if len(ix.haEnds) != len(want) {
		t.Fatal(ix.haEnds)
	}
	for k, w := range want {
		es := ix.haEnds[k]
		if len(es) != 1 || es[0].name != w[0] || es[0].form != w[1] || !es[0].wide || es[0].count != 0 || es[0].class != classClosing || es[0].g == nil {
			t.Fatal(string(k), es)
		}
	}
	if ix.named('치', "맺음:지").g != nil || ix.named('지', "맺음:지").form != "지" || ix.named('대', "맺음:건대").form != "건대" || ix.named('지', "맺음:지").wide {
		t.Fatal("the written forms are copies, not under the index's keys, and the plain closings stay as they are")
	}
	empty := &index{byKey: map[rune][]*entry{}}
	empty.addHa()
	if len(empty.haEnds) != 0 || empty.graphs != 0 || empty.edges != 0 {
		t.Fatal("no plain closing, nothing", empty.haEnds, empty.graphs)
	}
	before := &index{byKey: map[rune][]*entry{'지': {ix.named('지', "맺음:지")}}, lexed: true}
	before.addHa()
	if len(before.haEnds) != 1 || len(before.haEnds['치']) != 1 || before.graphs != 1 || before.edges < 2 {
		t.Fatal("one plain closing, one written form with its graph", before.haEnds, before.graphs, before.edges)
	}
}
