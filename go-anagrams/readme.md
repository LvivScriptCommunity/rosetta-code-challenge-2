# How to run

1. Install golang: https://github.com/moovweb/gvm
2. Run cmd `go run main.go  anagram.go` or `go run $(ls -1 *.go | grep -v _test.go)`  

3. Result should be similar to:
```shell
[angel angle galen glean lange]
[abel able bale bela elba]
[elan lane lean lena neal]
[caret carte cater crate trace]
[evil levi live veil vile]
[alger glare lager large regal]
```

4. Run tests: `go test ./... -v`
5. Run benchmarks: `go test -v -bench=. -benchmem`
```shell
=== RUN   TestSortAnagram_lviv_vivl
--- PASS: TestSortAnagram_lviv_vivl (0.00s)
=== RUN   TestSortAnagram_lviV_vvvl
--- PASS: TestSortAnagram_lviV_vvvl (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/LvivScriptCommunity/rosetta-code-challenge-2/go-anagrams
BenchmarkAnagram-8                  4120            280398 ns/op               0 B/op          0 allocs/op
BenchmarkAnagramSort-8              2134            529826 ns/op          127744 B/op       3992 allocs/op
PASS
ok      github.com/LvivScriptCommunity/rosetta-code-challenge-2/go-anagrams     2.386s

```
