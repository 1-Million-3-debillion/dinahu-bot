CREATE TABLE IF NOT EXISTS "user" (
    "user_id" INTEGER UNIQUE NOT NULL,
    "first_name" TEXT,
    "last_name" TEXT,
    "username" TEXT
);