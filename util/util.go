// util.go
package util

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// ValidateNoteContents
func ValidateNoteContents(contents string) error {
	if strings.Contains(contents, "baddata") {
		return fmt.Errorf("baddata in note contents")
	}
	return nil
}

// ParseMDToHTML
// Uses simplecss.org CSS
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

// TESTING STUFF //
// GotWantTest takes *testing.T
// Used for testing...
func GotWantTest[T comparable](got, want T, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
