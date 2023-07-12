# WanaKana
[![godoc](https://godoc.org/github.com/deelawn/wanakana?status.svg)](https://godoc.org/github.com/deelawn/wanakana)
[![Coverage Status](https://coveralls.io/repos/github/deelawn/wanakana/badge.svg)](https://coveralls.io/github/deelawn/wanakana)
[![Go Report Card](https://goreportcard.com/badge/github.com/deelawn/wanakana)](https://goreportcard.com/report/github.com/deelawn/wanakana)

This repo is a port of Wanikani's JS implementation located [here](https://github.com/WaniKani/WanaKana).

No errors are returned by any of the exported functions or methods; this is to keep the behavior the same as the original implementation, but there is an argument to be made to change this.

## Improvements
If anyone would like to make a PR to improve this repository, here are a few suggestions regarding what could be done:
- [ ] Improve tree caching -- currently only one tree can be stored in the cache and it is not modifiable
- [ ] Add a wasm compilation target and add support for binding and unbinding an element similar to what the original implementation is capable of
- [ ] Move the tree generation logic into a `generate` package. Running `go generate` should generate source code with the `go:embed` directive along with all of the key/value pairs to store in the tree. This makes it more transparent as to what each tree contains by default rather than having to decipher how it is being generated.

## Usage
Here is an example of how this package can be used:
``` go
package main

import (
	"fmt"

	"github.com/deelawn/wanakana"
	"github.com/deelawn/wanakana/config"
)

func main() {

	romaji := "okonomiyakinikuman"
	options := config.Options{
		IMEMode: config.ToKanaMethodKatakana,
	}

	kana := wanakana.ToKana(romaji, options, nil)

	fmt.Println(kana)
}

```

This produces `オコノミヤキニクマン`
