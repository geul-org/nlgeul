# koparse

한국어 문장을 문법 표만으로 분석하는 실험 도구다. Go 모듈 `github.com/geul-org/koparse`과 cobra CLI `kopas`이고, 하위 명령은 `mini` 하나다(cobra의 `help`는 남고 `completion`은 껐다). 2026-11-10 이전의 분석기(명사 분리·식별)는 git 기록에 남아 있다.

## kopas mini

```sh
go install ./cmd/kopas
kopas mini "이것은 문장을 분해한다"
```

문장 하나(인자 하나)를 띄어쓰기로 나눠 어절마다 `어절:문장성분`을 한 줄씩 찍는다. 이름은 `주어`, `목적어`, `보어`, `서술어`, `관형어`, `부사어`, `독립어`, `미정`이고, 풀이가 여럿이면 `|`로 잇는다. 단어와 조사, 용언의 줄기와 꼬리 조각(선어말 칸, 맺음, 요, 인용의 고, 명사 뒤 하다·되다·시키다) 사이에는 `+`를 넣는다. 줄기와 조각이 한 음절에 합쳐진 곳은 표면 꼴 그대로 둔다(`지시+해서`, `합+니다`).

```
이것+은:관형어|주어|목적어
문장+을:목적어|관형어
분해+한+다:서술어
```

사전과 낱말 목록 없이 조사, 어미 사슬, 문법 형태소의 표만으로 어디까지 되는지 재는 실험이다. 코드 주석의 `particles.md`, `endings.md`, `morphemes.md`는 그 표를 적은 작업 문서의 이름이고 이 저장소에는 없다. 표는 코드 안에 옮겨 두었고 실행 때 md를 읽지 않는다.

입력 문장의 전제조건(어긴 입력을 고쳐 읽지 않는다):

1. 동사 서술은 반드시 마지막에 온다.
2. 띄어쓰기를 올바로 한다.
3. 조사를 반드시 올바로 붙인다.

## 구성

- 어절의 오른쪽 끝을 음절·자모로 읽는 부분은 일반 Go 코드, 읽은 사실로 문장성분을 가리는 판정은 toulmin 그래프다([internal/mini/](internal/mini/README.md)).
- toulmin은 park-jun-woo/toulmin `geul` 브랜치의 Go 모듈(MIT)을 공개 커밋 버전으로 require한다.
- 제3자 구성요소: `github.com/park-jun-woo/toulmin/go`(MIT), `github.com/spf13/cobra` v1.10.2(Apache-2.0), `github.com/spf13/pflag` v1.0.9(BSD-3-Clause), `github.com/inconshreveable/mousetrap` v1.1.0(Apache-2.0). 우리 코드는 CC0다(LICENSE, NOTICE).
- 규칙: filefunc(`codebook.yaml`, 함수·타입별 파일), TSMA, `go test`·`go vet`·`go build`. `go.work`는 쓰지 않는다(`GOWORK=off`).
- 패키지는 [internal/](internal/README.md), 진입점은 [cmd/](cmd/README.md).
