-- +goose Up
-- Create "messages" table
CREATE TABLE "messages" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "body" TEXT NOT NULL,
  "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- Drop "messages" table
DROP TABLE "messages";