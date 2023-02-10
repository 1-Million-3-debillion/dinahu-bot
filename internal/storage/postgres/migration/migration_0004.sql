CREATE TABLE IF NOT EXISTS "stats" (
    "id" UUID NOT NULL UNIQUE,
    "user_id" BIGINT NOT NULL,
    "chat_id" BIGINT NOT NULL,
    "dinahu_count" BIGINT,
    CONSTRAINT "stats_user_chat_id" UNIQUE ("user_id", "chat_id")
);

CREATE INDEX IF NOT EXISTS "stats_index"
ON "user_chat" ("chat_id");