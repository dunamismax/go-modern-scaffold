-- Create "messages" table
CREATE TABLE "messages" (
  "id" bigserial NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);
