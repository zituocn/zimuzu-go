package main

import (
	"fmt"
	"github.com/zituocn/zimuzu-go/models"
	"github.com/zituocn/zimuzu-go/utils"
)

func main() {

	//测试搜索接口
	data, err := GetSearchList("西部", "resource", "pubtime", 20, 1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data.List)
}

//获取搜索结果
//k			关键字
//st 		搜索类型(resource=影视资源|subtitle=字幕资源|article=资讯以及影评和剧评)
//order 	排序方式(pubtime|uptime)
//limit	 	每页数量
//page  	页码
func GetSearchList(k string, st string, order string, limit int64, page int64) (data *models.SearchItemList, err error) {
	var (
		url = fmt.Sprintf("/search?k=%s&st=%s&order=%s&limit=%d&page=%d", k, st, order, limit, page)
		out models.SearchItemListOut
	)
	req := utils.HttpGet(url)
	req.ToJSON(&out)
	return &out.Data, err
}

//获取影片列表
//page 	   页码
//limit    每页数量，<=20
//channel  频道(movie|tv|openclass)
//area     国家(美国|日本|中国|英国)
//sort     排序方式(update|pubdate|premiere|name|rank|score|views)
//year     年代(最小为1990)
//category 影片类型(未知-:( )
func GetResourceList(page int64, limit int64, more ...string) (data *models.ResourceList, err error) {
	var (
		url   = fmt.Sprintf("/resource/fetchlist?page=%d&limit=%d", page, limit)
		out   models.ResourceListOut
		param string
	)
	for i, v := range more {
		if i == 0 {
			param = "&channel=%s"
		}
		if i == 1 {
			param = "&area=%s"
		}
		if i == 2 {
			param = "&sort=%s"
		}
		if i == 3 {
			param = "&year=%s"
		}
		if i == 4 {
			param = "&category=%s"
		}

		url += fmt.Sprintf(param, v)
	}
	req := utils.HttpGet(url)
	req.ToJSON(&out)
	return &out.Data, err
}

//
//获取某个影片下载信息及地址
//id		影片id
//file		是否同时显示下载方式及地址 0/1
//season    第x季
//episode   第x集
func GetItemInfo(id int64, file int64, season int64, episode int64) (data *[]models.Item, err error) {
	var (
		url = fmt.Sprintf("/resource/itemlist_web?id=%d&file=%d&season=%d&episode=%d", id, file, season, episode)
		out models.ItemOut
	)

	req := utils.HttpGet(url)
	err = req.ToJSON(&out)
	return &out.Data, err

}

//
//获取某个影片的信息
//id		影片id
func GetResourceInfo(id int64) (data *models.Resource, err error) {
	var (
		url = fmt.Sprintf("/resource/getinfo?id=%d", id)
		out models.ResourceOut
	)
	req := utils.HttpGet(url)
	err = req.ToJSON(&out)
	return &out.Data, err

}
