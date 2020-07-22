-- ----------------------------
-- Table structure for admins Model
-- ----------------------------
--
create table admins
(
    id           int unsigned auto_increment primary key COMMENT '自增ID',
    create_at    timestamp                                                       null COMMENT '创建时间',
    update_at    timestamp                                                       null COMMENT '更新时间',
    delete_at    timestamp                                                       null COMMENT '删除时间',
    uuid         varchar(255)                                                    null COMMENT '用户唯一标识UUID',
    nickname     varchar(255) default 'QMPlusUser'                               null COMMENT '用户昵称',
    header_img   varchar(255) default 'http://www.henrongyi.top/avatar/lufu.jpg' null COMMENT '用户头像',
    authority_id varchar(255) default '888'                                      null COMMENT '用户角色ID',
    username     varchar(255)                                                    null COMMENT '用户名',
    password     varchar(255)                                                    null COMMENT '用户登录密码'
)
    charset = utf8;

create index idx_admins_deleted_at
    on admins (delete_at);

-- ----------------------------
-- Table structure for jwt_blacklists Model
-- ----------------------------
DROP TABLE IF EXISTS `jwts`;
CREATE TABLE `jwts`
(
    `id`        int UNSIGNED                                    NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `create_at` timestamp                                       NULL DEFAULT NULL COMMENT '更新时间',
    `update_at` timestamp                                       NULL DEFAULT NULL COMMENT '更新时间',
    `delete_at` timestamp                                       NULL DEFAULT NULL COMMENT '删除时间',
    `jwt`       text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT 'jwt',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_jwts_deleted_at` (`delete_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 57
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;


-- ----------------------------
-- Table structure for apis Model
-- ----------------------------
DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis`
(
    `id`          int(10) UNSIGNED                                        NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `create_at`   timestamp                                               NULL DEFAULT NULL COMMENT '创建时间',
    `update_at`   timestamp                                               NULL DEFAULT NULL COMMENT '更新时间',
    `delete_at`   timestamp                                               NULL DEFAULT NULL COMMENT '删除时间',
    `path`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api路径',
    `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
    `api_group`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api组',
    `method`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'POST' COMMENT '方法',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_apis_delete_at` (`delete_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 91
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for casbin_rule Model
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`
(
    `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '类型',
    `v0`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色id',
    `v1`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求路由',
    `v2`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '请求方法',
    `v3`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v4`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
    `v5`     varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for authorities Model
-- ----------------------------
DROP TABLE IF EXISTS `authorities`;
CREATE TABLE `authorities`
(
    `authority_id`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '角色ID',
    `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色名',
    `parent_id`      varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '父角色ID',
    `create_at`      datetime(0)                                             NULL DEFAULT NULL COMMENT '创建时间',
    `update_at`      datetime(0)                                             NULL DEFAULT NULL COMMENT '更新时间',
    `delete_at`      datetime(0)                                             NULL DEFAULT NULL COMMENT '删除时间',
    `resources_id`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '资源权限ID',
    PRIMARY KEY (`authority_id`) USING BTREE,
    UNIQUE INDEX `authority_id` (`authority_id`) USING BTREE,
    INDEX `idx_sys_authorities_deleted_at` (`delete_at`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus Model
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`
(
    `id`           int(10) UNSIGNED                                        NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`   timestamp(0)                                            NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`   timestamp(0)                                            NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`   timestamp(0)                                            NULL DEFAULT NULL COMMENT '删除时间',
    `menu_level`   int(10) UNSIGNED                                        NULL DEFAULT NULL COMMENT '菜单等级(预留字段)',
    `parent_id`    int(10) UNSIGNED                                        NULL DEFAULT NULL COMMENT '父菜单ID',
    `path`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由path',
    `name`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '路由name',
    `hidden`       tinyint(1)                                              NULL DEFAULT NULL COMMENT '是否在列表隐藏',
    `component`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '前端文件路径',
    `title`        varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名',
    `icon`         varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
    `sort`         int(255)                                                NULL DEFAULT NULL COMMENT '排序标记',
    `keep_alive`   tinyint(1)                                              NULL DEFAULT NULL COMMENT '是否缓存',
    `default_menu` tinyint(1)                                              NULL DEFAULT NULL COMMENT '是否是基础路由(开发中)',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_menus_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 52
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;


-- ----------------------------
-- Table structure for dictionaries Model
-- ----------------------------
DROP TABLE IF EXISTS `dictionaries`;
CREATE TABLE `sys_dictionaries`
(
    `id`         int(10) UNSIGNED                                        NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at` datetime(0)                                             NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime(0)                                             NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime(0)                                             NULL DEFAULT NULL COMMENT '删除时间',
    `name`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
    `type`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
    `status`     tinyint(1)                                              NULL DEFAULT NULL COMMENT '状态',
    `desc`       varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '描述',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_sys_dictionaries_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 8
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for dictionary_details Model
-- ----------------------------
DROP TABLE IF EXISTS `dictionary_details`;
CREATE TABLE `dictionary_details`
(
    `id`                int(10) UNSIGNED                                        NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`        datetime(0)                                             NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`        datetime(0)                                             NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`        datetime(0)                                             NULL DEFAULT NULL COMMENT '删除时间',
    `label`             varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '展示值',
    `value`             int(11)                                                 NULL DEFAULT NULL COMMENT '字典值',
    `status`            tinyint(1)                                              NULL DEFAULT NULL COMMENT '启用状态',
    `sort`              int(11)                                                 NULL DEFAULT NULL COMMENT '排序标记',
    `dictionary_id` int(11)                                                 NULL DEFAULT NULL COMMENT '关联标记',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_sys_dictionary_details_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 38
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Compact;


-- ----------------------------
-- Table structure for files Model
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `created_at` timestamp(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp(0) NULL DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;


-- ----------------------------
-- Table structure for customers Model
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers`
(
    `id`                    int(10) UNSIGNED                                        NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`            timestamp(0)                                            NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`            timestamp(0)                                            NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`            timestamp(0)                                            NULL DEFAULT NULL COMMENT '删除时间',
    `customer_name`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NULL DEFAULT NULL COMMENT '客户名',
    `customer_phone_data`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '客户电话',
    `sys_user_id`           int(10) UNSIGNED                                        NULL DEFAULT NULL COMMENT '负责员工id',
    `sys_user_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '负责员工角色',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_exa_customers_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for operations Model
-- ----------------------------
DROP TABLE IF EXISTS `operations`;
CREATE TABLE `operations`
(
    `id`            int(10) UNSIGNED                                       NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`    datetime(0)                                            NULL DEFAULT NULL COMMENT '创建时间',
    `updated_at`    datetime(0)                                            NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`    datetime(0)                                            NULL DEFAULT NULL COMMENT '删除时间',
    `ip`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求ip',
    `method`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求方法',
    `path`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求路由',
    `status`        int(11)                                                NULL DEFAULT NULL COMMENT '状态',
    `latency`       bigint(20)                                             NULL DEFAULT NULL COMMENT '延迟',
    `agent`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '代理',
    `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '报错信息',
    `request`          text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin         NULL COMMENT '请求Body',
    `user_id`       int(11)                                                NULL DEFAULT NULL COMMENT '用户id',
    `response`          text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin         NULL COMMENT '响应Body',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_operations_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 342
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_bin
  ROW_FORMAT = Compact;

