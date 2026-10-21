package model

import "time"

// CREATE TABLE `tb_voucher_order`  (
//   `id` bigint(20) NOT NULL COMMENT '主键',
//   `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '下单的用户id',
//   `voucher_id` bigint(20) UNSIGNED NOT NULL COMMENT '购买的代金券id',
//   `pay_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '支付方式 1：余额支付；2：支付宝；3：微信',
//   `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '订单状态，1：未支付；2：已支付；3：已核销；4：已取消；5：退款中；6：已退款',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
//   `pay_time` timestamp NULL DEFAULT NULL COMMENT '支付时间',
//   `use_time` timestamp NULL DEFAULT NULL COMMENT '核销时间',
//   `refund_time` timestamp NULL DEFAULT NULL COMMENT '退款时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Compact;

type TbVoucherOrder struct {
	ID         int64      `gorm:"primaryKey;autoIncrement:false" json:"id"`
	UserID     uint64     `json:"userId"`
	VoucherID  uint64     `json:"voucherId"`
	PayType    uint8      `gorm:"default:1" json:"payType"`
	Status     uint8      `gorm:"default:1" json:"status"`
	CreateTime time.Time  `gorm:"->" json:"createTime"`
	PayTime    *time.Time `json:"payTime"`
	UseTime    *time.Time `json:"useTime"`
	RefundTime *time.Time `json:"refundTime"`
	UpdateTime time.Time  `gorm:"->" json:"updateTime"`
}

func (TbVoucherOrder) TableName() string {
	return "tb_voucher_order"
}
