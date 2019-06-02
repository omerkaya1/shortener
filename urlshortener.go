package shortener

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

var urlShortener URLShortener

type URLShortener interface {
	Shorten(string) string
	Resolve(string) string
}

type URLShortenerImpl map[string]string

func GetURLShortener() URLShortener {
	if urlShortener != nil {
		return urlShortener
	} else {
		urlShortener = make(URLShortenerImpl)
	}
	return urlShortener
}

func (us URLShortenerImpl) Shorten(longURL string) string {
	u, err := url.Parse(longURL)
	if err != nil {
		return ""
	}

	if len(u.Path) == 0 || len(u.Scheme) == 0 || len(u.Host) == 0 {
		return ""
	}

	if _, exist := us[longURL]; !exist {
		shortenedPath := generateShortURL(longURL)
		shortURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, shortenedPath)
		us[longURL] = shortURL
		return shortURL
	} else {
		return ""
	}
}

func (us URLShortenerImpl) Resolve(shortURL string) string {
	for k, v := range us {
		if v == shortURL {
			return k
		}
	}
	return fmt.Sprintf("%s URL cannot be resolved: unknown URL", shortURL)
}

func generateShortURL(long string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(long)))
}
