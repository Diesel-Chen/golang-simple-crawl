package parser

import (
	"golang-simple-crawl/engine"
	"golang-simple-crawl/model"
	"regexp"
)

//"basicInfo":["离异","49岁","天秤座(09.23-10.22)","158cm","58kg","工作地:合肥瑶海区","月收入:3千以下","自由职业","高中及以下"]
//"basicInfo":["丧偶","33岁","天秤座(09.23-10.22)","158cm","工作地:广州花都区","月收入:8千-1.2万","经销商","大专"],
var proReg = regexp.MustCompile(`"basicInfo":\["([^"]+)","(\d+岁)","([^"]+)","(\d+cm)","(\d+kg)","工作地:([^"]+)","月收入:([^"]+)","([^"]+)","([^"]+)"\]`)
var proReg1 = regexp.MustCompile(`"basicInfo":\["([^"]+)","(\d+岁)","([^"]+)","(\d+cm)","工作地:([^"]+)","月收入:([^"]+)","([^"]+)","([^"]+)"\]`)
var weightReg = regexp.MustCompile(`basicInfo":\[.+kg.+\]`)

func ParserProfile(name string, cnt []byte) engine.RequestResult {

	var result engine.RequestResult
	isWeights := weightReg.FindAllString(string(cnt), 1)
	if len(isWeights) > 0 {
		resps := proReg.FindAllSubmatch(cnt, -1)
		for _, v := range resps {
			p := model.Person{
				Name:           name,
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
	} else {
		resps := proReg1.FindAllSubmatch(cnt, -1)
		for _, v := range resps {
			p := model.Person{
				Name:           name,
				MarriageStatus: string(v[1]),
				Age:            string(v[2]),
				Constellation:  string(v[3]),
				Height:         string(v[4]),
				WorkStation:    string(v[5]),
				Salary:         string(v[6]),
				Job:            string(v[7]),
				Degree:         string(v[8]),
			}
			result.Items = append(result.Items, p)
		}
	}

	return result
}
