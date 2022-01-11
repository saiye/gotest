package enum

type UserInfo struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Sex      int    `json:"sex"`
	Age      int64  `json:"age"`
}
