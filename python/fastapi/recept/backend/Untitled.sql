CREATE TABLE IF NOT EXISTS "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "password" varchar,
  "person_id" varchar
);
CREATE INDEX idx_users_id ON "users" ("id");

CREATE TABLE IF NOT EXISTS "person" (
  "id" varchar PRIMARY KEY,
  "name" varchar,
  "birth" varchar,
  "sex" varchar,
  "profile" varchar,
  "phone_number" varchar
);
CREATE INDEX idx_person_id ON "person" ("id");

CREATE TABLE IF NOT EXISTS "receipt" (
  "id" varchar PRIMARY KEY,
  "user_id" varchar,
  "username" varchar,
  "title" varchar,
  "ingridient" text,
  "intructions" text,
  "notes" text,
  "like" bigint,
  "comment" varchar,
  "status" varchar,
  "created_at" bigint,
  "updated_at" bigint
);
CREATE INDEX idx_receipt_id ON "receipt" ("id");
CREATE INDEX idx_receipt_title ON "receipt" ("title");


ALTER TABLE "receipt" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "users" ADD FOREIGN KEY ("person_id") REFERENCES "person" ("id");
