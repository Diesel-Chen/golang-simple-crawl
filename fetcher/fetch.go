package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", RandomGetUA())

	resp, err := http.DefaultClient.Do(req)

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
