CREATE TABLE "servers" (
  "id" bigserial PRIMARY KEY,
  "Hostname" varchar NOT NULL,
  "Ip" bigint NOT NULL,
  "Version" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
