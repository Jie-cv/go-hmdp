package model

import "time"

// CREATE TABLE `tb_user_info`  (
//   `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '主键，用户id',
//   `city` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '城市名称',
//   `introduce` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '个人介绍，不要超过128个字符',
//   `fans` int(8) UNSIGNED NULL DEFAULT 0 COMMENT '粉丝数量',
//   `followee` int(8) UNSIGNED NULL DEFAULT 0 COMMENT '关注的人的数量',
//   `gender` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '性别，0：男，1：女',
//   `birthday` date NULL DEFAULT NULL COMMENT '生日',
//   `credits` int(8) UNSIGNED NULL DEFAULT 0 COMMENT '积分',
//   `level` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '会员级别，0~9级,0代表未开通会员',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`user_id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbUserInfo struct {
	UserID     uint64     `gorm:"primaryKey;autoIncrement:false" json:"userId"`
	City       *string    `json:"city"`
	Introduce  *string    `json:"introduce"`
	Fans       *uint32    `json:"fans"`
	Followee   *uint32    `json:"followee"`
	Gender     *uint8     `json:"gender"`
	Birthday   *time.Time `json:"birthday"`
	Credits    *uint32    `json:"credits"`
	Level      *uint8     `json:"level"`
	CreateTime time.Time  `gorm:"->" json:"createTime"`
	UpdateTime time.Time  `gorm:"->" json:"updateTime"`
}

func (TbUserInfo) TableName() string {
	return "tb_user_info"
}
