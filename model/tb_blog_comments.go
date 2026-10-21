package model

// CREATE TABLE `tb_blog_comments`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
//   `blog_id` bigint(20) UNSIGNED NOT NULL COMMENT '探店id',
//   `parent_id` bigint(20) UNSIGNED NOT NULL COMMENT '关联的1级评论id，如果是一级评论，则值为0',
//   `answer_id` bigint(20) UNSIGNED NOT NULL COMMENT '回复的评论id',
//   `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '回复的内容',
//   `liked` int(8) UNSIGNED NULL DEFAULT NULL COMMENT '点赞数',
//   `status` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '状态，0：正常，1：被举报，2：禁止查看',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

import "time"

type TbBlogComments struct {
	ID         uint64    `json:"id"`
	UserID     uint64    `json:"userId"`
	BlogID     uint64    `json:"blogId"`
	ParentID   uint64    `json:"parentId"`
	AnswerID   uint64    `json:"answerId"`
	Content    string    `json:"content"`
	Liked      *uint32   `json:"liked"`
	Status     *uint8    `json:"status"`
	CreateTime time.Time `gorm:"->" json:"createTime"`
	UpdateTime time.Time `gorm:"->" json:"updateTime"`
}

func (TbBlogComments) TableName() string {
	return "tb_blog_comments"
}
