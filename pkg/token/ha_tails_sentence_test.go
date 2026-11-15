//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestHaTailsSentence — the sentences of clause 40 read as a stream (the best candidate of each word, by lambda): 원치, 가능케, 불쌍타, 달성코자, 단언컨대 are read as the stem with 하 and the plain closing's name on the written letters, in one word; 분발토록 keeps the noun with the particle 토록 first and has the clause 40 reading after it; nouns with a particle right after (위치를, 정치는, 대치는, 강타를, 일치는) keep their reading and have no clause 40 reading; 영원토록 keeps the adverb first; the known harm: the dependent noun 만치 alone is read first as 만하+치
package token

import (
	"slices"
	"strings"
	"testing"
)

func TestHaTailsSentence(t *testing.T) {
	r := streamReader(t, "words-ha.tsv")
	for _, c := range []struct {
		sentence, word, first, also string
		none                        bool
	}{
		{"나는 그것을 원치 않는다.", "원치", "원하[열쇠:동사]+치[맺음:지]", "", false},
		{"모두가 가능케 했다.", "가능케", "가능하[열쇠:형용사]+케[맺음:게]", "", false},
		{"불쌍타 그 사람.", "불쌍타", "불쌍하[열쇠:형용사]+타[맺음:다]", "", false},
		{"목표를 달성코자 한다.", "달성코자", "달성하[열쇠:동사]+코자[맺음:고자]", "", false},
		{"단언컨대 그는 옳다.", "단언컨대", "단언하[열쇠:동사]+컨대[맺음:건대]", "", false},
		{"힘껏 분발토록 했다.", "분발토록", "분발[열쇠:체언]+토록[부사격:토록]", "분발하[열쇠:동사]+토록[맺음:도록]", false},
		{"위치를 바꿨다.", "위치를", "위치[열쇠:체언]+를[목적격:를]", "", true},
		{"정치는 어렵다.", "정치는", "정치[열쇠:체언]+는[보조사:는]", "", true},
		{"대치는 계속되었다.", "대치는", "대치[열쇠:체언]+는[보조사:는]", "", true},
		{"강타를 맞았다.", "강타를", "강타[열쇠:체언]+를[목적격:를]", "", true},
		{"일치는 없었다.", "일치는", "일치[열쇠:체언]+는[보조사:는]", "", true},
		{"영원토록 남는다.", "영원토록", "영원토록[열쇠:부사]", "영원하[열쇠:형용사]+토록[맺음:도록]", false},
		{"그럴 만치 컸다.", "만치", "만하[열쇠:형용사]+치[맺음:지]", "만치[열쇠:체언]", false},
	} {
		w, ok := jsonWordAt(r.JSONSentence(c.sentence, 0), c.word)
		if !ok {
			t.Errorf("%s: no word %s", c.sentence, c.word)
			continue
		}
		got := jsonWordShown(w)
		if got[0] != c.first || (c.also != "" && !slices.Contains(got, c.also)) {
			t.Errorf("%s: %v, want %s first (and %s)", c.sentence, got, c.first, c.also)
		}
		if c.none && strings.Contains(strings.Join(got, "\n"), "하[열쇠:") {
			t.Errorf("%s: a particle right after, yet a clause 40 reading: %v", c.sentence, got)
		}
	}
}
