package model

type AuthUser struct {
	UUID     string    `json:"uuid"`
	UserName string    `json:"userName"`
	NickName string    `json:"nickName"`
	Avatar   string    `json:"avatar"`
	Blog     string    `json:"blog"`
	Location string    `json:"location"`
	Email    string    `json:"email"`
	Remark   string    `json:"remark"`
	Gender   string    `json:"gender"`
	Source   string    `json:"source"`
	token    AuthToken `json:"token"`
}
