package tool

import (
	"errors"
)

//CheckId 检查留言是否存在，注意留言删除时的情况
func CheckId(userName string, id int) (err error) {
	if id != -1 {
		sqlSelect := "select content from message where user_name = ? and id = ?"
		rows, err := DB.Query(sqlSelect, userName, id)
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
			return errors.New("the message is wrong ,id should be different from others")
		}
	}
	return errors.New("the message is not correct leaving message")
}
