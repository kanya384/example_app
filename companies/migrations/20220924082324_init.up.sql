CREATE TYPE "employee_role" AS ENUM (
  'administrator',
  'deliveryman'
);

CREATE TABLE "company" (
  "id" uuid PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "inn" bigint NOT NULL,
  "address" varchar(300) NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

CREATE TABLE "project" (
  "id" uuid PRIMARY KEY,
  "name" varchar(150) NOT NULL,
  "description" varchar(400) NOT NULL,
  "address" varchar(300) NOT NULL,
  "longitude" float NOT NULL,
  "latitude" float NOT NULL,
  "company_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

CREATE TABLE "employee" (
  "id" uuid PRIMARY KEY,
  "role" employee_role,
  "user_id" uuid UNIQUE NOT NULL,
  "project_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);

COMMENT ON TABLE "company" IS 'table "company" conatins companies information';

COMMENT ON TABLE "project" IS 'table "project" conatins projects of the companies';

COMMENT ON TABLE "employee" IS 'table "employee" contains employess of the projects';

COMMENT ON COLUMN "employee"."user_id" IS 'users id form users service';

ALTER TABLE "project" ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;

ALTER TABLE "employee" ADD FOREIGN KEY ("project_id") REFERENCES "project" ("id") ON DELETE CASCADE ON UPDATE NO ACTION;