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

const tmpCookie = "sid=b6ac8ebe-f1d7-45de-8071-460194548fd4; ec=g5s2spdj-1616480328670-d49e7f968231c1248756317; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616480336; FSSBBIl1UgzbN7NO=5OlFjG_.fxShf3yo9inUrzz7ACHfPwtA6qpHzla3aAwR6z1hDQfb6OXVwrdqK2vml0aQ5Y1CEQ83srv.hHQZT4G; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616945573; _efmdata=krmqwcCi5F6JC0fs0IU3OI83eFfH6kQ%2BeeLWsRiRLHuQAwAXcfd%2Bst7vGR6sJpCvD9Ed0YcSUw7uUs3bFmaOEGNDrwUa7K5uM4gAS1UD7fk%3D; _exid=jB7TDeeqsOp2OnuT3Ztwn4vlfohVRMRSVXz4fCy2bUBKBTtHxe%2FYWk%2Fa6J52Hfq94ZH6844hTcnbMGvYJmTgrQ%3D%3D; FSSBBIl1UgzbN7NP=53DcPiKrGLO3qqqmg8UbXjqKsy45DUSBND8S.JkopA6urQlQkPytSmv1Vbo9uc9GWYfHxRkGIONcwtPckjYgPeoXj90qhjFKzyu8mUB_j3NP_zl2uvHSdv5CXm1JVkkZ8gDTezdNTpDZ29eLvCk32eYFqTYuJYzEBq1Mjdj.jFnBMW_WQLRUfz0DhGK_tnZatF9yiKEcGsO3Tn3P8qIEyiXyfHvv5gUgyhxTzTonDKQPUyBstuB0b4LlgEHkVOCCkg"

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
	IpQueue = strings.Split("49.87.133.239:4267|222.161.140.28:4278|113.141.222.241:4236|113.237.247.95:4252|117.95.46.6:4278|114.99.13.11:4225|111.173.38.129:4245|27.156.185.251:4235|117.95.1.216:4267|175.4.112.120:4230|115.209.126.169:4226|221.234.29.60:4245|119.52.132.20:4278|119.36.14.229:4213|42.84.163.8:4264|180.122.32.190:4226|125.87.89.130:4278|183.166.133.190:4214|111.127.116.141:4235|171.95.155.115:4284|", "|")
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
