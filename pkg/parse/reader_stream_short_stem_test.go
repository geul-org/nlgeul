//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamShortStem — the stream reads the shortened forms of a stem and of a fused word: 자랑스런 = 자랑스럽+(ㄴ) (스러운 without its 우, before the modifying ㄴ only: 자랑스런데 has it as the modifying form before 데, and 자랑스러면 has no 자랑스럽), 파라냐고 = 파랗+냐고 and 좋으냐고 = 좋+으+냐고 (냐고 takes the link vowel 으), 이걸로 = 이것+으로, and 걸로 and 데다 right after a modifying form = 거+(으)로 and 데+(에)다
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamShortStem(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-under.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, has, not string
	}{
		{"자랑스런 바람이 분다.", "\n자랑스럽[열쇠:형용사]+(ㄴ)[꾸밈:ㄴ]\n", "자랑스럽[열쇠:형용사]+(ㄴ)다"},
		{"자랑스러면 좋다.", "자랑", "자랑스럽[열쇠:형용사]+면"},
		{"그것이 파라냐고 한다.", "\n파랗[열쇠:형용사]+냐고[맺음:냐고]\n", "파랗[열쇠:형용사]+으"},
		{"그것이 좋으냐고 한다.", "\n좋[열쇠:형용사]+으[이음:으]+냐고[맺음:냐고]\n", "좋으[열쇠"},
		{"이걸로 씻는다.", "\n이것[열쇠:체언]+으로[부사격:으로]\n", "이것[열쇠:체언]+(으)로"},
		{"그걸로 씻는다.", "\n그것[열쇠:체언]+으로[부사격:으로]\n", "그것[열쇠:체언]+(으)로"},
		{"그렇게 하는 걸로 한다.", "\n거[의존명사:거]+(으)로[부사격:으로]\n", "\n걸[열쇠:체언]+로[부사격:로]\n"},
		{"바람이 부는 데다 비가 온다.", "\n데[의존명사:데]+(에)다[부사격:에다]\n", "데[의존명사:데]+(에)다가"},
		{"바람이 부는 데다가 비가 온다.", "\n데[의존명사:데]+(에)다가[부사격:에다가]\n", "데[의존명사:데]+다가"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || strings.Contains(got, c.not) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
