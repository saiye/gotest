package user

import "store/app/logic"

//Login 登录
func Login(account string, password string) error {
	//todo
	return logic.Login(account, password)
}

//Logout 登出
func Logout(token string) error {
	return nil
}

//Lock 锁定用户，多少秒:0是永久锁定
func Lock(uid uint64, lockTime int64) int64 {
	var endTime int64
	return endTime
}

//Unlock 解锁用户
func Unlock(uid uint64) error {
	return nil
}
