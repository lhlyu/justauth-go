package model

type User struct {
	Token Token `json:"token"`

	UUID      string `json:"uuid"`
	UserName  string `json:"userName"`
	NickName  string `json:"nickName"`
	Avatar    string `json:"avatar"`
	Company   string `json:"company"`
	Blog      string `json:"blog"`
	Location  string `json:"location"`
	Email     string `json:"email"`
	Remark    string `json:"remark"`
	Url       string `json:"url"`
	Source    string `json:"source"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	// 原始信息json
	Original string `json:"original"`
}
