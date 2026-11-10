# internal

함수 하나에 파일 하나(filefunc)라 패키지 폴더의 파일 수는 20개를 넘는다.

| 패키지 | 내용 |
|---|---|
| [parse/](parse/README.md) | `nlgeul parse`: 어절 끝을 문법 표로 읽어 사실을 내고 toulmin 그래프로 문장성분을 판정 |
| [cli/](cli/README.md) | cobra 루트 명령 `nlgeul`과 하위 명령 `parse` |
| [token/](token/README.md) | `nlgeul token`: 토큰 스트림을 색인과 표지별 판정 그래프로 읽어 어절로 끊어 올림(실험) |
