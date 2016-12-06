package models

//影片列表的json接收
type ResourceListOut struct {
	Status int64        `json:"id"`
	Info   string       `json:"info"`
	Data   ResourceList `json:"data"`
}

//影片列表
type ResourceList struct {
	Count int64      `json:"count"`
	List  []Resource `json:"list"`
}
