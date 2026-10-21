package model

import "time"

// CREATE TABLE `tb_voucher`  (
//   `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
//   `shop_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '商铺id',
//   `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '代金券标题',
//   `sub_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '副标题',
//   `rules` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '使用规则',
//   `pay_value` bigint(10) UNSIGNED NOT NULL COMMENT '支付金额，单位是分。例如200代表2元',
//   `actual_value` bigint(10) NOT NULL COMMENT '抵扣金额，单位是分。例如200代表2元',
//   `type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0,普通券；1,秒杀券',
//   `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '1,上架; 2,下架; 3,过期',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbVoucher struct {
	ID          uint64    `json:"id"`
	ShopID      *uint64   `json:"shopId"`
	Title       string    `json:"title"`
	SubTitle    *string   `json:"subTitle"`
	Rules       *string   `json:"rules"`
	PayValue    uint64    `json:"payValue"`
	ActualValue int64     `json:"actualValue"`
	Type        uint8     `json:"type"`
	Status      uint8     `gorm:"default:1" json:"status"`
	CreateTime  time.Time `gorm:"->" json:"createTime"`
	UpdateTime  time.Time `gorm:"->" json:"updateTime"`
}

func (TbVoucher) TableName() string {
	return "tb_voucher"
}
