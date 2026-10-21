package model

import "time"

// CREATE TABLE `tb_shop_type`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '类型名称',
//   `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图标',
//   `sort` int(3) UNSIGNED NULL DEFAULT NULL COMMENT '顺序',
//   `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbShopType struct {
	ID         uint64     `json:"id"`
	Name       *string    `json:"name"`
	Icon       *string    `json:"icon"`
	Sort       *uint32    `json:"sort"`
	CreateTime *time.Time `gorm:"->" json:"createTime"`
	UpdateTime *time.Time `gorm:"->" json:"updateTime"`
}

func (TbShopType) TableName() string {
	return "tb_shop_type"
}
