# rule

toulmin 그래프의 규칙 함수(`func(Context, Specs) (float64, any)`) 가운데 기준(Spec)만 바꿔 여러 노드와 여러 그래프가 함께 쓰는 것을 모았다. 함수가 읽는 사실 타입과 컨텍스트 열쇠도 여기 있다. `internal/`에 기대지 않는다.

## 사실과 열쇠

- `KeyWord` → `*Word`(어절: 글, 조사 사슬 `Chain`, 꼬리 풀이들 `[]Tail`, 가장 긴 꼬리, 마지막 어절인지, 의존명사로 시작하는지, 문장 `*Sentence`).
- `KeyTail` → `Tail`(판정 중인 꼬리 풀이 하나).
- 읽는 쪽이 이 타입을 채워 컨텍스트에 넣는다. 규칙은 `WordOf`, `TailOf`로 꺼낸다.
- `KeyAround` → `Around`(문장 안에서 판정하는 후보와 그 둘레: 후보 자신 `Self`, 앞 어절들 `Before`, 뒤 어절들 `After`. 앞뒤 모두 가까운 것부터이고 저마다 `Near`다). `Near`는 어절 하나를 그 후보 하나가 읽은 대로 적는다: 찍힌 줄기 `Stem`, 줄기의 이름 `Kind`(`열쇠:체언`, `보조 용언:있`. 새 낱말은 ""), 마지막 조각의 이름 `Last`(줄기뿐이면 ""), 서술어인가 `Pred`(맺음을 지녔다. 읽는 쪽이 채운다), 표지들의 이름 `Marks`(줄기 뒤 조각의 이름을 차례대로 `+`로 이은 한 문자열. `MarksOf`가 낱낱으로 돌려준다. 두 `Near`를 `==`로 견줄 수 있게 문자열 하나로 둔다). 앞 어절과 빈칸 없이 붙어 있는가 `Glued`(붙여 쓴 어절. 읽는 쪽이 채운다). 앞뒤 어절은 그 어절의 1위 후보로 적는다. 토큰 층의 문장 단위 규칙이 채우고(`pkg/token/README.md`의 "스트림 인터페이스"), 규칙은 `AroundOf`로 꺼낸다(없으면 빈 `Around`).
- `Near`의 `Glued`(앞 어절과 빈칸 없이 붙었다), `Nominal`(체언으로 끝난다: 마지막 조각이 체언 줄기·차례이거나 수 뒤 단위다), `Text`(어절의 글자), `Modifiable`(꾸밈꼴로 끝나는 어절이 바로 앞에 설 수 있다: 줄기가 체언·관형사이거나 꾸밈꼴·이름꼴로 끝난다), `Signs`(이름 없는 기호만의 열이다: 글에서 가린 이름 `○○`, `□□□`)도 읽는 쪽이 채운다. `AdnominalAll`(어절의 후보가 모두 관형어다: 꾸밈 맺음으로 끝나거나 표지 없이 홀로 선 관형사)과 `Unmodifiable`(관형어가 바로 앞에 설 수 없다: 새 낱말 읽기를 빼고 아는 후보가 모두 꾸밈꼴·이름꼴이 아닌 용언·부사·감탄사이거나, 문장을 끝내는 부호다)은 1위 하나가 아니라 어절의 후보 모두에서 읽는 쪽이 채운다. 뒤 어절의 `Glued`는 그 어절이 판정하는 후보의 어절에 붙어 있다는 뜻이다. `AroundIs`가 묻는다.

## 기준

