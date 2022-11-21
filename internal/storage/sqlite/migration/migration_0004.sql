CREATE TABLE IF NOT EXISTS "stats" (
    "id" TEXT NOT NULL UNIQUE,
    "user_id" INTEGER NOT NULL,
    "chat_id" INTEGER NOT NULL,
    "dinahu_count" INTEGER,
    CONSTRAINT "user_chat_id" UNIQUE ("user_id", "chat_id")
);

CREATE INDEX IF NOT EXISTS "stats_index"
ON "user_chat" ("chat_id");