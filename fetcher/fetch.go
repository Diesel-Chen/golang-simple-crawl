package fetcher

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const tmpCookie = "sid=b6ac8ebe-f1d7-45de-8071-460194548fd4; ec=g5s2spdj-1616480328670-d49e7f968231c1248756317; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616480336; FSSBBIl1UgzbN7NO=5OlFjG_.fxShf3yo9inUrzz7ACHfPwtA6qpHzla3aAwR6z1hDQfb6OXVwrdqK2vml0aQ5Y1CEQ83srv.hHQZT4G; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616772650; _efmdata=krmqwcCi5F6JC0fs0IU3OI83eFfH6kQ%2BeeLWsRiRLHuQAwAXcfd%2Bst7vGR6sJpCvIlB3r5nbgsZVZyAYktQU4wjt3VYez2ylOiPl0uQqKjk%3D; _exid=ZBRMpi8orCq3GQ5noR%2BqjrOejq6YLP2wRUxLQjHKjTD7YU%2BP9GOcBQaVAM2JFwsRglNQTFBJxIRE7mFKz3NsyA%3D%3D; FSSBBIl1UgzbN7NP=53m6eTKrwdwEqqqmgMnzwmAz9sZSHPF96tnd0I6mVVaBObuVGsTa5po.2lR2rO7Yi0yoJu5v_WyJQHttWEeCoesU6.yvBUx.RzgmxNAneyoMzNHRzBDPa216eGpCxMrYpLaqgX3WGVAHte2J_mmqwjhAO4lwoWBb9jvNbskrIWjHXnERYuFPm7790eAq.aDD_ZsmVINFP5RZCQialO.OxoN8ediLKFuikffNO1MyrI8XZANy8It2DxaJvUmrie45JzvSP8WVxPY0PiRaHFPo65F"

func Fetch(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", RandomGetUA())

	req.Header.Add("cookie", tmpCookie)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err status code %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
func ProxyClient() http.Client {
	ProxyServer := ""
	proxyURL, _ := url.Parse(ProxyServer)

	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}

}

func RandomGetUA() string {
	agent := []string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
	rand.Seed(time.Now().UnixNano())
	return agent[rand.Intn(len(agent))]
}