| 기준 | 쓰임 |
|---|---|
| `Claim{Name, Concludes}` | 주장 노드의 이름과 결론. 같은 함수를 여러 주장이 쓸 수 있게 하고, 평가 결과의 이름에서 `ClaimOf`로 되찾는다. 토큰 층은 반박과 예외 노드에도 붙여 그 노드의 이름(기각 이유)을 싣는다 |
| `Bits{Of, Mask}` | 격(`Case…`), 맺음 갈래(`Kind…`), 어절의 사실(`WordLast`, `WordBound`), 조각의 사실(`Piece…`), 기호의 사실(`Mark…`), 글자 갈래(`Side…`), 붙는 조건(`After…`), 문장 안 후보와 둘레의 사실(`Around…`) 들 가운데 어느 것을 볼지 |
| `List{Name, Of, Items}` | 목록과, 그 목록에서 찾을 사실(`List…`: 조각의 꼴, 앞 조각의 이름, 뒤 조각의 이름, 합쳐진 조각이 놓인 음절, 점 오른쪽 마디, 이름의 맨 오른쪽 마디, 어절, 꼬리의 맺음, 문장 안의 후보의 줄기·줄기의 이름·마지막 조각의 이름·어절의 글자, 앞 어절의 마지막 조각의 이름, 뒤 어절의 줄기) |
| `Not{}` | 예/아니오 규칙을 뒤집는다(toulmin에 부정이 없어 방향을 기준에 싣는다). `PieceIs`, `Listed`, `MarkIs`, `MarkBoth`, `StemKind`가 받는다(`Holds`) |
| `Versus{Mine, Rival}` | `Longer`가 견줄 두 길이(`SideChain`, `SideTail`, `SideLongest`) |
| `Shape{Of, Min, Max, Lens, Vals, Head}` | 기호가 잇는 숫자 마디들의 꼴 |
| toulmin `StringListSpec` | `ChainListed`의 조사 목록 |

## 규칙 함수

규칙 함수는 30개다. 값은 사실 하나를 보는 함수면 +1(적용)·-1(적용 안 됨)이고, 관측 횟수를 견주는 함수면 두 횟수 가운데 상대의 몫을 [-1, 1]에 옮긴 값이다. 같은 틀의 판정(참/거짓 사실 하나, 목록에 있다/없다)은 함수 하나가 기준으로 갈라 맡는다.

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `ChainCase` | rule | 사슬의 가장 바깥 격조사가 기준의 격이다 | `Bits`(case) |
| `ChainNoCase` | rule | 격조사 없이 보조사만 붙었다 | 없음 |
| `ChainListed` | rule | 격조사 없는 사슬에 기준 목록의 조사가 있다(사슬의 조각 여럿을 보고 격조사가 없는지도 함께 보아 `Listed`에 넣지 않았다) | 조사 목록 |
| `ClosingKind` | rule | 꼬리의 맺음이 기준의 갈래다 | `Bits`(kind) |
| `WordIs` | except | 어절이 기준의 사실을 갖는다(마지막 어절, 의존명사로 시작) | `Bits`(word) |
| `Longer` | except | 경쟁 풀이가 이 풀이보다 길게 읽는다 | `Versus` |
| `ChainSeenMore` | counter | 꼬리만큼 긴 조사가 관측됐다(횟수의 몫) | 없음 |
| `TailSeenMore` | counter | 사슬만큼 긴 꼬리가 관측됐다(횟수의 몫) | 없음 |
| `Listed` | rule, except, counter | 기준이 가리키는 사실이 기준 목록에 있다. `pkg/parse`는 어절(`ListWord`: 관형사 → 관형어, 접속 부사 → 독립어, 닫힌 집합의 낱말은 쪼개 읽지 않는다)과 꼬리의 맺음(`ListClosing`: 게 → 부사어)을, 토큰 층은 조각과 기호의 사실, 문장 안 후보와 둘레 어절의 사실을 묻는다 | `List`, 뒤집으려면 `Not` |

