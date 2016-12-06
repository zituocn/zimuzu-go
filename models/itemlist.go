package models

//下载资源接收
type ItemOut struct {
	Status int64  `json:"status"`
	Info   string `json:"info"`
	Data   []Item `json:"data"` //下载资源列表
}

//下载资源
type Item struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Format   string     `json:"format"`
	Season   string     `json:"season"`
	Episode  string     `json:"episode"`
	Size     string     `json:"size"`
	Dateline string     `json:"dateline"`
	Link     []LinkInfo `json:"link"` //下载地址列表
	Info     string     `json:"info"`
}

//下载方式及地址
type LinkInfo struct {
	Way     string `json:"way"`     //下载方式
	Address string `json:"address"` //下载地址
}
