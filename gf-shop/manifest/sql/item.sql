CREATE TABLE `shop`.`item` (
                               `id` BIGINT(20) NOT NULL COMMENT '商品id,商品编号',
                               `title` VARCHAR(50) NOT NULL COMMENT '商品标题',
                               `sell_point` VARCHAR(200) NULL COMMENT '商品卖点',
                               `price` BIGINT(20) NOT NULL COMMENT '商品价格单位元',
                               `num` INT(10) NOT NULL COMMENT '库存数量',
                               `barcode` VARCHAR(30) NULL COMMENT '商品条形码',
                               `image` VARCHAR(200) NULL COMMENT '商品图片',
                               `cid` BIGINT(10) NOT NULL COMMENT '商品类目',
                               `status` TINYINT(4) NOT NULL DEFAULT '1' COMMENT '商品状态 1-正常销售，2-停售',
                               `create_at` DATETIME NOT NULL COMMENT '创建时间',
                               `update_at` DATETIME NOT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`),
                               UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
    ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '商品表';