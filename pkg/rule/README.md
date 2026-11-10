# rule

toulmin 그래프의 규칙 함수(`func(Context, Specs) (float64, any)`) 가운데 기준(Spec)만 바꿔 여러 노드와 여러 그래프가 함께 쓰는 것을 모았다. 함수가 읽는 사실 타입과 컨텍스트 열쇠도 여기 있다. `internal/`에 기대지 않는다.

## 사실과 열쇠

- `KeyWord` → `*Word`(어절: 글, 조사 사슬 `Chain`, 꼬리 풀이들 `[]Tail`, 가장 긴 꼬리, 마지막 어절인지, 의존명사로 시작하는지, 문장 `*Sentence`).
- `KeyTail` → `Tail`(판정 중인 꼬리 풀이 하나).
- 읽는 쪽이 이 타입을 채워 컨텍스트에 넣는다. 규칙은 `WordOf`, `TailOf`로 꺼낸다.

## 기준

| 기준 | 쓰임 |
|---|---|
| `Claim{Name, Concludes}` | 주장 노드의 이름과 결론. 같은 함수를 여러 주장이 쓸 수 있게 하고, 평가 결과의 이름에서 `ClaimOf`로 되찾는다 |
| `Bits{Of, Mask}` | 격(`Case…`), 맺음 갈래(`Kind…`), 어절의 사실(`WordLast`, `WordBound`) 가운데 어느 것을 볼지 |
| `Versus{Mine, Rival}` | `Longer`가 견줄 두 길이(`SideChain`, `SideTail`, `SideLongest`) |
| toulmin `StringListSpec` | 낱말·조사·맺음의 목록 |

## 규칙 함수

값은 사실 하나를 보는 함수면 +1(적용)·-1(적용 안 됨)이고, 관측 횟수를 견주는 함수면 두 횟수 가운데 상대의 몫을 [-1, 1]에 옮긴 값이다.

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `ChainCase` | rule | 사슬의 가장 바깥 격조사가 기준의 격이다 | `Bits`(case) |
| `ChainNoCase` | rule | 격조사 없이 보조사만 붙었다 | 없음 |
| `ChainListed` | rule | 격조사 없는 사슬에 기준 목록의 조사가 있다 | 조사 목록 |
| `WordListed` | rule, except | 어절이 통째로 기준 목록의 낱말이다 | 낱말 목록 |
| `ClosingKind` | rule | 꼬리의 맺음이 기준의 갈래다 | `Bits`(kind) |
| `ClosingListed` | rule | 꼬리의 맺음이 기준 목록에 있다 | 맺음 목록 |
| `WordIs` | except | 어절이 기준의 사실을 갖는다(마지막 어절, 의존명사로 시작) | `Bits`(word) |
| `Longer` | except | 경쟁 풀이가 이 풀이보다 길게 읽는다 | `Versus` |
| `ChainSeenMore` | counter | 꼬리만큼 긴 조사가 관측됐다(횟수의 몫) | 없음 |
| `TailSeenMore` | counter | 사슬만큼 긴 꼬리가 관측됐다(횟수의 몫) | 없음 |

토큰 열 하나를 표지로 판정하는 그래프(컨텍스트 열쇠 `KeyPiece` → `Piece`)가 쓰는 함수:

| 함수 | 쓰임 | 뜻 | 기준 |
|---|---|---|---|
| `PieceRead` | rule, counter | 색인에서 걸린 토큰 열이 이 자리에 있다 | 없음 |
| `PieceNoStem` | except | 표지 앞에 줄기가 없다 | 없음 |
| `PieceAfter` | except | 앞 토큰의 끝이 붙는 조건과 안 맞는다 | `Bits`(after) |
| `ChainOrder` | except | 조사 사슬의 순서가 어긋난다 | 없음 |
| `SlotOrder` | except | 선어말 칸의 순서가 어긋난다 | `Bits`(repeat) |
| `PieceBase` | except | 합쳐진 조각이 기준 목록에 없는 음절 밑에 있다 | 음절 목록 |

도움 함수: `PieceOf`, `Applies`(예/아니오 → ±1), `Share`(횟수의 몫), `WordOf`, `TailOf`, `ClaimOf`.
