CREATE TABLE `tenants`
(
    `id`             int(11) NOT NULL AUTO_INCREMENT,
    `name`           varchar(255) NOT NULL COMMENT '租户名称',
    `parent_id`      int(11) NOT NULL DEFAULT -1 COMMENT '父级id',
    `model`          varchar(10)  NOT NULL COMMENT '模型：aliyun',
    `is_application` tinyint(4) NOT NULL DEFAULT 0 COMMENT '可构成应用',
    `access_key`     varchar(255) NOT NULL COMMENT '访问密钥',
    `secret_key`     varchar(255) NOT NULL COMMENT '签名密钥',
    `dashboard`      varchar(255) NOT NULL COMMENT '密钥管理面板',
    `desc`           varchar(255) NOT NULL COMMENT '描述',
    `created_at`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;