# go-fix-modernize-sample

Go 1.26 で導入された `go fix` modernize の動作を検証するためのサンプルプロジェクトです。

20 個の modernize ルールそれぞれに対応するコードを **修正前の状態** で収録しています。`go fix` を実行することで、各ファイルがモダンな Go のイディオムに自動変換されることを確認できます。

## 前提条件

- Go 1.26 以上（`.mise.toml` で管理）

## 使い方

```bash
# modernize で検出される箇所を確認（dry-run）
go fix -diff ./...

# 実際に修正を適用
go fix ./...

# 差分を確認
git diff
```

## 収録ルール一覧（20 rules）

| # | ルール名 | ファイル | 修正前 | 修正後 | 導入バージョン |
|---|---------|---------|--------|--------|-------------|
| 1 | **any** | `rule_any.go` | `interface{}` | `any` | Go 1.18 |
| 2 | **fmtappendf** | `rule_fmtappendf.go` | `[]byte(fmt.Sprintf(...))` | `fmt.Appendf(nil, ...)` | Go 1.19 |
| 3 | **forvar** | `rule_forvar.go` | `i := i`（ループ変数シャドウイング） | 削除（Go 1.22 でイテレーション毎スコープ化） | Go 1.22 |
| 4 | **mapsloop** | `rule_mapsloop.go` | `for k, v := range src { dst[k] = v }` | `maps.Copy(dst, src)` | Go 1.21 |
| 5 | **minmax** | `rule_minmax.go` | `if b < x { x = b }` パターン | `min(a, b)` / `max(a, b)` | Go 1.21 |
| 6 | **newexpr** | `rule_newexpr.go` | `func intPtr(v int) *int { return &v }` | `new(T)` でインライン化 | Go 1.20 |
| 7 | **omitzero** | `rule_omitzero.go` | `json:",omitempty"`（struct/time フィールド） | `json:",omitzero"` | Go 1.24 |
| 8 | **plusbuild** | `rule_plusbuild.go` | `// +build linux` | `//go:build linux` のみ（旧構文削除） | Go 1.16 |
| 9 | **rangeint** | `rule_rangeint.go` | `for i := 0; i < n; i++` | `for i := range n` | Go 1.22 |
| 10 | **reflecttypefor** | `rule_reflecttypefor.go` | `reflect.TypeOf((*T)(nil)).Elem()` | `reflect.TypeFor[T]()` | Go 1.20 |
| 11 | **slicescontains** | `rule_slicescontains.go` | ループで要素探索 | `slices.Contains()` | Go 1.21 |
| 12 | **slicessort** | `rule_slicessort.go` | `sort.Ints` / `sort.Strings` / `sort.Slice` | `slices.Sort` / `slices.SortFunc` | Go 1.21 |
| 13 | **stditerators** | `rule_stditerators.go` | `for i := 0; i < t.NumMethod(); i++` | `for m := range t.Methods()` | Go 1.22 |
| 14 | **stringsbuilder** | `rule_stringsbuilder.go` | `result += p`（ループ内文字列結合） | `strings.Builder` / `strings.Join` | Go 1.10 |
| 15 | **stringscut** | `rule_stringscut.go` | `strings.Index` + スライス | `strings.Cut()` | Go 1.18 |
| 16 | **stringscutprefix** | `rule_stringscutprefix.go` | `HasPrefix` + `TrimPrefix` | `strings.CutPrefix()` / `CutSuffix()` | Go 1.20 |
| 17 | **stringsseq** | `rule_stringsseq.go` | `strings.Split` でイテレーション | `strings.SplitSeq`（イテレータ） | Go 1.22 |
| 18 | **testingcontext** | `rule_testingcontext_test.go` | `context.WithCancel(context.Background())` + `defer cancel()` | `t.Context()` | Go 1.20 |
| 19 | **unsafefuncs** ⚠️ | `rule_unsafefuncs.go` | `unsafe.Pointer(uintptr(ptr) + uintptr(n))` | `unsafe.Add(ptr, n)` | Go 1.17 |
| 20 | **waitgroup** | `rule_waitgroup.go` | `wg.Add(1)` + `go func() { defer wg.Done() }` | `wg.Go(func() { ... })` | Go 1.22 |

> ⚠️ **unsafefuncs** は `golang.org/x/tools` の modernize パッケージに定義されていますが、Go 1.26 の `go tool fix` にバンドルされた `modernize.Suite` には含まれていません。`go fix ./...` では適用されません。

## 各ルールの詳細

### 1. any

`interface{}` を組み込み型エイリアス `any` に置換します。

```go
// Before
func ExampleAny(v interface{}) interface{} { ... }

// After
func ExampleAny(v any) any { ... }
```

### 2. fmtappendf

`[]byte(fmt.Sprintf(...))` の不要な中間文字列アロケーションを `fmt.Appendf` で排除します。

```go
// Before
return []byte(fmt.Sprintf("name: %s, age: %d", name, age))

// After
return fmt.Appendf(nil, "name: %s, age: %d", name, age)
```

### 3. forvar

Go 1.22 のイテレーション毎スコープにより不要になったループ変数のシャドウイング `i := i` を削除します。

