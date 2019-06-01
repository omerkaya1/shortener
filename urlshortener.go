package url_shortener

import (
	"crypto/md5"
	"fmt"
	"net/url"
)

var urlShortener URLShortener

type URLShortener interface {
	Shorten(string) string
	Resolve(string) string
	addToKnown(string) (string, error)
	alreadyExist(string) (string, bool)
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
	retURL, err := us.addToKnown(longURL)
	if err != nil {
		fmt.Printf("failure: %v", err)
	}
	return fmt.Sprintf("%s", retURL)
}

func (us URLShortenerImpl) Resolve(shortURL string) string {
	for k, v := range us {
		if v == shortURL {
			return k
		}
	}
	return fmt.Sprintf("%s URL cannot be resolved: unknown URL", shortURL)
}

func (us URLShortenerImpl) addToKnown(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	if len(u.Path) == 0 {
		return "", fmt.Errorf("the provided URL has no path part to shorten")
	}

	if _, exist := us.alreadyExist(rawUrl); !exist {
		shortenedPath := generateShortURL(rawUrl)
		shortURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, shortenedPath)
		us[rawUrl] = shortURL
		return shortURL, nil
	} else {
		return "", fmt.Errorf("URL %s was already shortened", rawUrl)
	}
}

func (us URLShortenerImpl) alreadyExist(url string) (string, bool) {
	if retVal, ok := us[url]; ok {
		return retVal, ok
	} else {
		return "", false
	}
}

func generateShortURL(long string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(long)))
}
