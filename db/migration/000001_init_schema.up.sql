-- Create the accounts table
CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create the entries table
CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,  -- Assuming a default of 0 instead of a timestamp
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create the transfers table
CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

-- Create indexes for improved query performance
CREATE INDEX "idx_accounts_owner" ON "accounts" ("owner");
CREATE INDEX "idx_entries_account_id" ON "entries" ("account_id");
CREATE INDEX "idx_transfers_from_account_id" ON "transfers" ("from_account_id");
CREATE INDEX "idx_transfers_to_account_id" ON "transfers" ("to_account_id");
CREATE INDEX "idx_transfers_from_to" ON "transfers" ("from_account_id", "to_account_id");

-- Add comments to clarify column usage
COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';
COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';

-- Add foreign key constraints to maintain referential integrity
ALTER TABLE "entries"
  ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
  ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers"
  ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id");
