package llmstxtgen

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mackee/go-readability"
	"github.com/snabb/sitemap"
)

func parseContent(body []byte) (*readability.ReadabilityArticle, error) {
	options := readability.DefaultOptions()
	article, err := readability.Extract(string(body), options)
	if err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}
	return &article, nil
}

func Full(sitemapReader io.Reader) (io.Reader, error) {
	index := sitemap.NewSitemapIndex()
	_, err := index.ReadFrom(sitemapReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read sitemap index: %w", err)
	}
	sitemaps := make([]*sitemap.Sitemap, len(index.URLs))
	for i, u := range index.URLs {
		m := sitemap.New()
		r, err := http.Get(u.Loc)
		if err != nil {
			return nil, fmt.Errorf("failed to get sitemap: %w", err)
		}
		defer r.Body.Close()
		_, err = m.ReadFrom(r.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read sitemap: %w", err)
		}
		sitemaps[i] = m
	}
	reader, writer := io.Pipe()
	go func() {
		for _, s := range sitemaps {
			for _, u := range s.URLs {
				r, err := http.Get(u.Loc)
				if err != nil {
					panic(fmt.Errorf("failed to get page: %w", err))
				}
				defer r.Body.Close()
				b, err := io.ReadAll(r.Body)
				if err != nil {
					panic(fmt.Errorf("failed to read page: %w", err))
				}
				article, err := parseContent(b)
				if err != nil {
					panic(fmt.Errorf("failed to parse content: %w", err))
				}
				fmt.Fprintln(writer, readability.ToMarkdown(article.Root))
				fmt.Fprintln(writer, "---")
			}
		}
		writer.Close()
	}()
	return reader, nil
}