```go
// Before
for i := range 5 {
    i := i // unnecessary since Go 1.22
    result = append(result, &i)
}

// After
for i := range 5 {
    result = append(result, &i)
}
```

### 4. mapsloop

マップのコピーループを `maps.Copy` に置換します。

```go
// Before
for k, v := range src { dst[k] = v }

// After
maps.Copy(dst, src)
```

### 5. minmax

if 文による min/max パターンを組み込み関数に置換します。

```go
// Before
x := a
if b < x { x = b }

// After
x := min(a, b)
```

### 6. newexpr

ポインタ取得ヘルパー関数を `new(T)` でインライン化します。

```go
// Before
func intPtr(v int) *int { return &v }
a := intPtr(42)

// After (inlined)
a := new(42)
```

### 7. omitzero

struct や `time.Time` フィールドの `omitempty` を、より意味的に正確な `omitzero` に置換します。

```go
// Before
CreatedAt time.Time `json:"created_at,omitempty"`

// After
CreatedAt time.Time `json:"created_at,omitzero"`
```

### 8. plusbuild

`//go:build` がある場合に旧形式の `// +build` 行を削除します。

```go
// Before
//go:build linux
// +build linux

// After
//go:build linux
```

### 9. rangeint

C スタイルの for ループを `range` に置換します。

```go
// Before
for i := 0; i < n; i++ { sum += i }

// After
for i := range n { sum += i }
```

### 10. reflecttypefor

冗長な `reflect.TypeOf` パターンをジェネリクスベースの `reflect.TypeFor` に置換します。

```go
// Before
reflect.TypeOf((*error)(nil)).Elem()

// After
reflect.TypeFor[error]()
```

### 11. slicescontains

手動の要素探索ループを `slices.Contains` に置換します。

```go
// Before
for _, item := range items {
    if item == target { return true }
}
return false

// After
return slices.Contains(items, target)
```

### 12. slicessort

レガシーな `sort` パッケージ関数を型安全な `slices` パッケージに置換します。

```go
// Before
sort.Ints(s)

// After
slices.Sort(s)
```

### 13. stditerators

インデックスベースのカウンターループを標準ライブラリのイテレータに置換します。

```go
// Before
for i := 0; i < t.NumMethod(); i++ {
    m := t.Method(i)
    fmt.Println(m.Name)
}

// After
for m := range t.Methods() {
    fmt.Println(m.Name)
}
```

### 14. stringsbuilder

ループ内の非効率な文字列結合を `strings.Builder` や `strings.Join` に置換します。

```go
// Before
var result string
for _, p := range parts { result += p }

// After
var b strings.Builder
for _, p := range parts { b.WriteString(p) }
return b.String()
```

### 15. stringscut

`strings.Index` + スライスパターンを `strings.Cut` に置換します。

```go
// Before
i := strings.Index(s, "=")
if i < 0 { return s, "", false }
return s[:i], s[i+1:], true

// After
before, after, found := strings.Cut(s, "=")
```

### 16. stringscutprefix

`HasPrefix` + `TrimPrefix`（および `HasSuffix` + `TrimSuffix`）を 1 回の呼び出しに統合します。

```go
// Before
if strings.HasPrefix(s, "prefix_") {
    return strings.TrimPrefix(s, "prefix_"), true
}

// After
if after, ok := strings.CutPrefix(s, "prefix_"); ok {
    return after, true
}
```

### 17. stringsseq

`strings.Split` / `strings.Fields` のイテレーションをイテレータ版に置換し、不要なスライスアロケーションを回避します。

```go
// Before
for _, part := range strings.Split(s, ",") { ... }

// After
for part := range strings.SplitSeq(s, ",") { ... }
```

### 18. testingcontext

テスト内の `context.WithCancel(context.Background())` + `defer cancel()` を `t.Context()` に置換します。

```go
// Before
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// After
ctx := t.Context()
```

### 19. unsafefuncs

> **注意:** このルールは Go 1.26 の `go tool fix` にバンドルされた `modernize.Suite` に含まれていないため、`go fix ./...` では適用されません。
> `golang.org/x/tools` の最新版を直接実行すれば適用可能です。
>
> ```bash
> go run golang.org/x/tools/go/analysis/passes/modernize/cmd/modernize@latest -fix ./...
> ```

冗長な unsafe ポインタ演算を `unsafe.Add` に置換します。

```go
// Before
unsafe.Pointer(uintptr(ptr) + uintptr(n))

// After
unsafe.Add(ptr, n)
```

### 20. waitgroup

WaitGroup の定型パターンを `wg.Go` に置換します。

```go
// Before
wg.Add(1)
go func() {
    defer wg.Done()
    // do work
}()

// After
wg.Go(func() {
    // do work
})
```

## 参考リンク

- [Go 1.26 Release Notes](https://go.dev/doc/go1.26)
- [Proposal: cmd/go: fix: apply fixes from modernizers, inline, and other analyzers (#71859)](https://github.com/golang/go/issues/71859)
- [modernize package - golang.org/x/tools/go/analysis/passes/modernize](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/modernize)

## License

MIT
