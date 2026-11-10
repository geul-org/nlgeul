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
읽는다(영숙+이+가). 문장성분은 정하지 않는다.

낱말 표(--words 또는 환경 변수 KOPAS_WORDS)가 없으면 문법 표에 적힌 닫힌 집합
낱말만 안다. 낱말 표는 탭으로 나눈 칸 넷(검색 키, 갈래, 불규칙 유형, 횟수)의
줄이고 "#"으로 시작하는 줄은 건너뛴다. 표가 있으면 표에 있는 줄기를 아는 낱말로
치고(1겹), 줄기의 갈래가 뒤의 표지와 맞는지 보고(2겹), 경계·흔적 없는 조각·조각
수에 값을 매긴다(3겹). 아는 낱말만으로 된 읽기가 서면 새 낱말에 기댄 읽기는 뺀다.`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			f := cmd.Flags()
			trace, _ := f.GetBool("trace")
			cuts, _ := f.GetBool("cuts")
			stdin, _ := f.GetBool("stdin")
			words, _ := f.GetString("words")
			layers, _ := f.GetInt("layers")
			if words == "" {
				words = os.Getenv("KOPAS_WORDS")
			}
			sentences, err := tokenInput(args, stdin, cmd.InOrStdin())
			if err != nil {
				return err
			}
			r, err := token.New(token.Options{Words: words, Layers: layers})
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
	cmd.Flags().Int("layers", 0, "측정용: 낱말 표를 몇 겹까지 쓸지(1 여부, 2 갈래 맞춤, 3 구조 값; 0은 전부)")
	return cmd
}
