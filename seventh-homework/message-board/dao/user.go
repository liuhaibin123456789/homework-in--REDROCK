package dao

import (
	"log"
	"message-board/models"
	"message-board/tool"
)

func AddUser(u models.User) (err error) {
	sqlInsertStr := "insert into `user` (`account`,`password`) values (?,?)"
	_, err = tool.DB.Exec(sqlInsertStr, u.Account, u.Password)
	return err
}

func UpdatePwd(account, newPwd string) (err error) {
	sqlUpdateStr := "update `user` set password=? where account=?"
	_, err = tool.DB.Exec(sqlUpdateStr, newPwd, account)
	return err
}

//SelectPwd 查询对应账户密码
func SelectPwd(account string) (string, error) {
	sqlSelectStr := "select password from `user` where account= ?"
	var pwd string
	err := tool.DB.QueryRow(sqlSelectStr, account).Scan(&pwd)
	if err != nil {
		log.Println(err)
		return "", err
	} else {
		return pwd, nil
	}
}

//IsExist 检查用户是否存在
func IsExist(account string) (err error) {
	sqlSelectStr := "select password from `user` where account= ?"
	var pwd string
	err = tool.DB.QueryRow(sqlSelectStr, account).Scan(&pwd)
	if err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

//CheckUser 校验账户密码
func CheckUser(account, password string) (bool, error) {
	pwd, err := SelectPwd(account)
	if password == pwd {
		return true, err
	} else {
		return false, err
	}
}
