CREATE TYPE "user_role" AS ENUM (
  'administrator',
  'deliveryman',
  'root'
);

CREATE TYPE "device_types" AS ENUM (
  'web',
  'ios',
  'android'
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "surname" varchar(100) NOT NULL,
  "phone" varchar(100) NOT NULL,
  "pass" varchar(300) NOT NULL,
  "email" varchar(300) NOT NULL,
  "role" user_role NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

CREATE TABLE "device" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "device_id" varchar(150),
  "ip" varchar(15) NOT NULL,
  "agent" varchar(150) NOT NULL,
  "type" device_types NOT NULL,
  "refresh_token" varchar(200) NOT NULL,
  "refresh_exp" timestamp NOT NULL,
  "created_at" timestamp NOT NULL,
  "last_seen" timestamp NOT NULL
);

COMMENT ON TABLE "user" IS 'table "user" conatins users of service';

COMMENT ON COLUMN "device"."device_id" IS 'if provided (for android and ios users)';

COMMENT ON COLUMN "device"."agent" IS 'browser info or mobile device info';

COMMENT ON COLUMN "device"."refresh_exp" IS 'after this time token expires';

ALTER TABLE "device" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;