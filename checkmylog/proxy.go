package checkmylog

import (
	"bufio"
	"fmt"
	http "github.com/saucesteals/fhttp"
	"net/url"
	"os"
	"strings"
)

var proxies []proxyInfo

func init() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the path to the proxies (enter if none): ")
	for scanner.Scan() {
		var err error
		proxies, err = loadProxies(strings.ReplaceAll(scanner.Text(), "\"", ""))
		if err != nil {
			fmt.Println(err)
		}
		break
	}
}

// proxyInfo ...
type proxyInfo struct {
	IP       string
	Port     string
	Email    string
	Password string
	Auth     bool
}

// loadProxies ...
func loadProxies(filename string) ([]proxyInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		proxies []proxyInfo
		scanner = bufio.NewScanner(file)
	)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		if len(s) == 2 {
			proxies = append(proxies, proxyInfo{IP: s[0], Port: s[1]})
		} else if len(s) == 4 {
			proxies = append(proxies, proxyInfo{IP: s[0], Port: s[1], Email: s[2], Password: s[3], Auth: true})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return proxies, nil
}

// newTransportWithProxy ...
func newTransportWithProxy(proxy proxyInfo) *http.Transport {
	if proxy.Auth {
		return &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				User:   url.UserPassword(proxy.Email, proxy.Password),
				Host:   fmt.Sprintf("%s:%s", proxy.IP, proxy.Port),
			}),
		}
	}
	return &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("%s:%s", proxy.IP, proxy.Port),
		}),
	}
}
