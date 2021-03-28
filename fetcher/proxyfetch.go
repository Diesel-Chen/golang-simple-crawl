package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const tmpCookie = "sid=b6ac8ebe-f1d7-45de-8071-460194548fd4; ec=g5s2spdj-1616480328670-d49e7f968231c1248756317; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616480336; FSSBBIl1UgzbN7NO=5OlFjG_.fxShf3yo9inUrzz7ACHfPwtA6qpHzla3aAwR6z1hDQfb6OXVwrdqK2vml0aQ5Y1CEQ83srv.hHQZT4G; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616949353; _efmdata=krmqwcCi5F6JC0fs0IU3OI83eFfH6kQ%2BeeLWsRiRLHuQAwAXcfd%2Bst7vGR6sJpCvzJdrPGI4rbubQB7ob%2F%2BxywbiE9BpKaOuZszbEjJoLoU%3D; _exid=vSjgxarHDrnB%2FUSRRiJZaSwBV3a%2FGP58j1KkmTqxu3pNHcpi2PfjwDWm3nvbmPVau%2BOlQAKtDWCRT0fInWwlJg%3D%3D; FSSBBIl1UgzbN7NP=53Dca4CrG7iaqqqmg8MWvpGJkq.A2eCokYHgVrENMikFAyUWji8RCcmEayJHgyoa_mdzRX1agCR8PZj6pRIbHK78Q5gGERtjuPQFS5UsEqfwpQNmg_LlkmZq7FzQzXkkf3r8gsvKkXt5GFDHK6QaB28zuiFWAzohNe9RX7QxYxyIb40LL3TN2Olhx.jHva.UIE_P2_mriXW6jNxgxG59cpMOxMhjz81nX74OXeXsdZgj7jg8eEZDq6X5AQP9pVDWdX20W4vIaRCSfmDw6Q3tOeP"

var ProxyServer string
var IpQueue []string

func ProxyFetch(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", RandomGetUA())

	req.Header.Add("cookie", tmpCookie)
	client := ProxyClient()
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("client.Do出错 返回码：")
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err status code %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
func ProxyClient() http.Client {

	go func() {
		for {
			time.Sleep(100 * time.Microsecond)
			getRandomIp()
		}
	}()

	proxyURL, _ := url.Parse(ProxyServer)

	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
}
func init() {

	//resp, err := http.Get("http://webapi.http.zhimacangku.com/getip?num=40&type=1&pro=&city=0&yys=0&port=1&time=1&ts=0&ys=0&cs=0&lb=6&sb=|&pb=4&mr=1&regions=")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	IpQueue = strings.Split("106.111.136.145:4267|119.36.156.69:4213|49.88.235.129:4214|36.6.58.68:4270|180.111.139.155:4278|120.34.12.234:4232|101.23.115.65:4223|223.243.48.126:4216|175.167.20.106:4270|115.211.39.233:4274|123.186.228.151:4223|60.184.167.83:4245|121.233.161.253:4262|114.99.0.197:4213|59.62.25.242:4264|182.100.138.74:4212|223.243.54.157:4278|183.164.238.27:4225|118.123.43.157:4284|27.44.220.169:4265|", "|")
	ProxyServer = IpQueue[0]
}

func getRandomIp() {
	rand.Seed(time.Now().UnixNano())
	target := IpQueue[rand.Intn(len(IpQueue))]
	ProxyServer = "http://" + target
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
