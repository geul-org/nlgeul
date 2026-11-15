//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestComplementRuleKeeps — complementRule takes no candidate away and adds none: every word of the sentences of the other tests has as many candidates, with the same names, as in the stream with no rule
package token

import "testing"

func TestComplementRuleKeeps(t *testing.T) {
	for _, s := range []string{"물이 얼음이 되었다", "그는 학생이 아니다", "학생이 밥을 먹었다", "반장이 누가 됐어", "누가 반장이 됐어", "커서 뭐가 될래", "반장이 누구를 됐어", "반장이 누가 먹었어"} {
		complementKept(t, s)
	}
}
