package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

func download(url string) string {
	fmt.Println("Downloading ", url)

	response, err := http.Get(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("Link downloading error")
		}
	}(response.Body)

	if err != nil {
		return ""
	}

	utf8, err := charset.NewReader(response.Body, response.Header.Get("Content-Type"))
	if err != nil {
		fmt.Println("Encoding error:", err)
		return ""
	}

	if b, err := io.ReadAll(utf8); err == nil {
		return string(b)
	}

	return ""
}

func title(HTMLString string) (title string) {
	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)

	var i int
	for {
		tt := z.Next()

		i++
		if i > 100 { // Title should be one of the first tags
			return
		}

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <Title> tag
			if t.Data != "title" {
				continue
			}

			tt := z.Next()

			if tt == html.TextToken {
				t := z.Token()
				title = t.Data
				return
			}
		}
	}
}
