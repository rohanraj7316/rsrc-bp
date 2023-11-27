package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Parallel()

	t.Run("url exists", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{
				"abc123": "https://example.com",
			},
		}

		expectedURL := "https://example.com"
		resultURL, err := urls.Get("abc123")

		assert.NoError(t, err)
		assert.Equal(t, expectedURL, resultURL)
	})

	t.Run("url don't exists", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{
				"abc123": "https://example.com",
			},
		}

		resultURL, err := urls.Get("xyz456")

		assert.Error(t, err)
		assert.Equal(t, "", resultURL)
		assert.EqualError(t, err, "not found")
	})

	t.Run("empty struct", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{},
		}

		resultURL, err := urls.Get("abc123")

		assert.Error(t, err)
		assert.Equal(t, "", resultURL)
		assert.EqualError(t, err, "not found")
	})
}

func TestCreate(t *testing.T) {
	t.Parallel()

	t.Run("create new shortId", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{},
		}

		originalURL := "https://example.com"

		shortID, err := urls.Create(originalURL)

		assert.NoError(t, err)
		assert.Equal(t, originalURL, urls.shortIdToOriginalUrls[shortID])
	})

	t.Run("use existing short123", func(t *testing.T) {
		existingURL := "https://existing.com"
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{
				"short123": existingURL,
			},
		}

		shortID, err := urls.Create(existingURL)

		assert.NoError(t, err)
		assert.Equal(t, "short123", shortID)
		assert.Equal(t, existingURL, urls.shortIdToOriginalUrls["short123"])
	})

	t.Run("empty map", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{},
		}

		shortID, err := urls.Create("")

		assert.Error(t, err)
		assert.Equal(t, "", shortID)
		assert.EqualError(t, err, "empty original url")
	})
}

func TestTop3ShortedDomain(t *testing.T) {
	t.Parallel()

	t.Run("top 3 shorted domain", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{
				"url1": "https://example.com/page1",
				"url2": "https://example.com/page2",
				"url3": "https://anotherdomain.com/page3",
				"url4": "https://anotherdomain.com/page4",
				"url5": "https://differentdomain.com/page5",
			},
		}

		expectedTop3ShortedDomain := map[string]int{
			"example.com":         2,
			"anotherdomain.com":   2,
			"differentdomain.com": 1,
		}

		haveTop3ShortedDomain, err := urls.Top3ShortedDomain()

		assert.NoError(t, err)
		assert.Equal(t, expectedTop3ShortedDomain, haveTop3ShortedDomain)
	})

	t.Run("empty shorted list", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{},
		}

		_, err := urls.Top3ShortedDomain()

		assert.Error(t, err, "domain count 0 which is less than 3")
	})

	t.Run("error since domain count is less than 3", func(t *testing.T) {
		urls := &shortUrls{
			shortIdToOriginalUrls: map[string]string{
				"url1": "https://example.com/page1",
				"url2": "https://example.com/page2",
			},
		}

		_, err := urls.Top3ShortedDomain()

		assert.Error(t, err, "domain count 1 which is less than 3")
	})
}
