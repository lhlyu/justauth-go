package model

type AuthUser struct {
	UUID     string     `json:"uuid"`
	UserName string     `json:"userName"`
	NickName string     `json:"nickName"`
	Avatar   string     `json:"avatar"`
	Company  string     `json:"company"`
	Blog     string     `json:"blog"`
	Location string     `json:"location"`
	Email    string     `json:"email"`
	Remark   string     `json:"remark"`
	Gender   string     `json:"gender"`
	Source   string     `json:"source"`
	Token    *AuthToken `json:"token"`
}
