CREATE DATABASE IF NOT EXISTS hello_cafe COLLATE utf8_general_ci;

CREATE TABLE `admin` (
     `admin_seq` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'PK ',
     `phone` varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '핸드폰번호',
     `password` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '비밀번호',
     `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '이름',
     `reg_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '등록일',
     `mod_dt` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '수정일',
     PRIMARY KEY (`admin_seq`),
     UNIQUE KEY `phone` (`phone`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `logout_token` (
    `token_seq` bigint(20) NOT NULL AUTO_INCREMENT,
    `admin_seq` bigint(20) NOT NULL COMMENT 'admin sequence',
    `token` varchar(255) CHARACTER SET utf8mb4 NOT NULL COMMENT 'token',
    PRIMARY KEY (`token_seq`),
    UNIQUE KEY `admin_seq_token` (`admin_seq`,`token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `item` (
    `item_seq` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `admin_seq` bigint(20) NOT NULL COMMENT 'admin sequence',
    `category` tinyint(4) NOT NULL COMMENT '카테고리(0:음료, 1:음식)',
    `barcode` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '바코드',
    `price` bigint(20) NOT NULL COMMENT '가격',
    `cost` bigint(20) NOT NULL COMMENT '원가',
    `name` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '이름',
    `consonant` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '이름 초성',
    `description` text CHARACTER SET utf8mb4 COMMENT '설명',
    `expire_dt` datetime NOT NULL COMMENT '유통기한',
    `size` tinyint(4) NOT NULL COMMENT '사이즈(0:small, 1:large)',
    `reg_dt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '등록일',
    `mod_dt` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '수정일',
    PRIMARY KEY (`item_seq`),
    UNIQUE KEY `barcode` (`barcode`) USING BTREE,
    KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;