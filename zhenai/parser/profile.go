package parser

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/model"
	"regexp"
)

//"basicInfo":["离异","49岁","天秤座(09.23-10.22)","158cm","58kg","工作地:合肥瑶海区","月收入:3千以下","自由职业","高中及以下"]
var proReg = regexp.MustCompile(`"basicInfo":\["([^"]+)","(\d+岁)","([^"]+)","(\d+cm)","(\d+kg)","工作地:([^"]+)","月收入:([^"]+)","([^"]+)","([^"]+)"\]`)

func ParserProfile(cnt []byte) engine.RequestResult {
	resps := proReg.FindAllSubmatch(cnt, -1)
	var result engine.RequestResult
	for _, v := range resps {
		p := model.Person{
			MarriageStatus: string(v[1]),
			Age:            string(v[2]),
			Constellation:  string(v[3]),
			Height:         string(v[4]),
			Weight:         string(v[5]),
			WorkStation:    string(v[6]),
			Salary:         string(v[7]),
			Job:            string(v[8]),
			Degree:         string(v[9]),
		}
		result.Items = append(result.Items, p)
	}
	return result
}
