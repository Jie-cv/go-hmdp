package model

import "time"

// CREATE TABLE `tb_user`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号码',
//   `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '密码，加密存储',
//   `nick_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '昵称，默认是用户id',
//   `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '人物头像',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`),
//   UNIQUE INDEX `uniqe_key_phone`(`phone`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbUser struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Password   string    `gorm:"column:password" json:"password"`
	NickName   string    `gorm:"column:nick_name" json:"nickName"`
	Icon       string    `gorm:"column:icon" json:"icon"`
	CreateTime time.Time `gorm:"->" json:"createTime"`
	UpdateTime time.Time `gorm:"->" json:"updateTime"`
}

func (TbUser) TableName() string {
	return "tb_user"
}