토큰 열 하나를 표지나 줄기로 판정하는 그래프(컨텍스트 열쇠 `KeyPiece` → `Piece`)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `PieceRead` | rule, counter | 색인에서 걸린 토큰 열이 이 자리에 있다(주장, 늘 서는 반박) | 없음 |
| `PieceIs` | except, counter | 조각이 기준의 사실을 갖는다(`Piece…`): 앞에 줄기가 있다(`PieceStem`), 바로 앞이 수다(`PieceNum`), 그 수가 낱말 표 없이 홀로 선 한자어 한 음절 수사다(`PieceLone`), 수와 합쳐 통째로 낱말이다(체언 뒤 접미는 앞 줄기와 합쳐 낱말 표의 낱말이다, `PieceWhole`), 앞 어절이 꾸밈 맺음으로 끝난다(`PieceAdn`), 그 낱말 자신이 꾸밈꼴·이름꼴이다(`PieceTurned`), 경계 뒤의 새 낱말이 표지로 읽힌다(`PieceMarked`), 뒤에 이어지는 조각이 없다(`PieceLast`), 줄기가 어/아 앞에서 끝 모음을 지킨다(`PieceKeeps`), 줄기의 낱말이 단위 전체다(`PieceSole`), 앞 어절이 관형사다(`PieceDet`), 조각이 줄기 바로 뒤에서 단위를 끝내고 바로 뒤에서 마침표·물음표·느낌표가 단위를 닫는다(`PieceClosed`), 되돌린 음절이 글자에 그대로 있다(합쳐진 조각이 흔적을 남기지 않았다, `PieceSame`), 줄기 또는 표지의 글자가 다른 낱말·맺음의 틀린 표기다(`PieceMisspelt`), 바로 뒤에 보조사 요가 오고 조각이 줄기 바로 뒤에서 받침 없는 한글 음절 다음에 서거나 줄기 위의 접미 이 바로 뒤에 선다(`PieceYoOpen`), 바로 앞의 수가 접두사 `제`로 시작한다(`PieceOrder`), 새 낱말의 글자가 영문 글자 열 바로 뒤에 띄어쓰기 조항이 띄어 쓰는 낱말(제48항 호칭어·관직명, 제42항 의존 명사) 하나만 품는다(`PieceApart`), 앞 어절이 어/아 맺음으로 끝나고 조각이 서술어 꼬리가 붙은 보조 용언의 줄기다(`PieceConn`), 새 낱말이 영숫자 열과 한글이 맞닿은 자리를 넘고 그 자리에서 가른 두 쪽이 모두 아는 낱말이나 영숫자 열로 읽힌다(`PieceChange`), 단위의 글자(머리 뒤)가 한글 음절이고 그 바로 뒤의 소괄호가 같은 수의 한자만 품는다(한자 병기, `PieceGlossed`), 수 바로 뒤에서 시작하는 줄기의 글자가 단위 표의 단위와 단위 뒤 낱말이다(`5년간`, `PieceUnitWord`), 바로 앞 어절이 꾸밈 맺음으로 끝나고 이 줄기가 그 꾸밈 위에 서는 의존 명사다(`것`, `거`, `PieceBound`), 새 낱말의 글자가 꾸미는 말과 그 꾸밈을 받는 아는 낱말을 붙여 쓴 것이다(`못와`, `두손`, `PieceModifier`), 단위를 여는 줄기이고 그 낱말이 어/아 맺음으로 끝나며 바로 뒤에 보조 용언의 줄기가 붙어 있다(`가봐`의 `가`, `PieceMain`) | `Bits`(piece), 뒤집으려면 `Not` |
| `Listed` | except, counter | 조각의 꼴(`ListForm`), 앞 조각의 이름(`ListBefore`), 뒤 조각의 이름(`ListAfter`), 합쳐진 조각이 놓인 음절(`ListBase`)이 기준 목록에 있다 | `List`, 뒤집으려면 `Not` |
| `PieceAfter` | except, counter | 앞 토큰의 끝이 붙는 조건과 안 맞는다(`Fits`) | `Bits`(after) |
| `PieceRight` | except | 조각 바로 뒤에 선 것이 기준의 것이다(조사, 그 밖의 글자) | `Bits`(right) |
| `PieceClash` | except, counter | 어/아/여로 시작하는 조각이 앞 음절의 모음과 어긋난다(`Harmony`) | 없음 |
| `ChainLinks` | except | 조사 사슬에서 앞 조사의 갈래가 뒤 조사의 갈래를 받지 않는다(`Links`) | 없음 |
| `SlotOrder` | except | 선어말 칸의 순서가 어긋난다(`SlotBefore`) | 없음 |
| `StemTail` | except | 줄기의 갈래가 뒤에 붙은 것과 맞지 않는다 | `Bits`(tail) |
| `StemKind` | except | 줄기가 표에서 기준 갈래 가운데 하나다(앞말이 표의 용언 줄기다, 표가 아는 낱말이다) | `Bits`(kinds), 뒤집으려면 `Not` |
| `StemCoda` | except | 뒤에 표지가 붙은 줄기의 마지막 음절이 기준 받침으로 끝난다 | `Bits`(jong) |
| `SeenFew` | counter | 이 낱말은 관측이 적다(적은 정도. 횟수 0은 한 번과 같고, 횟수를 모르면 걸리지 않는다. 넉넉한 횟수가 1 이하면 걸리지 않는다) | `Bits`(enough) |
| `ShortKey` | counter | 이 낱말은 짧아 우연히 겹치기 쉽다 | `Bits`(syllables) |

