package handler

import (
	"net/http"
	"regexp"
	"strings"
	"golang.org/x/net/html"
)

// Extracts the first URL from a string (simple regex)
func extractFirstURL(text string) string {
	re := regexp.MustCompile(`https?://[^\s]+`)
	match := re.FindString(text)
	return match
}

// Fetches Open Graph/meta tags from a URL
func fetchLinkPreview(url string) (title, description, image string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return // End of document
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "meta" {
				var property, content string
				for _, a := range t.Attr {
					if a.Key == "property" || a.Key == "name" {
						property = a.Val
					}
					if a.Key == "content" {
						content = a.Val
					}
				}
				if property == "og:title" && title == "" {
					title = content
				}
				if property == "og:description" && description == "" {
					description = content
				}
				if property == "og:image" && image == "" {
					image = content
				}
				if property == "description" && description == "" {
					description = content
				}
			}
			if t.Data == "title" && title == "" {
				z.Next()
				title = strings.TrimSpace(z.Token().Data)
			}
		}
	}
}
