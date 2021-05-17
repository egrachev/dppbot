package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func download(url string) string {
	fmt.Println("Downloading ", url)

	response, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	if b, err := io.ReadAll(response.Body); err == nil {
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

			// Check if the token is an <title> tag
			if t.Data != "title" {
				continue
			}

			// fmt.Printf("%+v\n%v\n%v\n%v\n", t, t, t.Type.String(), t.Attr)
			tt := z.Next()

			if tt == html.TextToken {
				t := z.Token()
				title = t.Data
				return
				// fmt.Printf("%+v\n%v\n", t, t.Data)
			}
		}
	}
}

// func reverse(str string) string {
// 	r := []rune(str)
// 	var res []rune

// 	for i := len(r) - 1; i >= 0; i-- {
// 		res = append(res, r[i])
// 	}

// 	return string(res)
// }
