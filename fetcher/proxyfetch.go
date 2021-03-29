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

const tmpCookie = "sid=61d6b068-2873-4173-a51b-31e359072622; ec=kmYC32Ea-1617027918782-fd4810a6340531550703258; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1617027927; FSSBBIl1UgzbN7NO=55fsThGrmgLSsKxVnlBFUR_ebsf4vtD5jp_z99KZ3i2WOHi.fNzCjpBwzkzA3p_oPNS7oTsN6aWvzsWCIJqgV8a; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1617029602; _exid=tcBU05sRFkKUJGYZWZnW7BDKnMBAhMH9IsvPcU1hA%2FAzh9qekVfL1Clu0dIeH7z8X43fcNYSk%2FQ7B6dDX5%2BYlQ%3D%3D; _efmdata=M2ffXuF3m6LGxRfCJVLqYHjLX44Bb8o1QGLoo1ey2nVVwNpYpWmXSdrjGMamQ3XhBXZlQVyUPScM4xkNLHv6tHKZe1SNqYCThhOykcSTisk%3D; FSSBBIl1UgzbN7NP=53DovXbrfy49qqqmgILiqNGS1NHvTTR_GKWcoLL7AXlatq1.u_8qdR7yLBLyG_tDLPDqBpBSyG1Sx_vAqeTCWxB2JIVRYDPby0CLv39eDx66X618Rymt9VYK.FQcWOxRvAkU2SzlQSM39hr2C.a14o9ztXqgjYDQHw50Q9EpB2EycbvyA8h0_PASjMmUv_AB3V0zOAPSUWo8AGkQjQIuRriPb9j2TSkvAEIU9oz1vEHm1m2QsDEoMgHUPfU0NOidCq"

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
	IpQueue = strings.Split("223.156.84.141:4243|223.243.177.250:4272|106.6.232.61:4217|175.147.117.1:4260|110.90.220.152:4270|110.82.166.47:4245|121.238.151.142:4236|36.62.211.183:4245|117.69.182.229:4226|125.87.84.238:4246|111.72.140.29:4221|1.70.76.28:4236|124.94.251.145:4230|114.103.80.6:4210|114.101.252.238:4272|", "|")

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
