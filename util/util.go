package util

import (
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// GotWantTest takes *testing.T
func GotWantTest[T comparable](got, want T, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

// ParseMDToHTML
// Uses simplecss.org
func ParseMDToHTML(md []byte) []byte {
	head := []byte(`<link rel="stylesheet" href="https://cdn.simplecss.org/simple.min.css">`)

	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags, Head: head}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
