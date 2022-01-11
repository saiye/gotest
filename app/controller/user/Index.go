package user

import "store/app/logic"

func Login(account string, password string) error {
	//todo
	return logic.Login(account, password)
}
