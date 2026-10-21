package model

import "time"

// CREATE TABLE `tb_follow`  (
//   `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
//   `follow_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '关联的用户id',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbFollow struct {
	ID           int64     `json:"id"`
	UserID       uint64    `json:"user_id"`
	FollowUserID uint64    `json:"follow_user_id"`
	CreateTime   time.Time `gorm:"->" json:"create_time"`
}

func (TbFollow) TableName() string {
	return "tb_follow"
}
