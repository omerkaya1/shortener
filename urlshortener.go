package shortener

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

var urlShortener URLShortener
var m sync.RWMutex

const ab = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type URLShortener interface {
	Shorten(string) string
	Resolve(string) string
}

type URLShortenerImpl struct {
	Store map[string]string
	count int64
}

func GetURLShortener() URLShortener {
	if urlShortener != nil {
		return urlShortener
	} else {
		urlShortener = &URLShortenerImpl{Store: map[string]string{}, count: 0}
	}
	return urlShortener
}

func (us *URLShortenerImpl) Shorten(longURL string) string {
	u, err := url.Parse(longURL)
	if err != nil {
		return ""
	}

	if len(u.Path) == 0 || len(u.Scheme) == 0 || len(u.Host) == 0 {
		return ""
	}
	m.Lock()
	defer m.Unlock()
	shortened := strings.Builder{}
	counter := us.count
	for {
		temp := counter % int64(len(ab))
		shortened.WriteByte(ab[temp])
		if counter = counter / int64(len(ab)); counter == 0 {
			break
		}
	}

	shortURL := fmt.Sprintf("%s://%s/%s", u.Scheme, u.Host, shortened.String())
	if _, exist := us.Store[shortURL]; !exist {
		us.Store[shortURL] = longURL
		us.count++
		return shortURL
	} else {
		return ""
	}
}

func (us *URLShortenerImpl) Resolve(shortURL string) string {
	m.RLock()
	defer m.RUnlock()
	return us.Store[shortURL]
}
