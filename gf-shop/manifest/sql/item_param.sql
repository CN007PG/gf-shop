CREATE TABLE `shop`.`item_param` (
                                    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
                                    `item_cat_id` BIGINT(20) NULL COMMENT '商品类目',
                                    `param_data` JSON NULL COMMENT '商品json参数',
                                    `create_at` DATETIME NOT NULL COMMENT '创建时间',
                                    `update_at` DATETIME NOT NULL COMMENT '更新时间',
                                    PRIMARY KEY (`id`),
                                    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
    ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '商品参数';