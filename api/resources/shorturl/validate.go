package shorturl

import (
	"errors"
	"net/url"
)

func ValidateShortUrl(shortId string) error {
	if len(shortId) == 0 {
		return errors.New("empty short id")
	}

	return nil
}

func ValidateOriginalUrl(originalUrl string) error {
	if _, err := url.ParseRequestURI(originalUrl); err != nil {
		return err
	}

	return nil
}
