package item

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//MetaData is the picture related metadata
type MetaData struct {
	twitterImage string
	ogImage      string
}

// GetMetaData gets the picture metadata
func GetMetaData(url string) string {
	res, _ := http.Get(url)

	data := new(MetaData)
	doc, err := goquery.NewDocumentFromReader((res.Body))
	if err != nil {
		panic(err)
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		property, _ := s.Attr("property")
		if property == "og:image" {
			data.ogImage, _ = s.Attr("content")
		}
		name, _ := s.Attr("name")
		if name == "twitter:image" {
			data.twitterImage, _ = s.Attr("content")
		}
		if name == "twitter:image:src" {
			data.twitterImage, _ = s.Attr("content")
		}
	})

	if data.ogImage != "" {
		return data.ogImage
	}
	return data.twitterImage
}
