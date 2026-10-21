-- 业务模块：voucher 初始数据
-- 先执行 schema 下的建表 SQL，再导入本文件；仅需导入一次。
SET NAMES utf8mb4;

START TRANSACTION;

-- 表：tb_voucher（1 条初始记录）

INSERT INTO `tb_voucher` VALUES (1, 1, '50元代金券', '周一至周日均可使用', '全场通用\\n无需预约\\n可无限叠加\\不兑现、不找零\\n仅限堂食', 4750, 5000, 0, 1, '2022-01-04 09:42:39', '2022-01-04 09:43:31');

-- 表：tb_seckill_voucher（0 条初始记录）

-- 原始 init.sql 未提供该表的数据。

-- 表：tb_voucher_order（0 条初始记录）

-- 原始 init.sql 未提供该表的数据。

COMMIT;
