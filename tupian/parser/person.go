package parser

import (
	"fmt"
	"golang-simple-crawl/engine"
	"golang-simple-crawl/model"
	"regexp"
)

//<img src="http://nbvvv.8iwvsl.com/15031B/p01/%E5%96%9D%E9%86%89%E5%90%8E%E8%A2%AB%E7%8B%82%E7%8E%A92[21P]/01.jpg" border="0">
var perReg = regexp.MustCompile(`<img src="(http://[^\.]+\.8iwvsl[^"]+)" border="0">`)

func ParserPerson(cnt []byte, name string) engine.RequestResult {
	resp := perReg.FindAllSubmatch(cnt, -1)
	result := engine.RequestResult{}
	//fmt.Println("resp:", resp)
	//fmt.Println("name:", name)
	for _, v := range resp {
		fmt.Println("url:", string(v[1]))
		picture := model.Picture{
			Name: name,
			Url:  string(v[1]),
		}
		result.Items = append(result.Items, picture)
	}
	return result
}
