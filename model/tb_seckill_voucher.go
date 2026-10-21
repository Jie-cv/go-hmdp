package model

import "time"

// CREATE TABLE `tb_seckill_voucher`  (
//   `voucher_id` bigint(20) UNSIGNED NOT NULL COMMENT '关联的优惠券的id',
//   `stock` int(8) NOT NULL COMMENT '库存',
//   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//   `begin_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生效时间',
//   `end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '失效时间',
//   `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//   PRIMARY KEY (`voucher_id`)
// ) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '秒杀优惠券表，与优惠券是一对一关系' ROW_FORMAT = Compact;

type TbSeckillVoucher struct {
	VoucherID  uint64    `gorm:"primaryKey;autoIncrement:false" json:"voucher_id"`
	Stock      int32     `json:"stock"`
	CreateTime time.Time `gorm:"->" json:"create_time"`
	BeginTime  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"begin_time"`
	EndTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"end_time"`
	UpdateTime time.Time `gorm:"->" json:"update_time"`
}

func (TbSeckillVoucher) TableName() string {
	return "tb_seckill_voucher"
}
