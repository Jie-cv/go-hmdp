package model

import "time"

// CREATE TABLE `tb_shop`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商铺名称',
//   `type_id` bigint(20) UNSIGNED NOT NULL COMMENT '商铺类型的id',
//   `images` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商铺图片，多个图片以\',\'隔开',
//   `area` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '商圈，例如陆家嘴',
//   `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '地址',
//   `x` double UNSIGNED NOT NULL COMMENT '经度',
//   `y` double UNSIGNED NOT NULL COMMENT '维度',
//   `avg_price` bigint(10) UNSIGNED NULL DEFAULT NULL COMMENT '均价，取整数',
//   `sold` int(10) UNSIGNED ZEROFILL NOT NULL COMMENT '销量',
//   `comments` int(10) UNSIGNED ZEROFILL NOT NULL COMMENT '评论数量',
//   `score` int(2) UNSIGNED ZEROFILL NOT NULL COMMENT '评分，1~5分，乘10保存，避免小数',
//   `open_hours` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '营业时间，例如 10:00-22:00',
//   `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`),
//   INDEX `foreign_key_type`(`type_id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbShop struct {
	ID         uint64     `json:"id"`
	Name       string     `json:"name"`
	TypeID     uint64     `json:"typeId"`
	Images     string     `json:"images"`
	Area       *string    `json:"area"`
	Address    string     `json:"address"`
	X          float64    `json:"x"`
	Y          float64    `json:"y"`
	AvgPrice   *uint64    `json:"avgPrice"`
	Sold       uint32     `json:"sold"`
	Comments   uint32     `json:"comments"`
	Score      uint32     `json:"score"`
	OpenHours  *string    `json:"openHours"`
	CreateTime *time.Time `gorm:"->" json:"createTime"`
	UpdateTime *time.Time `gorm:"->" json:"updateTime"`
}

func (TbShop) TableName() string {
	return "tb_shop"
}
