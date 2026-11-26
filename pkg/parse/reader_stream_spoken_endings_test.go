//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamSpokenEndings — with the word table of testdata, the endings added at the end of the closing table are read after a stem of the table: 니깐 (with the link vowel 으 after a consonant, and after a dropped copula stem), 구료·는구료·로구료, 고나·로고나, 더만, 라던가 after the copula stem, 리라고·으리라고, 게마는, 다마는·ㄴ다마는·는다마는, 련마는, ㄹ련·을련, 예요 after a final consonant, the gated 련·으련 (not without a word table), and 느냐 right after the closing 려·으려 (closing-chains.tsv)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamSpokenEndings(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for in, want := range map[string]string{
		"가니깐":    "가[열쇠:동사]+니깐[맺음:니깐]",
		"먹으니깐":   "먹[열쇠:동사]+으[이음:으]+니깐[맺음:니깐]",
		"학교니깐":   "학교[열쇠:체언]+(이)[서술격:이(빠짐)]+니깐[맺음:니깐]",
		"가시구료":   "가[열쇠:동사]+시[선어말:시]+구료[맺음:구료]",
		"먹는구료":   "먹[열쇠:동사]+는구료[맺음:는구료]",
		"마당이로구료": "마당[열쇠:체언]+이[서술격:이]+로구료[맺음:로구료]",
		"작고나":    "작[열쇠:형용사]+고나[맺음:고나]",
		"마당이로고나": "마당[열쇠:체언]+이[서술격:이]+로고나[맺음:로고나]",
		"먹더만":    "먹[열쇠:동사]+더만[맺음:더만]",
		"마당이라던가": "마당[열쇠:체언]+이[서술격:이]+라던가[맺음:라던가]",
		"사과라던가":  "사과[열쇠:체언]+(이)[서술격:이(빠짐)]+라던가[맺음:라던가]",
		"가리라고":   "가[열쇠:동사]+리라고[맺음:리라고]",
		"먹으리라고":  "먹[열쇠:동사]+으리라고[맺음:으리라고]",
		"가게마는":   "가[열쇠:동사]+게마는[맺음:게마는]",
		"작다마는":   "작[열쇠:형용사]+다마는[맺음:다마는]",
		"간다마는":   "가[열쇠:동사]+(ㄴ)다마는[맺음:ㄴ다마는]",
		"먹는다마는":  "먹[열쇠:동사]+는다마는[맺음:는다마는]",
		"가련마는":   "가[열쇠:동사]+련마는[맺음:련마는]",
		"먹으련마는":  "먹[열쇠:동사]+으련마는[맺음:으련마는]",
		"갈련":     "가[열쇠:동사]+(ㄹ)련[맺음:ㄹ련]",
		"먹을련":    "먹[열쇠:동사]+을련[맺음:을련]",
		"마당예요":   "마당[열쇠:체언]+예요[맺음:예요]",
		"가련":     "가[열쇠:동사]+련[맺음:련]",
		"먹으련":    "먹[열쇠:동사]+으련[맺음:으련]",
		"가려느냐":   "가[열쇠:동사]+려[맺음:려]+느냐[맺음:느냐]",
		"먹으려느냐":  "먹[열쇠:동사]+으려[맺음:으려]+느냐[맺음:느냐]",
	} {
		if got := streamText(r, in); !strings.Contains(got, want) {
			t.Errorf("%s: no %s\n%s", in, want, got)
		}
	}
	for in, not := range map[string]string{
		"작는구료": "작[열쇠:형용사]+는구료[맺음:는구료]",
		"작을련":  "작[열쇠:형용사]+을련[맺음:을련]",
		"가고느냐": "+고[맺음:고]+느냐[맺음:느냐]",
	} {
		if got := streamText(r, in); strings.Contains(got, not) {
			t.Errorf("%s: %s\n%s", in, not, got)
		}
	}
	bare, _ := New(Options{})
	if got := streamText(bare, "가련"); strings.Contains(got, "[맺음:련]") {
		t.Fatal("련 without a word table\n" + got)
	}
}
