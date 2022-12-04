CREATE TABLE IF NOT EXISTS "user" (
    "user_id" BIGINT UNIQUE NOT NULL,
    "first_name" TEXT,
    "last_name" TEXT,
    "username" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL
);