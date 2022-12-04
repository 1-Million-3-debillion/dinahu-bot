CREATE TABLE IF NOT EXISTS "chat" (
    "chat_id" BIGINT NOT NULL UNIQUE,
    "name" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL
);