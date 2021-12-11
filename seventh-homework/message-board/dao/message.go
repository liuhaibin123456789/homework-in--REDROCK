package dao

import (
	"message-board/models"
	"message-board/tool"
)

//AddLeavingMessage 添加留言.留言标记p_id=0,id不为默认值
func AddLeavingMessage(m models.Message) (err error) {
	sqlInsertStr := "insert into `message` (id,p_id,c_id,user_name,content) values (?,default,default,?,?)"

	_, err = tool.DB.Exec(sqlInsertStr, m.Id, m.UserName, m.Content)
	return err
}

//AddCommentMessage 添加评论,评论p_id对应留言id,id为默认值
func AddCommentMessage(m models.Message) (err error) {
	sqlInsertStr := "insert into `message` (id,p_id,c_id,user_name,content) values (default,?,?,?,?)"

	_, err = tool.DB.Exec(sqlInsertStr, m.PId,m.CId, m.UserName, m.Content)
	return err
}

//UpdateLeavingMessage 修改留言
func UpdateLeavingMessage(userName string, id int, content string) (err error) {
	sqlUpdateStr := "update message set content=? where user_name=? and id =?"

	_, err = tool.DB.Exec(sqlUpdateStr, content, userName, id)
	return err
}

//UpdateCommentMessage 修改评论
func UpdateCommentMessage(userName string, CId int, content string) (err error) {
	sqlUpdateStr := "update message set content=? where user_name=? and c_id =?"

	_, err = tool.DB.Exec(sqlUpdateStr, content, userName, CId)
	return err
}

//DeleteLeavingMessage 删除留言
func DeleteLeavingMessage(userName string, id int) (err error) {
	sqlUpdateStr := "update message set content=? where user_name=? and id =?"

	_, err = tool.DB.Exec(sqlUpdateStr, "该留言已删除~", userName, id)
	return err
}

//DeleteCommentMessage 删除评论
func DeleteCommentMessage(userName string, Cid int) (err error) {
	sqlDeleteStr := "delete from message where user_name = ? and c_id = ?"
	_, err = tool.DB.Exec(sqlDeleteStr, userName, Cid)
	return err
}

//SelectAllLeavingMessages 查看所有用户的所有留言,不包括评论
func SelectAllLeavingMessages() (AllLeavingMessages []models.Message, err error) {
	sqlSelectStr := "select * from message where c_id = 0 " //p_id为默认值，表示为留言
	MSGS := make([]models.Message,5)

	rows, err := tool.DB.Query(sqlSelectStr)
	if err != nil {
		return nil, err
	} else {
		var m models.Message
		for rows.Next() {
			err = rows.Scan(&m.Id,&m.PId,&m.CId,&m.UserName,&m.Content)
			if err != nil {
				return nil, err
			}
			MSGS = append(MSGS, m)
		}
		return MSGS, nil
	}
}

//SelectSingleLeavingMessage 查看某用户的某条留言（userName、id标记对应某条留言）
func SelectSingleLeavingMessage(userName string, id int) (LeavingMessage models.Message, err error) {
	sqlSelectStr := "select * from message where user_name=? and id=?"
	err = tool.DB.QueryRow(sqlSelectStr, userName, id).Scan(&LeavingMessage.Id,&LeavingMessage.PId,
		&LeavingMessage.CId,&LeavingMessage.UserName,&LeavingMessage.Content)
	if err != nil {
		return models.Message{}, err
	} else {
		return LeavingMessage, nil
	}
}

//SelectCommentMessages 查看某个用户所属某条留言的下属所有评论(包括留言本身)
func SelectCommentMessages(userName string, id int) (MSGS []models.Message, err error) {
	sqlSelectStr := "select * from message where p_id=? and user_name=?"
	MSGS = make([]models.Message,5)

	//查找并添加首留言至返回值切片
	firstLeavingMessage, err := SelectSingleLeavingMessage(userName, id)
	if err != nil {
		return MSGS, err
	}
	MSGS = append(MSGS, firstLeavingMessage)

	rows, err := tool.DB.Query(sqlSelectStr, id, userName)
	if err != nil {
		return MSGS, err
	}
	defer rows.Close() //延迟释放资源

	//追加后续评论至返回值切片
	var m models.Message
	for rows.Next() {
		err = rows.Scan(&m.Id, &m.PId,&m.CId,&m.UserName,&m.Content)
		if err != nil {
			return MSGS, err
		}
		MSGS = append(MSGS, m)
	}
	return MSGS, err
}