`Piece`의 `Before`와 `After`는 앞 조각과 뒤 조각의 이름이다. 서술격 `이`의 `After`는 이 바로 뒤에 선 것의 이름이다: 꼬리가 맺음으로 시작하면 그 맺음, 선어말이 먼저 서면 그 선어말이다.

도움 함수: `PieceOf`, `MarkOf`, `AroundOf`, `MarksOf`, `Applies`(예/아니오 → ±1), `Holds`(기준에 `Not`이 있으면 뒤집은 ±1), `Share`(횟수의 몫), `WordOf`, `TailOf`, `ClaimOf`.

기호 하나를 풀이별로 판정하는 그래프(컨텍스트 열쇠 `KeyMark` → `Mark`)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `MarkIs` | except | 기호가 기준의 사실을 갖는다(`Mark…`): 같은 기호가 둘 이상 이어진다(`MarkRepeated`), 왼쪽에 글자가 붙어 있다(`MarkAfterWord`), 이름의 마디가 모두 숫자다(`MarkNumeric`), 숫자 바로 뒤에 있다(`MarkAfterDigit`), 단위 전체다(`MarkAlone`), 읽는 글자 뒤에 있다(`MarkEnds`), 읽는 글자가 바로 뒤에서 시작한다(`MarkLeads`), 입력 줄의 맨 처음이다(`MarkHeads`), 바로 뒤가 빈칸이다(`MarkSpaced`), 같은 종류의 여는 기호가 열려 있을 수 있다(`MarkOpened`), 앞뒤가 모두 빈칸이다(`MarkApart`), 붙임표나 등호와 맞닿아 있다(`MarkArrow`), 자기 말고 다른 문장부호나 기호가 바로 옆에 있다(`MarkBySign`), 어느 방식에서든 같은 종류의 여는 기호가 열려 있을 수 있다(`MarkOpenedAny`), 기호가 잇는 숫자 마디의 마지막 바로 뒤에 마침표가 붙어 있다(`MarkStopped`: `1.28.은`의 날짜), 문장을 끝내는 부호(. ? ! …)가 바로 앞에 있다(`MarkAfterStop`: `다.』`의 닫는 괄호). `MarkAlone`부터는 글자만으로 알 수 없어 읽는 쪽이 `Mark.Facts`에 싣는다 | `Bits`(mark), 뒤집으려면 `Not` |
| `MarkBoth` | except | 기호의 양쪽이 모두 기준의 글자 갈래다 | `Bits`(sides), 뒤집으려면 `Not`(한쪽이라도 아니다) |
| `MarkRight` | except | 기호의 오른쪽이 기준의 글자 갈래가 아니다 | `Bits`(sides), 뒤집으려면 `Not`(오른쪽이 그 글자 갈래다) |
| `PartsShape` | except | 기호가 잇는 숫자 마디들이 기준 꼴이 아니다(`FitsShape`) | `Shape` |
| `Listed` | except | 점 오른쪽 마디(`ListExt`), 이름의 맨 오른쪽 마디(`ListLast`)가 기준 목록에 있다 | `List`, 뒤집으려면 `Not` |

