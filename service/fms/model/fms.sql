CREATE SCHEMA `fms`;

CREATE TABLE `fms`.`billbook` (
  `id` varchar(64) UNIQUE,
  `name` varchar(16) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);

CREATE TABLE `fms`.`category` (
  `id` varchar(64) UNIQUE,
  `parent_id` varchar(64),
  `name` varchar(16) NOT NULL,
  `level` int NOT NULL,
  `billbook_id` varchar(64) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);

CREATE TABLE `fms`.`asset_category` (
  `id` varchar(64) UNIQUE,
  `name` varchar(16) NOT NULL,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);

CREATE TABLE `fms`.`asset` (
  `id` varchar(64) UNIQUE,
  `name` varchar(255) NOT NULL,
  `asset_cate_id` varchar(64) NOT NULL,
  `hide` boolean,
  `count_into` boolean,
  `balance` int NOT NULL,
  `unit` varchar(255) NOT NULL DEFAULT "FEN",
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);

CREATE TABLE `fms`.`bill` (
  `id` varchar(64) UNIQUE,
  `category_id` varchar(64) NOT NULL,
  `billbook_id` varchar(64) NOT NULL,
  `asset_out_id` varchar(64),
  `asset_in_id` varchar(64),
  `debt_id` varchar(64),
  `bill_type` varchar(255) NOT NULL,
  `amount` int NOT NULL,
  `unit` varchar(255) NOT NULL DEFAULT "FEN",
  `remark` varchar(255),
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);

CREATE TABLE `fms`.`debt` (
  `id` varchar(64) UNIQUE,
  `debt_type` varchar(255) NOT NULL,
  `object` varchar(255) NOT NULL,
  `amount` int NOT NULL,
  `repay_amount` int NOT NULL,
  `unit` varchar(255) NOT NULL DEFAULT "FEN",
  `pay_date` datetime,
  `repay_date` datetime,
  `count_into` boolean,
  `ended` boolean,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime,
  `is_deleted` boolean DEFAULT false,
  PRIMARY KEY ( `id` )
);