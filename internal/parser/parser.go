package parser

import (
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/humtta/rosalind-cli/internal/model"
)

func ParseProblemListPage(r io.Reader, base string) (*[]model.Problem, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("parse HTML: %w", err)
	}

	rows := doc.Find("table.problem-list tbody tr")
	problems := make([]model.Problem, 0, rows.Length())

	for i := range rows.Length() {
		problem, err := parseProblemListRow(i, rows.Eq(i), base)
		if err != nil {
			return nil, fmt.Errorf("parse row %d: %w", i, err)
		}
		problems = append(problems, *problem)
	}

	return &problems, nil
}

func parseProblemListRow(
	i int,
	row *goquery.Selection,
	base string,
) (*model.Problem, error) {
	cells := row.Find("td")
	if cells.Length() < 2 {
		return nil, fmt.Errorf("unexpected layout")
	}

	link := cells.Eq(1).Find("a[href]").First()

	href, ok := link.Attr("href")
	if !ok || href == "" {
		return nil, fmt.Errorf("missing link")
	}

	url, err := resolveURL(base, href)
	if err != nil {
		return nil, fmt.Errorf("resolve URL: %w", err)
	}

	id := strings.TrimSpace(cells.Eq(0).Text())
	if id == "" {
		return nil, fmt.Errorf("missing id")
	}

	title := strings.TrimSpace(link.Text())
	if title == "" {
		return nil, fmt.Errorf("missing title")
	}

	return &model.Problem{
		Index: i,
		ID:    id,
		Title: title,
		URL:   url,
	}, nil
}

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
