# go-textstyle
An implement of transform.Transformer to change text style. For example, Gopher to 𝔾𝕠𝕡𝕙𝕖𝕣, 𝕲𝖔𝖕𝖍𝖊𝖗 and so on.

## Note
This package is on proof of concept stage. There would be some breaking changes. Do not use this on your products.

## Usage
This package provides some implementations of transform.Transformer. Combining with transform.Reader enables you to accept any kind of io.Reader to transform text styles.

Here is a simple example.

```go
package main

import (
	"io"
	"os"
	"strings"

	"github.com/tenkoh/go-textstyle"
	"golang.org/x/text/transform"
)

func main() {
	s := "Hello, Gophers"
	r := transform.NewReader(strings.NewReader(s), textstyle.Bold())
	io.Copy(os.Stdout, r)
	//Output: 𝐇𝐞𝐥𝐥𝐨, 𝐆𝐨𝐩𝐡𝐞𝐫𝐬
}
```

You can use other io.Reader like os.Stdin, *os.File, http.Response.Body and so on.

## Available styles
- Bold
- Italic
- BoldItalic
- Script
- BoldScript
- Fraktur
- BoldFraktur
- DoubleStruck
- SansSerif
- SansSerifBold
- SansSerifItalic
- SansSerifBoldItalic
- Monospce

## Contribution
I'm very welcome your contribution. There are some assets to help adding text styles. If you are eager to add a complex transformation, do it your way without the assets.

### Code generation
If you want to add simple replacement, which means just giving codepoint offsets, you can use code generation tool. Please open `generate_style.go`, then add `Style` refering to existing styles. After editting, `go generate` command generates `styles.go` including added styles.

### Alternative character's map
Unicode table has some lacks (or reserved part) to alphabets. If you can not realize simple repalcement due to a few excepts, please consider to add a pairs to alternate them into `altMap` in `textstyle.go`.

## Roadmap
- 0.0.1: Basic functions and some styles (current)
- 0.0.x: Add Greek styles and enclosed styles.

## Thanks
I was very impressed by https://github.com/ikanago/omekasy , CLI tool to change text style implemented in Rust.

## License
MIT

## Author
tenkoh
