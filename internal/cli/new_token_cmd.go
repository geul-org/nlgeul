//ff:func feature=cli type=command control=sequence
//ff:what newTokenCmd — kopas token: read a token (syllable) stream into words with a graph for every marker reading, and print contests where the readings do not agree
package cli

import (
	"os"

	"github.com/geul-org/koparse/internal/token"
	"github.com/spf13/cobra"
)

// newTokenCmd is `kopas token "<문장>"`. The word table comes from
// --words, else from the environment variable KOPAS_WORDS.
func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `token "<문장>"`,
		Short: "토큰(음절) 스트림을 표지별 판정 그래프로 읽어 어절로 끊어 올린다(실험)",
		Long: `문장을 띄어쓰기와 문장 부호(확정 경계)에서 단위로 나누고, 단위마다 토큰(음절)을
차례로 읽으며 색인에서 표지(조사, 선어말 칸, 맺음)와 낱말을 찾아 표지 풀이마다
판정 그래프를 돌린다. 살아남은 읽기들이 모두 합의한 자리에서 끊어 한 줄씩 찍는다.
조각은 "+"로 잇고 조각 뒤 대괄호에 그래프 이름을 단다. 줄기에 합쳐진 조각은
줄기를 본디 꼴로 되돌리고 조각을 괄호에 적는다(마신다 → 마시+(ㄴ)다, 써 → 쓰+(어)).
읽기가 갈리면 "경합"으로 갈래를 줄마다 들여 찍고 성립(조각 verdict의 최솟값)과
선호(관측 횟수에서 나온 값)를 적는다. 받침으로 끝난 이름 뒤의 "이"는 접미로도
읽는다(영숙+이+가). 체언 뒤의 서술격 이다는 줄기 "이"를 따로 읽는다
(그림+이[서술격:이]+다, 학생+이+(ㅂ)니다, 문제+이+(었)+다). 모음 뒤에서 빠진 이는
괄호에 적는다(학교+(이)[서술격:이(빠짐)]+다). 두 조각이 합쳐진 글자는 풀어 읽는다:
셨·셔·십의 높임 시(가+시+(었)+다, 받침 뒤에는 읽+으+시+(었)+다), 잖(지+(않)), 죠(지+(요)).
앞말이 정해진 꼴(이게, 이건, 이걸, 뭘, 누가, 여긴, 우린, 그땐 들)은 낱말과 그 아래의
조사로 푼다(이것+(이)[주격:이]). 라틴 문자와 숫자가 이어진 열은 조각 하나이고, 점(.)은
풀이마다 그래프가 있다: 파일명·도메인·소수점·버전의 점은 양쪽 열을 잇고
(task.py[파일명|도메인]+에), 마침표·줄임표는 앞 어절을 닫는다. 숫자 열은 수이고
(3[수]), 숫자 사이의 기호는 풀이마다 그래프가 판정해 선 풀이를 이름에 단다
(1,000[수], 10:30[시각|비율], 2026-11-10[날짜], 10~20[범위], 11/10[날짜|나눔],
010-1234-5678[전화], 192.168.0.1[버전|IP]). 수의 맨 앞 기호(-3[부호|수], $40[돈|수])도
수에 든다. 수 바로 뒤의 단위 표 낱말은 단위다(3[수]+개[단위:개]+를, 3+번+째[차례:째],
10+여[어림:여]). 단위 앞의 수는 숫자 열, 한글 수사(한, 두, 삼, 십, 몇, 여러 들),
표의 영문 글자(N, x, k)이고 숫자 뒤에는 큰 수(만, 천, 억)가 이어진다
(한[수]+개, 7[수]+만[큰 수:만]+장, N[수]+개). 단위 뒤에는 조사 사슬, 이다, 째가 오거나
단위가 끝나야 한다. 주소, 전자우편, 경로, 명령의 옵션으로 시작하는 단위는 그 부분이
조각 하나다(~/a/b[경로]+를, --xx[옵션]). 두 영숫자 열 사이의 빗금과 붙임표는 이름을
잇는다(a/b[경로], h-Cat[붙임]). 이음 으(읽+으+시+고)와 연결 어미 뒤의 보조사(보+지+도)는
조각으로 읽는다.
문장성분은 정하지 않는다.

낱말 표(--words 또는 환경 변수 KOPAS_WORDS)가 없으면 문법 표에 적힌 닫힌 집합
낱말만 안다. 낱말 표는 탭으로 나눈 칸 넷(검색 키, 갈래, 불규칙 유형, 횟수)의
줄이고 "#"으로 시작하는 줄은 건너뛴다. 표가 있으면 표에 있는 줄기를 아는 낱말로
치고(1겹), 줄기의 갈래가 뒤의 표지와 맞는지와 관측 횟수·길이로 그 세기를 보고
(2겹), 경계·흔적 없는 조각·조각 수에 값을 매긴다(3겹). 새 낱말에 기댄 읽기도
낮은 성립으로 묶음에 남는다. 낱말 속에도 흔한 합쳐진 글자는 낱말 표가 있을 때만,
앞말이 표의 용언 줄기(줄어든 조사는 체언)일 때 다른 풀이와 나란히 읽는다:
높임 시가 든 신·실·심·세(가+시+(ㄴ)다, 가+시+(어)+요), 인용이 줄어든 대·래·재·단·란·잔
(간+대[맺음:ㄴ대(ㄴ다고 해)]), 줄어든 선·론·진·긴·길(여기+선[부사격:선(에서는)]).
낱말 표가 있으면 불규칙 용언의 바뀐 줄기를 본디 줄기로 되돌리고(듣+었+다, 돕+(아),
모르+(아), 살+(ㄴ)다), 받침 하나뿐인 맺음을 용언 줄기 밑에서 읽고(예쁘+(ㄴ)), 본용언에
붙여 쓴 보조 용언을 읽는다(지우+(어)+지+(ㄴ)다). 단위의 읽기 가운데 1등과 차이가 큰
것은 올리지 않고 재워 둔다(--asleep로 찍는다).`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := cmd.Flags()
			trace, _ := f.GetBool("trace")
			cuts, _ := f.GetBool("cuts")
			stdin, _ := f.GetBool("stdin")
			words, _ := f.GetString("words")
			layers, _ := f.GetInt("layers")
			asleep, _ := f.GetBool("asleep")
			if words == "" {
				words = os.Getenv("KOPAS_WORDS")
			}
			sentences, err := tokenInput(args, stdin, cmd.InOrStdin())
			if err != nil {
				return err
			}
			r, err := token.New(token.Options{Words: words, Layers: layers, Asleep: asleep})
			if err != nil {
				return err
			}
			runToken(cmd.OutOrStdout(), r, sentences, trace, cuts)
			return nil
		},
	}
	cmd.Flags().Bool("trace", false, "단위마다 걸린 토큰 열, 그래프, verdict, 기각 이유를 함께 찍는다")
	cmd.Flags().Bool("cuts", false, "측정용: 단위마다 살아남은 읽기들의 어절 끝 자리(토큰 수)만 찍는다")
	cmd.Flags().Bool("stdin", false, "표준 입력의 줄마다 문장 하나로 읽고, 문장 사이에 빈 줄을 찍는다")
	cmd.Flags().String("words", "", "낱말 표 파일(없으면 환경 변수 KOPAS_WORDS, 그것도 없으면 낱말 표 없이)")
	cmd.Flags().Bool("asleep", false, "단위마다 잠든 갈래(1등과의 차이가 커서 올리지 않고 재워 둔 읽기)도 찍는다")
	cmd.Flags().Int("layers", 0, "측정용: 낱말 표를 몇 겹까지 쓸지(1 여부, 2 갈래 맞춤, 3 구조 값; 0은 전부)")
	return cmd
}