어절의 후보 하나를 문장 안의 둘레 어절과 함께 판정하는 그래프(컨텍스트 열쇠 `KeyAround` → `Around`, 토큰 층의 문장 단위 규칙)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `Listed` | rule, except | 후보의 줄기(`ListSelfStem`), 그 줄기의 이름(`ListSelfKind`), 후보의 마지막 조각의 이름(`ListSelfLast`), 바로 앞 어절의 마지막 조각의 이름(`ListPrevLast`), 바로 앞 어절의 줄기 이름의 갈래(`ListPrevKind`: `닫기:1`의 `닫기`), 바로 앞 어절의 줄기(`ListPrevStem`), 바로 앞 어절의 줄기 이름 통째(`ListPrevName`: `열쇠:동사`), 바로 뒤 어절의 줄기(`ListNextStem`), 바로 뒤 어절의 마지막 조각의 이름(`ListNextLast`), 그 뒤 어절의 줄기(`ListNext2Stem`), 그 뒤 어절의 마지막 조각의 이름(`ListNext2Last`), 후보 어절의 글자(`ListSelfText`: 하다+ㄴ의 `한`. 읽는 쪽이 `Self.Text`에 채운다)가 기준 목록에 있다. 앞이나 뒤에 어절이 없으면 그 사실은 ""다(`nearAt`) | `List`, 뒤집으려면 `Not` |
| `Listed` | rule, except | 바로 뒤 어절의 줄기 이름의 갈래(`ListNextKind`: `닫기:1`의 `닫기`, `ListPrevKind`의 짝)가 기준 목록에 있다. 뒤에 어절이 없으면 ""다 | `List`, 뒤집으려면 `Not` |
| `SelfHas` | rule, except | 후보의 표지(줄기 뒤 조각) 가운데 하나의 이름이 기준 목록에 있다(`에서는`의 `부사격:에서`). 볼 사실은 `ListSelfMarks`로 적는다(`Listed`는 이 사실을 읽지 않는다). 증거는 목록에 있던 이름 | `List`, 뒤집으려면 `Not` |
| `PrevHas` | rule, except | 후보 바로 앞 어절의 표지(줄기 뒤 조각) 가운데 하나의 이름이 기준 목록에 있다(`있기는`의 `이름:기`). 볼 사실은 `ListPrevLast`로 적는다(앞 어절. `Listed`는 표지를 낱낱이 읽지 않는다). 앞 어절이 없으면 없다. 증거는 목록에 있던 이름 | `List`, 뒤집으려면 `Not` |
| `AroundIs` | rule, except | 후보와 둘레가 기준의 사실을 갖는다(`Around…`): 후보의 어절이 바로 앞 어절에 빈칸 없이 붙어 있다(`AroundGlued`), 바로 앞 어절이 체언으로 끝난다(`AroundAfterNominal`), 바로 뒤 어절이 꾸밈꼴 뒤에 설 수 있다(`AroundBeforeModifiable`), 바로 뒤 어절이 빈칸 없이 붙어 있다(`AroundBeforeGlued`: 그 어절의 `Glued`), 바로 앞 어절이 이름 없는 기호만의 열이다(`AroundAfterSigns`: 그 어절의 `Signs`, 가린 이름 `○○`), 바로 앞 어절의 후보가 모두 관형어다(`AroundAfterAdnominal`: 그 어절의 `AdnominalAll`), 바로 뒤 어절에 관형어가 설 수 없다(`AroundBeforeUnmodifiable`: 그 어절의 `Unmodifiable`). 앞이나 뒤에 어절이 없으면 그 사실은 없다 | `Bits`(around), 뒤집으려면 `Not` |
| `KindListed` | rule, except | `Listed`와 같은 사실(`ListSelfKind`, `ListSelfLast`, `ListPrevLast` …)에서 쌍점 앞의 갈래(`꾸밈:ㄹ`의 `꾸밈`, `의존명사:수`의 `의존명사`)가 기준 목록에 있다. 쌍점이 없는 이름은 이름 자체가 갈래다. 증거는 그 갈래 | `List`(갈래 목록), 뒤집으려면 `Not` |
| `Headed` | rule, except | 후보 앞에 어절이 없거나(문장 첫머리), 바로 앞 어절의 줄기 이름의 갈래가 기준 목록에 있다(여닫는 기호는 줄기 자리에 서서 `열기:1`이 줄기 이름이다). 줄기 이름이 없는 새 낱말은 첫머리가 아니다(`ListPrevKind`만으로는 새 낱말과 첫머리가 같은 ""라 가를 수 없다). 증거는 바로 앞 어절의 줄기 이름 | `List`(갈래 목록, 볼 사실은 `ListPrevKind`), 뒤집으려면 `Not` |

