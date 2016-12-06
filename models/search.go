package models

//搜索结果的json接收
type SearchItemListOut struct {
	Status int64          `json:"status"`
	Info   string         `json:"info"`
	Data   SearchItemList `json:"data"`
}

//包括记录总数的搜索结果列表
type SearchItemList struct {
	Count int64        `json:"count"`
	List  []SearchItem `json:"list"`
}

//搜索结果数据项
type SearchItem struct {
	Itemid  string `json:"itemid"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Pubtime string `json:"pubtime"`
	Uptime  string `json:"uptime"`
}
