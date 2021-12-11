package tool

import "errors"

//CheckCId 校验评论是否存在
func CheckCId(userName string, Pid int) (err error) {
	if Pid != 0 {
		sqlSelect := "select content from message where user_name = ? and p_id = ?"
		rows, err := DB.Query(sqlSelect, userName, Pid)
		if err != nil {
			return err
		}
		var Lines = 0 //记录相同留言的行数
		for rows.Next() {
			var c string
			err = rows.Scan(&c)
			if err != nil {
				return err
			}
			if c != "该留言已删除~" {
				Lines++
				break
			} else {
				continue
			}
		}
		//fmt.Println("相同的lines行数：",Lines)
		if Lines == 0 {
			return nil
		} else {
			return errors.New("the comment is wrong ,id should be different from others")
		}
	}
	return errors.New("the message is not comment message")
}
