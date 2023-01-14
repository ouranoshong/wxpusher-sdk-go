package model

// WxUser 微信用户
type WxUser struct {
	Uid        string `json:"uid"`
	CreateTime int64  `json:"createTime"`
	NickName   string `json:"nickName"`
	HeadImg    string `json:"headImg"`

	AppOrTopicId int    `json:"appOrTopicId"`
	Reject       bool   `json:"reject"`
	Id           int    `json:"id"`
	Type         int    `json:"type"`
	Target       string `json:"target"`
	PayEndTime   int    `json:"payEndTime"`
}

// QueryWxUserResult 分页查询微信用户返回结果
type QueryWxUserResult struct {
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
	Records  []WxUser `json:"records"`
}
