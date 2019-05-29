package url_shortener

import (
	"fmt"
	"math/rand"
	"net/url"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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
		fmt.Printf("failure: %v\n", err)
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
	shortenedPath := generateShortURL()
	if _, exist := us.alreadyExist(rawUrl); !exist {
		shortURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, shortenedPath)
		us[rawUrl] = shortURL
		return shortURL, nil
	} else {
		return "", fmt.Errorf("URL %s was already shortened\n", rawUrl)
	}
}

func (us URLShortenerImpl) alreadyExist(url string) (string, bool) {
	if retVal, ok := us[url]; ok {
		return retVal, ok
	} else {
		return "", false
	}
}

func generateShortURL() string {
	short := make([]byte, 10)
	for i := range short {
		short[i] = letters[rand.Intn(len(letters))]
	}
	return string(short)
}
