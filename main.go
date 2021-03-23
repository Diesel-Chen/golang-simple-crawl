package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main(){
	url:="https://www.zhenai.com/zhenghun"
	req,_:=http.NewRequest(http.MethodGet,url,nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
	resp,err:=http.DefaultClient.Do(req)

	if err!=nil{
		panic("http.Get Error")
	}
	defer resp.Body.Close()
	if resp.StatusCode!=http.StatusOK{
		fmt.Println("status code ",resp.StatusCode)
		panic("Error status code")
	}
	body,_:=ioutil.ReadAll(resp.Body)
	process(body)

}
func process(cnt []byte){
	reg:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	resp:=reg.FindAllSubmatch(cnt,-1)
	for _,val:=range resp{
		fmt.Printf("url:%s  cityname:%s\n",val[1],val[2])
	}
	fmt.Printf("total count %d\n",len(resp))
}
