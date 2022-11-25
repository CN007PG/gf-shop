CREATE TABLE `shop`.`item_cat` (
                                   `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
                                   `parent_id` BIGINT(20) NULL COMMENT '父类目id',
                                   `name` VARCHAR(50) NULL COMMENT '类目名称',
                                   `status` TINYINT(4) NULL DEFAULT '1' COMMENT '状态 1-正常销售，2-停售',
                                   `sort_order` INT(4) NULL COMMENT '排序号',
                                   `is_parent` TINYINT(1) NULL DEFAULT '1' COMMENT '是否为父类目',
                                   `create_at` DATETIME NOT NULL COMMENT '创建时间',
                                   `update_at` DATETIME NOT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
    ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COMMENT = '商品类目表';
