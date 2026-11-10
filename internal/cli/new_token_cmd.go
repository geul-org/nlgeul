//ff:func feature=cli type=command control=sequence
//ff:what newTokenCmd — kopas token: read a token (syllable) stream into words with a graph for every marker reading, and print contests where the readings do not agree
package cli

import "github.com/spf13/cobra"

// newTokenCmd is `kopas token "<문장>"`.
func newTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `token "<문장>"`,
		Short: "토큰(음절) 스트림을 표지별 판정 그래프로 읽어 어절로 끊어 올린다(실험)",
		Long: `문장을 띄어쓰기와 문장 부호(확정 경계)에서 단위로 나누고, 단위마다 토큰(음절)을
차례로 읽으며 색인에서 표지(조사, 선어말 칸, 맺음)와 닫힌 집합 낱말을 찾아 표지 풀이마다
판정 그래프를 돌린다. 살아남은 읽기들이 모두 합의한 자리에서 끊어 한 줄씩 찍는다.
조각은 "+"로 잇고 조각 뒤 대괄호에 그래프 이름을 단다. 읽기가 갈리면 "경합"으로
갈래를 줄마다 들여 찍고 성립(조각 verdict의 최솟값)과 선호(관측 횟수에서 나온 값)를 적는다.
문장성분은 정하지 않는다. 열쇠(낱말) 표는 없고 표에 적힌 닫힌 집합 낱말만 안다.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trace, _ := cmd.Flags().GetBool("trace")
			cuts, _ := cmd.Flags().GetBool("cuts")
			return runToken(cmd.OutOrStdout(), args[0], trace, cuts)
		},
	}
	cmd.Flags().Bool("trace", false, "단위마다 걸린 토큰 열, 그래프, verdict, 기각 이유를 함께 찍는다")
	cmd.Flags().Bool("cuts", false, "측정용: 단위마다 살아남은 읽기들의 어절 끝 자리(토큰 수)만 찍는다")
	return cmd
}
