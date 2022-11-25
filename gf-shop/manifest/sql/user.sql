CREATE TABLE `shop`.`user` (
                               `id` BIGINT NOT NULL AUTO_INCREMENT,
                               `name` VARCHAR(20) CHARACTER SET 'utf8mb4' NOT NULL COMMENT '用户名',
                               `password` VARCHAR(20) NOT NULL COMMENT '密码',
                               `phone` VARCHAR(20) NOT NULL COMMENT '电话',
                               `buy_history` JSON NULL DEFAULT NULL COMMENT '购买记录',
                               `buy_cart` JSON NULL DEFAULT NULL COMMENT '购物车',
                               `create_at` DATETIME NOT NULL COMMENT '创建时间',
                               `update_at` DATETIME NOT NULL COMMENT '更新时间',
                               UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE,
                               PRIMARY KEY (`id`))
    ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '用户表';