## 조건식과 도움 함수 (toulmin 규칙이 아니다)

규칙 함수와 읽기 코드가 같은 문법 조건을 따로 갖지 않도록 조건식을 한 곳에 둔다. `//ff:what` 첫머리에 "조건식" 또는 "도움 함수(규칙 아님)"이라고 적혀 있다.

| 함수 | 뜻 | 쓰는 곳 |
|---|---|---|
| `Fits(prev, after)` | 앞 음절의 끝이 붙는 조건(`After…`: 받침, ㄹ 아닌 받침, 모음이나 ㄹ, 모음, ㄹ 받침, 한글 음절, ㄴ 아닌 받침(`AfterNotN`))과 맞는가 | `PieceAfter`, 읽기 코드 |
| `Follows(left, right)` | 조사 사슬에서 자리 left 뒤에 자리 right가 올 수 있는가(자리 번호) | `pkg/parse`의 읽기 코드 |
| `Harmony(prev, first)` | 줄기 끝 음절 뒤에 어/아/여로 시작하는 조각이 모음에 맞는가(ㅏ·ㅑ·ㅗ 뒤 아, 그 밖 어, 하 뒤 여, ㅆ 받침 뒤 어) | `PieceClash` |
| `FitsShape(parts, shape)`, `InRange` | 숫자 마디들이 꼴(마디 수, 자리 수, 값의 범위, 머리)에 맞는가 | `PartsShape` |
| `Links(left, right)` | 조사 사슬에서 갈래 left 뒤에 갈래 right가 올 수 있는가(갈래 표, `Link…`. 서술격 이다의 줄기 `LinkCopula`는 접미 뒤에만 선다) | `ChainLinks` |
| `SlotOf(r)`, `SlotBefore(slot, next)` | 선어말 칸과 그 순서(시 < 었 < 겠 < 더 < 맺음, 었은 겹침) | `SlotOrder`, 읽기 코드 |
| `BareBases(lead)` | 받침 하나뿐인 맺음을 읽는 음절 | `Listed`의 기준(`ListBase`), 읽기 코드 |
| `ChoOf`, `JongOf`, `JungOf` | 음절의 첫소리·받침·모음 번호 | 읽기 코드 |
| `FusedVowel(r, more)` | 어/아가 합쳐진 음절인가 | 읽기 코드 |
| `FusedPast(r)` | 었/았이 합쳐진 음절인가(받침 ㅆ과 모음 조건) | 읽기 코드 |
| `LeadText(lead)` | 조각의 머리(받침, 합쳐지는 어)를 글자로 | 읽기 코드 |
| `RunRune(r)` | 영문·숫자 열을 잇는 글자인가(라틴 문자, 숫자, 밑줄) | 읽기 코드 |
| `Trim(field)` | 어절 앞뒤의 문장 부호 떼기 | 읽기 코드 |

상수 `After…`, `Jong…`, `LeadVowel`, `Slot…`도 여기 한 벌만 있다.
