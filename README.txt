# homework-in--REDROCK

some thinking of coding...

          p_id指的是添加在某个留言对应的评论id->目前只实现留言以及留言对应的所有同级评论，还在思考如何实现套娃评论...
          (感觉数据库设计的不合理。。。

sql语句：
#留言就是id标识，评论就是c_id标识
#p_id表示所属留言id,一级套娃
CREATE TABLE `message`(
    `id` INT  DEFAULT -1 ,#默认为-1表示，不是留言，表示评论，标签留言的唯一性
    `p_id` INT  DEFAULT 0, #默认为零表示，不是评论，表示最顶一级的留言，评论的p_id指向表示对应id的留言。
    `c_id` INT DEFAULT 0, #标签评论的唯一性
    `user_name` VARCHAR(20),
    `content` VARCHAR(255)#表示评论或留言的内容
)CHARSET=utf8;
#留言和评论放在一张表上
DROP TABLE `message`;

CREATE TABLE `user`(
     `id` INT AUTO_INCREMENT PRIMARY KEY,
     `account` VARCHAR(20) NOT NULL UNIQUE,
     `password` VARCHAR(16) NOT NULL
)CHARSET=utf8;

DROP TABLE `message`;

SELECT * FROM `message`;
SELECT * FROM `user`;
DELETE FROM message;
DELETE FROM `user`;
