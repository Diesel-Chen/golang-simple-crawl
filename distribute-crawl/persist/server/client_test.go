package main

import (
	"golang-simple-crawl/crawl/model"
	"golang-simple-crawl/distribute-crawl/config"
	"golang-simple-crawl/distribute-crawl/rpcsupport"
	"testing"
	"time"
)

func TestCallItemSaverRpc(t *testing.T) {
	//起itemsaver服务端
	host := ":2222"
	go ItemSaverRpc(host, "test")
	time.Sleep(time.Second)
	//起客户端
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	excepted := model.Person{
		Id:             "123123",
		Url:            "http://test/123123",
		Name:           "test",
		Age:            18,
		Height:         160,
		Weight:         60,
		Constellation:  "处女座",
		Salary:         "10000-20000",
		Job:            "老师",
		Degree:         "本科",
		WorkStation:    "上海",
		MarriageStatus: "未婚",
	}
	result := ""
	err = client.Call(config.ItemSaverRPC, excepted, &result)
	if err != nil || result != "ok" {
		t.Errorf("call error . result %s error %s", result, err)
	}

	//验证数据是否插入
}
