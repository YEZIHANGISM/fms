\c "fms";

CREATE TABLE "billbook" (
  "id" bigint UNIQUE PRIMARY KEY,
  "name" varchar(16) NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);

CREATE TABLE "category" (
  "id" bigint UNIQUE PRIMARY KEY,
  "parent_id" bigint,
  "name" varchar(16) NOT NULL,
  "level" int NOT NULL,
  "billbook_id" bigint NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);

CREATE TABLE "asset_category" (
  "id" bigint UNIQUE PRIMARY KEY,
  "name" varchar(16) NOT NULL,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);

CREATE TABLE "asset" (
  "id" bigint UNIQUE PRIMARY KEY,
  "name" varchar NOT NULL,
  "asset_cate_id" bigint NOT NULL,
  "hide" boolean,
  "count_into" boolean,
  "balance" int NOT NULL,
  "unit" varchar NOT NULL DEFAULT 'FEN',
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);

CREATE TABLE "bill_type" (
  "id" bigint UNIQUE PRIMARY KEY,
  "bill_type" varchar NOT NULL,
  "asset_out_id" bigint,
  "asset_in_id" bigint
);

CREATE TABLE "bill" (
  "id" bigint UNIQUE PRIMARY KEY,
  "billbook_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "bill_type_id" bigint NOT NULL,
  "debt_id" bigint,
  "amount" decimal(10,2),
  "remark" varchar,
  "bill_date" date DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);

CREATE TABLE "debt" (
  "id" bigint UNIQUE PRIMARY KEY,
  "debt_type" varchar NOT NULL,
  "object" varchar NOT NULL,
  "count_into" boolean,
  "ended" boolean,
  "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp,
  "is_deleted" boolean DEFAULT false
);