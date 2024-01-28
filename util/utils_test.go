// utils_test.go
package util

import (
	"strings"
	"testing"
)

func TestParseMDToHTML(t *testing.T) {
	got := string(ParseMDToHTML([]byte(md_txt)))

	// Map to see if the HTML contains everything expected from the MD
	contains := map[string]bool{
		"css":      false,
		"bold":     false,
		"li_1":     false,
		"nested_2": false,
		"quote":    false,
		"link":     false,
	}

	split_html := strings.Split(got, "\n")

	// Loop thru HTML, confirm stuff exists
	for line_num, line := range split_html {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "cdn.simplecss.org") {
			contains["css"] = true
		}
		if strings.Contains(line, "<p><strong>bold</strong></p>") {
			contains["bold"] = true
		}
		if strings.Contains(line, "<li>li_1</li>") {
			// TODO: make sure this is in a ul?
			contains["li_1"] = true
		}
		if strings.Contains(line, "<li>nested_2</li>") {
			// make sure this is nested
			if strings.Contains(split_html[line_num+1], "</ul></li>") {
				contains["nested_2"] = true
			}
		}
		if strings.Contains(line, "<p>quote</p>") {
			// make sure this is nested
			if strings.Contains(split_html[line_num+1], "</blockquote>") {
				contains["quote"] = true
			}
		}
		if strings.Contains(line, "href") {
			// TODO: confirm link
			contains["link"] = true
		}

	}

	// Check if anything in contains is false, complain about it
	for k, v := range contains {
		if !v {
			t.Fatalf("html missing %s output.\n\nmd: %s\n\n. html: %s\n\n", k, md_txt, got)
		}
	}

	// t.Log(got)
}

var md_txt = `# Title
**bold**

- li_1
- li_2
	- nested_1
	- nested_2

> quote

[link_text](http://example.com/link_url)
`
