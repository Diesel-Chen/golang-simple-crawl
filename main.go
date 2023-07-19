package main

import (
	"github.com/PuerkitoBio/goquery"
	"golang-simple-crawl/model"
	"net/http"
	"regexp"
	"strings"
)

func CrawlHtml(url string) *goquery.Document {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

func GetCityList() []model.CityInfo {
	res := make([]model.CityInfo, 0)
	doc := CrawlHtml("https://www.zhenai.com/zhenghun")
	doc.Find(".city-list dd a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		res = append(res, model.CityInfo{
			Name: s.Text(),
			Url:  href,
		})

	})
	return res
}

func GetPersonByCity(cities []model.CityInfo) []model.Person {
	res := make([]model.Person, 0)

	for _, city := range cities {
		doc := CrawlHtml(city.Url)
		doc.Find(".g-list .list-item").Each(func(i int, s *goquery.Selection) {

			avatar, _ := s.Find(".photo img").Attr("src")
			personInfo, _ := s.Find(".photo a").Attr("href")
			reg := regexp.MustCompile(`/u/(\d+)`)
			match := reg.FindStringSubmatch(personInfo)
			if len(match) < 2 {
				panic("未找到匹配项")
			}
			id := match[1]

			nick := s.Find(".content th a").Text()
			tmpPerson := model.Person{
				Id:        id,
				Avatar:    avatar,
				Nick:      nick,
				City:      city.Name,
				Signature: s.Find(".content .introduce").Text(),
			}

			s.Find(".content tr").Each(func(i1 int, s1 *goquery.Selection) {
				s1.Find("td").Each(func(i2 int, s2 *goquery.Selection) {
					attr := s2.Text()
					index := strings.Index(attr, "：")

					updateAttr := attr[index+3:]

					switch i1 {

					case 1:
						switch i2 {
						case 0:
							tmpPerson.Gender = updateAttr
						case 1:
							tmpPerson.Residence = updateAttr
						}

					case 2:
						switch i2 {
						case 0:
							tmpPerson.Age = updateAttr
						case 1:
							tmpPerson.Salary = updateAttr
						}
					case 3:
						switch i2 {
						case 0:
							tmpPerson.MarriageStatus = updateAttr
						case 1:
							tmpPerson.Height = updateAttr
						}
					}

				})

			})

			res = append(res, tmpPerson)

		})

	}
	return res
}

func main() {
	cityList := GetCityList()
	persons := GetPersonByCity(cityList[:3])
	if err := model.WriteToCSV(persons); err != nil {
		panic(err)
	}

}
