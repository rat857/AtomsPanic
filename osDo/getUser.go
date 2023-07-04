package osDo

import (
	"github.com/fatih/color"
	"os/user"
	"strconv"
)

func Getuser() (string, int) {
	//获取用户名和UID
	currentUser, err := user.Current()
	if err != nil {
		color.Red("%v,%s", err, "获取用户失败")
	}
	Uid, err := strconv.Atoi(currentUser.Uid)
	if err != nil {
		color.Red("%v,%s", err, "用户Uid转换为int型数据失败")
	}
	return currentUser.Name, Uid
}
