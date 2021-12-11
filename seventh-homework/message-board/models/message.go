package models

//Message 留言就是id标识，评论就是p_id标识
type Message struct {
	Id       int    `json:"id" form:"id"`//默认为-1表示，不是留言，表示评论，标签留言的唯一性
	PId      int    `json:"p_id" form:"p_id"`//标签评论的唯一性
	CId      int    `json:"c_id" form:"c_id"` //默认为零表示，不是评论，表示最顶一级的留言
	UserName string `json:"user_name" form:"user_name"`
	Content  string `json:"content" form:"content"`
}
