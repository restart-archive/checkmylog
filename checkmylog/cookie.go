package checkmylog

import (
	"errors"
	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/fhttp/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// LoadCookies loads cookies from the specified content.
func LoadCookies(content string) (*cookiejar.Jar, bool) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, false
	}

	cookies := map[url.URL][]*http.Cookie{}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		c, err := resolveCookie(line)
		if err != nil {
			continue
		}
		prefix := "http://"
		if c.Secure {
			prefix = "https://"
		}
		uri, err := url.Parse(prefix + c.Domain)
		if err != nil {
			continue
		}

		cookies[*uri] = append(cookies[*uri], c)
	}

	for uri, cs := range cookies {
		jar.SetCookies(&uri, cs)
	}

	return jar, len(cookies) > 0
}

// resolveCookie resolves a cookie from the specified line.
func resolveCookie(line string) (*http.Cookie, error) {
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, "\"", "")

	c := strings.Split(line, "	")
	if len(c) != 7 {
		return nil, errors.New("invalid cookie")
	}

	expiry, _ := strconv.ParseFloat(c[4], 64)
	co := &http.Cookie{
		Name:     c[5],
		Value:    c[6],
		Path:     c[2],
		Domain:   c[0],
		Expires:  time.Unix(int64(expiry), 0),
		HttpOnly: strings.EqualFold(c[3], "TRUE"),
		Secure:   strings.EqualFold(c[1], "TRUE"),
	}

	return co, nil
}
