CREATE TABLE "mail" (
  "id" uuid PRIMARY KEY,
  "recipient" varchar(300) NOT NULL,
  "subject" varchar(78) NOT NULL,
  "message" text NOT NULL,
  "status" smallint NOT NULL,
  "created_at" timestamp NOT NULL,
  "modified_at" timestamp NOT NULL
);