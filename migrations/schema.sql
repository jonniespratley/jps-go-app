CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "users" (
"id" TEXT PRIMARY KEY,
"name" TEXT NOT NULL,
"email" TEXT NOT NULL,
"bio" text,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "themes" (
"id" TEXT PRIMARY KEY,
"uri" TEXT NOT NULL,
"base_uri" TEXT NOT NULL,
"main" TEXT NOT NULL,
"error" TEXT NOT NULL,
"error_chromeless" TEXT NOT NULL,
"options" text,
"display_name" TEXT NOT NULL,
"demo_link" text,
"description" text,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
