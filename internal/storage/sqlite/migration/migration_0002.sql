CREATE TABLE IF NOT EXISTS "chat" (
    "chat_id" INTEGER NOT NULL UNIQUE,
    "name" TEXT,
    "created_at" INTEGER
);