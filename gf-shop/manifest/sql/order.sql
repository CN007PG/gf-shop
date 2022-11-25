CREATE TABLE `shop`.`order` (
                                `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
                                `order_no` VARCHAR(20) NULL COMMENT '订单号',
                                `order_unique` VARCHAR(50) NULL COMMENT '流水号',
                                `money` INT(11) NOT NULL COMMENT '金额',
                                `status` TINYINT(4) NULL DEFAULT '1' COMMENT '状态',
                                `create_at` DATETIME NOT NULL COMMENT '创建时间',
                                `update_at` DATETIME NOT NULL COMMENT '更新时间',
                                PRIMARY KEY (`id`),
                                UNIQUE INDEX `orderNo_UNIQUE` (`orderNo` ASC) VISIBLE)
    ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '订单';