package utils

import (
	"fmt"
	"net/url"
)

func ValidateURL(raw string) error {
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("invalid url scheme: %s", u.Scheme)
	}

	return nil
}
