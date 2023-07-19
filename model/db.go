package model

type CityInfo struct {
	Name string
	Url  string
}

type Person struct {
	Id             string `json:"id"`             //id
	Nick           string `json:"nick"`           //昵称
	Avatar         string `json:"avatar"`         //头像
	Gender         string `json:"gender"`         //性别
	Age            string `json:"age"`            //年龄
	City           string `json:"city"`           //城市
	Residence      string `json:"residence"`      //居住地
	Salary         string `json:"salary"`         //月薪
	Height         string `json:"height"`         //高度
	MarriageStatus string `json:"marriageStatus"` //婚姻状态
	Signature      string `json:"signature"`      //个签
}
