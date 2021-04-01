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

const tmpCookie = "sid=b6ac8ebe-f1d7-45de-8071-460194548fd4; ec=g5s2spdj-1616480328670-d49e7f968231c1248756317; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616480336; FSSBBIl1UgzbN7NO=5OlFjG_.fxShf3yo9inUrzz7ACHfPwtA6qpHzla3aAwR6z1hDQfb6OXVwrdqK2vml0aQ5Y1CEQ83srv.hHQZT4G; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1617103258; FSSBBIl1UgzbN7NP=53DWxvDrOcR9qqqmgGzCAraTM1RYZFqgILDZka.1vlVCC62Gf6QXXM8aQyPhie9OjuUu3dOoJNNGlR3ANbN0BNRwufwrvUGuBZth72MXsZdsn1hZ4z1EiXsgNBabTmdKCoInUMblaRb6yvr.h6yNJAGI9seXg2gy5ztF7yGOLVggVAzt0msDuecU3PAFoUN7vGYdN_EzwGwCuGpiZBZZM2O43Oi_L9ECP8ZbxJO0xzdNWJoAKac3qBLi0vhL3CuJomfERhi.94ng.bI8h9og69s; _efmdata=krmqwcCi5F6JC0fs0IU3OI83eFfH6kQ%2BeeLWsRiRLHuQAwAXcfd%2Bst7vGR6sJpCvDVZKsCDnl64umHg8YuQojV%2B2BYrfVo%2F9BLFBwf%2FtUpY%3D; _exid=erYATsc7EXj8yrIpeFNFiOIFF7kDJf8dwHP9JBt5x7yXzTb%2BJFeGHGCtKH2cbZP0KIJIOfYBwSrPvJTMfH1NSg%3D%3D"

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
	IpQueue = strings.Split("112.194.92.74:4245｜27.155.222.100:4267｜27.152.193.150:4246｜118.123.41.232:4218｜59.58.47.16:4245｜218.86.19.184:4242｜221.234.28.75:4278｜124.112.223.159:4203｜121.206.5.157:4224｜59.60.130.144:4242｜", "|")

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
