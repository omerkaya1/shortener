package shortener

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	longURLOne    = `https://www.google.com/search?q=wikipedia&oq=wiki&aqs=chrome.0.0j69i60j69i57j69i60j69i65l2.2896j0j7&sourceid=chrome&ie=UTF-8`
	longURLTwo    = `https://www.google.com/search?ei=Ax_0XN2bA6-trgSJhKeoCw&q=wikipedia+website&oq=wikipedia&gs_l=psy-ab.1.0.0i71l8.0.0..33818...0.0..0.0.0.......0......gws-wiz.5DOrqI04Ohw`
	longURLThree  = `https://www.google.com/search?ei=Jx_0XLWsAvDkrgTL1JL4DQ&q=wikipedia+english&oq=wikipedia+website&gs_l=psy-ab.1.0.0i71l8.0.0..1072572...0.0..0.0.0.......0......gws-wiz.lpsA8VirRYU`
	unknownShort  = "djfjhjhwjkfhjj23401984nb34jkh"
	malformedURL  = "htttrrexncliclsl"
	urlWithNoPath = "https://google.com"
)

func TestGetURLShortener(t *testing.T) {
	us := GetURLShortener()
	assert.NotNil(t, us)
	assert.Implements(t, (*URLShortener)(nil), us)
}

func TestURLShortenerImpl_Shorten(t *testing.T) {
	us := GetURLShortener()
	link1 := us.Shorten(longURLOne)
	link2 := us.Shorten(longURLTwo)
	assert.NotEmptyf(t, link1, "%s\n", link1)
	assert.NotEmptyf(t, link2, "%s\n", link2)
	assert.NotEqual(t, link1, link2)
	assert.Equal(t, "", us.Shorten(malformedURL))
	assert.Equal(t, "", us.Shorten(urlWithNoPath))
}

func TestURLShortenerImpl_Resolve(t *testing.T) {
	us := GetURLShortener()
	link := us.Shorten(longURLThree)
	assert.Equal(t, longURLThree, us.Resolve(link))
	assert.Equal(t, fmt.Sprintf("%s URL cannot be resolved: unknown URL", unknownShort), us.Resolve(unknownShort))
}
