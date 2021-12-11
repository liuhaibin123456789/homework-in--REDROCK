package service

import (
	"errors"
	"fmt"
	"message-board/dao"
	"message-board/models"
	"message-board/tool"
)

func AddLeavingMessage(userName string, id int, content string) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	}
	//校验留言是否已经存在
	err = tool.CheckId(userName, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	//初始化结构体
	var m models.Message
	m.Id = id
	m.Content = content
	m.UserName = userName
	err = dao.AddLeavingMessage(m)
	return err
}

func AddCommentMessage(userName string, Pid int,Cid int, content string) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	} else {
		if Cid==0 {
			return errors.New("cid is wrong")
		}
		//校验评论是否已经存在
		err = tool.CheckCId(userName, Cid)
		if err != nil {
			fmt.Println(err)
			return err
		}
		//初始化结构体
		var m models.Message
		m.PId = Pid
		m.CId=Cid
		m.Content = content
		m.UserName = userName
		err = dao.AddCommentMessage(m)
		if err != nil {
			return err
		}
	}
	return err
}

func UpdateLeavingMessage(userName string, id int, content string) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	} else {
		//校验留言是否已经存在
		err = tool.CheckId(userName, id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = dao.UpdateLeavingMessage(userName, id, content)
		if err != nil {
			return err
		}
	}
	return err
}

func UpdateCommentMessage(userName string, cId int, content string) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	} else {
		//校验评论是否已经存在
		err = tool.CheckCId(userName, cId)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = dao.UpdateCommentMessage(userName, cId, content)
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteLeavingMessage(userName string, id int) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	} else {
		//校验评论是否已经存在
		err = tool.CheckId(userName, id)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = dao.DeleteLeavingMessage(userName, id)
		if err != nil {
			return err
		}
	}
	return err
}

func DeleteCommentMessage(userName string, Cid int) (err error) {
	//校验账户是否存在
	err = dao.IsExist(userName)
	if err != nil {
		return errors.New("the user don`t exist")
	} else {
		//校验评论是否已经存在
		err = tool.CheckCId(userName, Cid)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = dao.DeleteCommentMessage(userName, Cid)
		if err != nil {
			return err
		}
	}
	return err
}

func SelectAllLeavingMessages() (AllLeavingMessages []models.Message, err error) {
	AllLeavingMessages, err = dao.SelectAllLeavingMessages()
	return AllLeavingMessages, err
}

func SelectCommentMessages(userName string, id int) (singleMessages []models.Message, err error) {
	singleMessages, err = dao.SelectCommentMessages(userName, id)
	return singleMessages, err
}
