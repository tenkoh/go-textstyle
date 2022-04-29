package style_test

import (
	"io"
	"os"
	"strings"
	"testing"

	style "github.com/tenkoh/go-transform-style"
	"golang.org/x/text/transform"
)

func ExampleTransformer() {
	s := "Hello, Gophers"
	r := transform.NewReader(strings.NewReader(s), style.Bold)
	io.Copy(os.Stdout, r)
	// Output: 𝐇𝐞𝐥𝐥𝐨, 𝐆𝐨𝐩𝐡𝐞𝐫𝐬
}

func TestBold(t *testing.T) {
	src := "My new gear..."
	want := "𝐌𝐲 𝐧𝐞𝐰 𝐠𝐞𝐚𝐫..."

	reader := transform.NewReader(strings.NewReader(src), style.Bold)
	b, err := io.ReadAll(reader)
	if err != nil {
		t.Errorf("got error: %s\n", err)
		return
	}
	if got := string(b); got != want {
		t.Errorf("want %s, got %s\n", want, got)
	}
}
