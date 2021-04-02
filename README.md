# Singlish
[![Test](https://github.com/wgarunap/singlish/actions/workflows/go.yml/badge.svg)](https://github.com/wgarunap/singlish/actions/workflows/go.yml/badge.svg)
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/wgarunap/singlish/blob/main/LICENSE)
[![GitHub Release](https://img.shields.io/github/release/wgarunap/singlish.svg?style=flat)](https://github.com/wgarunap/singlish/releases)
### Overview
Sinhala to Singlish converter in Golang. accuracy is not 100% guaranteed. Tested for main words and strings only.

### Use Repository
```shell
go get -u github.com/wgarunap/singlish
```

### Example
```go
package main

import (
	"fmt"
	singlish "github.com/wgarunap/singlish"
)

func main() {
	txt:=singlish.TranslateString("මෙය url generate කිරීම උදෙසා ලියන ලද්දක් බැවින් නිරවද්‍යතාවය 100% තහවුරු කර නැත.")
	fmt.Println(txt) 
	//Output
	//meya url generate kiriima udhesa liyana ladhdhak bewin nirawadhyathawaya 100 thahawuru kara netha
	
	txt2:=singlish.TranslateSinhalaStr("මෙය සිංහල එකකි")
	fmt.Println(txt2)
	//Output
	//meya sinhala ekaki

	txt3:=singlish.TranslateString(`මෙය !"#$%&'()*+,-./:;<=>?@[\]^_{|}~ එකකි `)
	fmt.Println(txt3)
	// Output 
	//meya ekaki
}

```

### Run Test
```shell
go test ./...
```

### Benchmark
Run benchmarks
```shell
go test -bench=.
```
Benchmark Results
```shell
goos: darwin
goarch: amd64
pkg: github.com/wgarunap/singlish
cpu: Intel(R) Core(TM) i7-5557U CPU @ 3.10GHz
BenchmarkTranslate1Char-4        4068814               283.9 ns/op
BenchmarkTranslate4Char-4        1000000              1018 ns/op

```

### Contributions

I am more than happy to receive contributions to this repository.

contact : wg.aruna.p@gmail.com
