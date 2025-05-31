CREATE TABLE "accounts" (
  "id_account" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" varchar NOT NULL,
  "currency" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id_entries" bigserial PRIMARY KEY,
  "id_account" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id_transfer" bigserial PRIMARY KEY,
  "from_id_account" bigint NOT NULL,
  "to_id_account" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("id_account");

CREATE INDEX ON "transfers" ("from_id_account");

CREATE INDEX ON "transfers" ("to_id_account");

CREATE INDEX ON "transfers" ("from_id_account", "to_id_account");

ALTER TABLE "entries" ADD FOREIGN KEY ("id_account") REFERENCES "accounts" ("id_account");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_id_account") REFERENCES "accounts" ("id_account");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_id_account") REFERENCES "accounts" ("id_account");
