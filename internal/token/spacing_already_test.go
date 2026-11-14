//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingAlready — the rules of articles 41 and 47 and 43 find boundaries the search already makes: 마당에서놀았다 and 읽어보았다 have the same candidates with every rule on as off (마당+에서 놀+았+다, 읽+어 보+았+다 are in both), and 고양이세마리 has 세 마리 in both and one candidate more with article 43 on (with every rule on 양 and 이 are titles and more are made)
package token

import "testing"

func TestSpacingAlready(t *testing.T) {
	off, all := spacingReader(t, "words-spacing.tsv"), spacingReader(t, "words-spacing.tsv", SpacingArticles...)
	for in, line := range map[string]string{
		"마당에서놀았다": "마당[열쇠:체언]+에서[부사격:에서] 놀[열쇠:동사]+았[선어말:았]+다[맺음:다]",
		"읽어보았다":   "읽[열쇠:동사]+어[맺음:어] 보[열쇠:동사]+았[선어말:았]+다[맺음:다]",
	} {
		if !branchLines(off, in)[line] || len(branchLines(off, in)) != len(branchLines(all, in)) {
			t.Fatal(in, len(branchLines(off, in)), len(branchLines(all, in)))
		}
	}
	const line = "고양이[열쇠:체언] 세[열쇠:관형사] 마리[열쇠:체언]"
	unit := spacingReader(t, "words-spacing.tsv", "43")
	if !branchLines(off, "고양이세마리")[line] || !branchLines(unit, "고양이세마리")[line] || len(branchLines(unit, "고양이세마리")) != len(branchLines(off, "고양이세마리"))+1 {
		t.Fatal(len(branchLines(off, "고양이세마리")), len(branchLines(unit, "고양이세마리")))
	}
}
