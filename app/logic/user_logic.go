package logic

import "store/app/enum"

//UserInfo 获取玩家个人信息
func UserInfo(userId uint64) *enum.UserInfo {
	var user enum.UserInfo

	return &user
}

//UserList 获取指定玩家列表
func UserList(idArr []uint64) []*enum.UserInfo {
	var userList []*enum.UserInfo

	return userList
}
