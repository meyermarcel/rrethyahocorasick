# ahocorasick

The fastest Golang implementation of the Aho-Corasick algorithm for string-searching.

## Usage

```bash
go get github.com/meyermarcel/ahocorasick@v1.0.0
```

[Documentation](https://godoc.org/github.com/meyermarcel/ahocorasick)

```go
matcher := CompileByteSlices([][]byte{
	[]byte("he"),
	[]byte("she"),
	[]byte("his"),
	[]byte("hers"),
	[]byte("she"),
})
fmt.Print(matcher.FindAllByteSlice([]byte("ushers")))

// Output:
// [{ "he" 2 } { "she" 1 } { "she" 1 } { "hers" 2 }]
```

```go
matcher := CompileStrings([]string{
	"he",
	"she",
	"his",
	"hers",
	"she",
})
fmt.Print(matcher.FindAllString("ushers"))
}

// Output:
// [{ "he" 2 } { "she" 1 } { "she" 1 } { "hers" 2 }]
```

## Benchmark

See [benchmark/result.txt](benchmark/result.txt) or execute

```
(cd benchmark/ && go test -bench . -benchmem | tee result.txt)
```

### Reference Papers

[1] A. V. Aho, M. J. Corasick, "Efficient String Matching: An Aid to Bibliographic Search," Communications of the ACM, vol. 18, no. 6, pp. 333-340, June 1975.

[2] J.I. Aoe, "An Efficient Digital Search Algorithm by Using a Doble-Array Structure," IEEE Transactions on Software Engineering, vol. 15, no. 9, pp. 1066-1077, September 1989.

[3] J.I. Aoe, K. Morimoto, T. Sato, "An Efficient Implementation of Trie Stuctures," Software - Practice and Experience, vol. 22, no.9, pp. 695-721, September 1992.

## License

`MIT`
