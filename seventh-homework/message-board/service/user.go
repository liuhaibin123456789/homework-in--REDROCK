package service

import (
	"message-board/dao"
	"message-board/models"
	"net/http"
)

//Register 注册服务
func Register(u models.User) (bool, error) {
	err := dao.AddUser(u)
	//  鉴于本层代码太少，错误就不上抛api层处理，就地解决...   ❌❌❌❌
	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

//LoginByPwd 密码登录
func LoginByPwd(account, password string) (bool, error) {
	pwd, err := dao.SelectPwd(account)
	if pwd == password {
		return true, err
	}
	return false, err
}

//LoginByCookie cookie登录
func LoginByCookie(account, password string) (*http.Cookie, error) {
	res, err := LoginByPwd(account, password)

	if err != nil && res {
		return &http.Cookie{Name: "message_cookie", Value: account, Path: "/", HttpOnly: true, MaxAge: 1}, err
	} else {
		return &http.Cookie{}, err //todo 空cookie是否会引起error???
	}
}

//UpdatePwd 更新密码
func UpdatePwd(account, newPwd string) (err error) {
	err = dao.UpdatePwd(account, newPwd)
	return err
}

//CheckUser 核查账户及密码
func CheckUser(account, password string) (bool, error) {
	return dao.CheckUser(account, password)
}
