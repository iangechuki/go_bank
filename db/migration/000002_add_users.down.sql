-- ↓↓↓ Down Migration ↓↓↓

BEGIN;

-- 1. Drop the owner–currency UNIQUE constraint on accounts
ALTER TABLE IF EXISTS accounts
  DROP CONSTRAINT IF EXISTS owner_currency_key;

-- 2. Drop the foreign-key from accounts.owner → users.username
--    (Postgres default name: accounts_owner_fkey; adjust if yours differs)
ALTER TABLE IF EXISTS accounts
  DROP CONSTRAINT IF EXISTS accounts_owner_fkey;

-- 3. Drop the standalone unique index on (owner,currency), if you created one
--    (Default would be something like accounts_owner_currency_idx)
DROP INDEX IF EXISTS accounts_owner_currency_idx;

-- 4. Drop the users table
DROP TABLE IF EXISTS users;

COMMIT;
