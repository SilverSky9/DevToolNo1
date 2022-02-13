-- -------------------------------------------------------------
-- TablePlus 4.5.2(402)
--
-- https://tableplus.com/
--
-- Database: Market
-- Generation Time: 2022-02-14 12:59:04.8240 AM
-- -------------------------------------------------------------


-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS interested_post_post_id_seq;
CREATE SEQUENCE IF NOT EXISTS interested_post_user_id_seq;

-- Table Definition
CREATE TABLE "public"."interested_post" (
    "post_id" int4 NOT NULL DEFAULT nextval('interested_post_post_id_seq'::regclass),
    "user_id" int4 NOT NULL DEFAULT nextval('interested_post_user_id_seq'::regclass)
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS pin_pin_id_seq;

-- Table Definition
CREATE TABLE "public"."pin" (
    "user_id" int4,
    "pin_id" int4 NOT NULL DEFAULT nextval('pin_pin_id_seq'::regclass),
    "pin" varchar(255),
    PRIMARY KEY ("pin_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS post_post_id_seq;
DROP TYPE IF EXISTS "public"."option_product";
CREATE TYPE "public"."option_product" AS ENUM ('buy', 'sell');

-- Table Definition
CREATE TABLE "public"."post" (
    "post_id" int4 NOT NULL DEFAULT nextval('post_post_id_seq'::regclass),
    "pin_id" int4,
    "product_name" varchar(255),
    "post_date" date,
    "product_option" "public"."option_product",
    "price" int4,
    "amount" int4,
    "tag_id" int4,
    PRIMARY KEY ("post_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS tag_tag_id_seq;

-- Table Definition
CREATE TABLE "public"."tag" (
    "tag_id" int4 NOT NULL DEFAULT nextval('tag_tag_id_seq'::regclass),
    "tag_name" varchar,
    PRIMARY KEY ("tag_id")
);

-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS user1_user_id_seq;

-- Table Definition
CREATE TABLE "public"."user1" (
    "user_id" int4 NOT NULL DEFAULT nextval('user1_user_id_seq'::regclass),
    "name" varchar(255),
    "address" varchar(255),
    "contact" varchar(255),
    "phone_number" varchar(10),
    PRIMARY KEY ("user_id")
);

INSERT INTO "public"."pin" ("user_id", "pin_id", "pin") VALUES
(1, 1, '$2a$14$0AcCPReUHV7jj70qrydRrOKBg1C87M0.BczJmzFaN4tsnlMmNj3hK'),
(2, 2, '$2a$14$BZ6KsYwiHOS/W62KG1UEveQNvubOXWlL8DFXKb1xlQ2X4goZky6wm'),
(3, 3, '$2a$14$8n.qzdlZqBonJK3.GCbZnuHywcQLNeT8pT9.0erPcTxJm0aYe2WFS'),
(4, 4, '$2a$14$4n9.3hsKlqpCFG584af4ruCbmQcJsTWXYiASgS/MZMoGFixGHUCG6');

INSERT INTO "public"."post" ("post_id", "pin_id", "product_name", "post_date", "product_option", "price", "amount", "tag_id") VALUES
(1, 1, 'ขายเร้าเตอร์ต่อครับ', '2022-02-13', 'sell', 400, 1, 1),
(2, 2, 'อยากขายน้ำลำไย', '2022-02-13', 'sell', 20, 10, 6),
(3, 3, 'ขายต่อกรงกระต่าย', '2022-02-13', 'sell', 150, 1, 7),
(4, 4, 'ขายต่อเก้าอี้คอม', '2022-02-13', 'sell', 900, 1, 8);

INSERT INTO "public"."tag" ("tag_id", "tag_name") VALUES
(1, 'it'),
(2, 'food'),
(3, 'food'),
(4, 'food'),
(5, 'food'),
(6, 'food'),
(7, 'pet'),
(8, 'com');

INSERT INTO "public"."user1" ("user_id", "name", "address", "contact", "phone_number") VALUES
(1, 'mark', 'jinda', 'markza', '0926894589'),
(2, 'milk', 'rnp', 'milktest', '12123'),
(3, 'fah', 'k1', 'fahfah1234', '0926897975'),
(4, 'mark', 'jinda', 'markza', '0926894589');

ALTER TABLE "public"."interested_post" ADD FOREIGN KEY ("post_id") REFERENCES "public"."post"("post_id");
ALTER TABLE "public"."interested_post" ADD FOREIGN KEY ("user_id") REFERENCES "public"."user1"("user_id");
ALTER TABLE "public"."pin" ADD FOREIGN KEY ("user_id") REFERENCES "public"."user1"("user_id");
ALTER TABLE "public"."post" ADD FOREIGN KEY ("pin_id") REFERENCES "public"."pin"("pin_id");
