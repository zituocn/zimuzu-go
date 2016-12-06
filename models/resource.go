package models

//影片_接收json
type ResourceOut struct {
	Status int64    `json:"id"`
	Info   string   `json:"info"`
	Data   Resource `json:"data"`
}

//影片详情
type Resource struct {
	Id              string `json:"id"`
	Cnanme          string `json:"cnname"`
	Enname          string `json:"enname"`
	Poster          string `json:"poster"`
	Play_status     string `json:"play_status"`
	Area            string `json:"area"`
	Category        string `json:"category"`
	Views           string `json:"views"`
	Score           string `json:"score"`
	Channel         string `json:"channel"`
	Premiere        string `json:"premiere"`
	Remark          string `json:"remark"`
	Rank            string `json:"rank"`
	Lang            string `json:"lang"`
	Content         string `json:"content"`
	Poster_a        string `json:"poster_a"`
	Poster_b        string `json:"poster_b"`
	Poster_m        string `json:"poster_m"`
	Poster_s        string `json:"poster_s"`
	Itemupdate      string `json:"itemupdate"`
	Item_permission int64  `json:"item_permission"`
}
