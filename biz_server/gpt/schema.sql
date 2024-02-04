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

CREATE TABLE `chat`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `model_id`   int(11) NOT NULL COMMENT '模型id',
    `session_id` varchar(64)  NOT NULL COMMENT '会话',
    `name`       varchar(255) NOT NULL COMMENT '描述',
    `created_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_session` (`session_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `chat_message`
(
    `id`         int(11) NOT NULL AUTO_INCREMENT,
    `chat_id`    int(11) NOT NULL COMMENT '模型id',
    `request_id` varchar(64) NOT NULL COMMENT '请求标识',
    `text`       text        NOT NULL COMMENT '内容',
    `version`    int(11) NOT NULL COMMENT '版本，0为prompt',
    `created_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY          `idx_chat_request` (`chat_id`,`request_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;