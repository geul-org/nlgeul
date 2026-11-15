//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestMarkerGraphGlossSentence — end to end, by the lines of the stream (streamText) and by its JSON (JSONSentence): a Hangul word glossed right after with as many Han characters in round brackets reads no marker (대곤(大棍) loses 대+곤 of 고는, 관치(官治) loses 관하+치 of clause 40, 오기(傲氣) loses 오+기, "관치(官治)" with a quotation mark before it and '관치(官治)' with a quotation mark at its head that may close as well) and keeps its known and new readings; the same words with no gloss, with a space before the bracket, with a count of characters that differs and with a comma in the bracket keep the marker; the known harm: a gloss of meaning (운용한다(行陣和睦)) loses 운용하+(ㄴ)다; a suffix is not turned down so: 보장률(保障率) keeps 보장+률 as 보장률이 does
package token

import (
	"strings"
	"testing"
)

func TestMarkerGraphGlossSentence(t *testing.T) {
	r := streamReader(t, "words-gloss.tsv")
	for _, c := range []struct {
		sentence string
		gone     []string
		kept     []string
	}{
		{"그는 대곤(大棍)을 맞았다.", []string{"곤[맺음:곤(고는)]", "(ㄴ)[꾸밈:ㄴ]"}, []string{"\n대곤\n", "\n대[열쇠:체언] 곤[열쇠:체언]\n", "\n+을[목적격:을]\n"}},
		{"관치(官治)의 폐해가 크다.", []string{"치[맺음:지]"}, []string{"\n관치\n", "\n관[열쇠:체언] 치[열쇠:체언]\n", "\n+의[관형격:의]\n"}},
		{"\"관치(官治)\"의 폐해", []string{"치[맺음:지]"}, []string{"\n관치\n", "\n관[열쇠:체언] 치[열쇠:체언]\n"}},
		{"그가 '대곤' '관치(官治)'라 했다.", []string{"치[맺음:지]"}, []string{"\n관치\n", "\n관[열쇠:체언] 치[열쇠:체언]\n", "\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"오기(傲氣)를 부렸다.", []string{"기[이름:기]"}, []string{"\n오기[열쇠:체언]\n"}},
		{"그는 대곤 했지만 소용없었다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"관치 않다.", nil, []string{"\n관하[열쇠:동사]+치[맺음:지]\n"}},
		{"오기를 부렸다.", nil, []string{"\n오[열쇠:동사]+기[이름:기]+를[목적격:를]\n"}},
		{"대곤 (大棍)을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"대곤(大)을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"대곤(大棍,)을 맞았다.", nil, []string{"\n대[열쇠:동사]+곤[맺음:곤(고는)]\n"}},
		{"군사를 운용한다(行陣和睦).", []string{"(ㄴ)다[맺음:ㄴ다]"}, []string{"\n운용[열쇠:체언] 한다\n"}},
		{"그 보장률(保障率)이 낮다.", nil, []string{"\n보장[열쇠:체언]+률[접미:률]\n"}},
		{"그 보장률이 낮다.", nil, []string{"\n보장[열쇠:체언]+률[접미:률]+이[주격:이]\n"}},
	} {
		lines := streamText(r, c.sentence)
		var shown []string
		for _, sg := range r.JSONSentence(c.sentence, 0).Segments {
			for _, w := range sg.Words {
				shown = append(shown, jsonWordShown(w)...)
			}
		}
		stream := strings.Join(shown, "\n")
		for _, g := range c.gone {
			if strings.Contains(lines, g) || strings.Contains(stream, g) {
				t.Errorf("%s: %s stands\n%s\n%s", c.sentence, g, lines, stream)
			}
		}
		for _, k := range c.kept {
			if !strings.Contains(lines, k) {
				t.Errorf("%s: want %q\n%s", c.sentence, k, lines)
			}
		}
		if c.gone == nil && !strings.Contains(stream, strings.TrimSpace(strings.Split(c.kept[0], "\t")[0])) {
			t.Errorf("%s: the stream lost the marker\n%s", c.sentence, stream)
		}
	}
}
