package shorturl

import (
	"context"
	"fmt"

	"github.com/rohanraj7316/rsrc-bp-testing/repository"
)

type Model interface {
	Top3ShortedDomain(ctx context.Context) (map[string]int, error)
	Create(context.Context, string) (string, error)
	Get(context.Context, string) (string, error)
}

type model struct {
	redirectHost string
	shortUrls    repository.ShortUrls
}

func NewModel(
	shortUrls repository.ShortUrls,
	redirectHost string,
) Model {
	return &model{
		redirectHost: redirectHost,
		shortUrls:    shortUrls,
	}
}

func (m *model) Top3ShortedDomain(ctx context.Context) (map[string]int, error) {
	return m.shortUrls.Top3ShortedDomain()
}

func (m *model) Create(ctx context.Context, originalUrl string) (string, error) {
	shortId, err := m.shortUrls.Create(originalUrl)
	if err != nil {
		return "", err
	}

	redirectionUrl := fmt.Sprintf("%s/%s", m.redirectHost, shortId)

	return redirectionUrl, nil
}

func (m *model) Get(ctx context.Context, shortId string) (string, error) {
	originalUrl, err := m.shortUrls.Get(shortId)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
