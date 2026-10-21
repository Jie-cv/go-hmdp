package model

import "time"

// CREATE TABLE `tb_sign`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
//   `year` year NOT NULL COMMENT '签到的年',
//   `month` tinyint(2) NOT NULL COMMENT '签到的月',
//   `date` date NOT NULL COMMENT '签到的日期',
//   `is_backup` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '是否补签',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbSign struct {
	ID       uint64    `json:"id"`
	UserID   uint64    `json:"userId"`
	Year     uint16    `json:"year"`
	Month    int8      `json:"month"`
	Date     time.Time `json:"date"`
	IsBackup *uint8    `json:"isBackup"`
}

func (TbSign) TableName() string {
	return "tb_sign"
}
