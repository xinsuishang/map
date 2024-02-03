CREATE TABLE `tenants`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `name`       varchar(255) NOT NULL COMMENT '租户名称',
    `region`     tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否有region区分',
    `type`       varchar(10)  NOT NULL COMMENT '资源类型：qiniu',
    `access_key` varchar(255) NOT NULL COMMENT '访问密钥',
    `secret_key` varchar(255) NOT NULL COMMENT '签名密钥',
    `dashboard`  varchar(255) NOT NULL COMMENT '密钥管理面板',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

CREATE TABLE `domain_mapping`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `tenant_id`   int(11) NOT NULL COMMENT '租户',
    `region_id`   varchar(10)  NOT NULL DEFAULT '' COMMENT 'regionId',
    `domain`      varchar(255) NOT NULL DEFAULT '' COMMENT '域名',
    `bucket_name` varchar(50)  NOT NULL COMMENT '存储空间名称',
    `desc`        varchar(255) NOT NULL COMMENT '存储空间名称',
    `created_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;