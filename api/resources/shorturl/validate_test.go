package shorturl_test

import (
	"testing"

	"github.com/rohanraj7316/rsrc-bp-testing/api/resources/shorturl"
	"github.com/stretchr/testify/assert"
)

func TestValidateShortUrl(t *testing.T) {
	t.Parallel()

	t.Run("invalid short id", func(t *testing.T) {
		shortId := ""

		err := shorturl.ValidateShortUrl(shortId)

		assert.Error(t, err)
	})

	t.Run("valid short id", func(t *testing.T) {
		shortId := "aw123"

		err := shorturl.ValidateShortUrl(shortId)

		assert.NoError(t, err)
	})
}

func TestValidateOriginalUrl(t *testing.T) {
	t.Parallel()

	t.Run("valid url", func(t *testing.T) {
		originalUrl := "https://exampl.com/testing"

		err := shorturl.ValidateOriginalUrl(originalUrl)

		assert.NoError(t, err)
	})

	t.Run("invalid url", func(t *testing.T) {
		originalUrl := "exampl/testing"

		err := shorturl.ValidateOriginalUrl(originalUrl)

		assert.Error(t, err)
	})
}
