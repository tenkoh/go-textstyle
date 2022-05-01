package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tenkoh/go-textstyle"
	"golang.org/x/text/transform"
)

func main() {
	s := "abcdefghijklmnopqrstuvwxyz0123456789\n\n"
	trs := []*textstyle.Transformer{
		textstyle.Bold(),
		textstyle.Italic(),
		textstyle.BoldItalic(),
		textstyle.Script(),
		textstyle.BoldScript(),
		textstyle.Fraktur(),
		textstyle.BoldFraktur(),
		textstyle.DoubleStruck(),
		textstyle.SansSerif(),
		textstyle.SansSerifBold(),
		textstyle.SansSerifItalic(),
		textstyle.SansSerifBoldItalic(),
		textstyle.Monospace(),
	}
	for _, tr := range trs {
		sr := tr.Rep.(*textstyle.SimpleReplacer)
		fmt.Println(sr.Name, ": ")
		r := transform.NewReader(strings.NewReader(s), tr)
		io.Copy(os.Stdout, r)
	}
}
