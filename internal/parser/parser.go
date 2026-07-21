package parser

import (
	"net/url"

	_ "github.com/PuerkitoBio/goquery"
)

func resolveURL(base, path string) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	p, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	return b.ResolveReference(p).String(), nil
}
