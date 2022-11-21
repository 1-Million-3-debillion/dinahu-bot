CREATE TABLE IF NOT EXISTS "user_chat" (
    "id" TEXT NOT NULL UNIQUE,
    "user_id" INTEGER NOT NULL,
    "chat_id" INTEGER NOT NULL,
    CONSTRAINT "user_chat_id" UNIQUE ("user_id", "chat_id")
);

CREATE INDEX IF NOT EXISTS "chat_id_index"
ON "user_chat" ("chat_id");