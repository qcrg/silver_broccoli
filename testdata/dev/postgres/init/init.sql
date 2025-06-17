CREATE TABLE users (
  id BIGINT PRIMARY KEY
);
COPY users FROM '/csv_import/users.csv' DELIMITER ',' CSV HEADER;

CREATE TABLE user_extra_privileges(
  id BIGINT PRIMARY KEY REFERENCES users(id),
  privileges BIT(9) NOT NULL
);
COPY user_extra_privileges FROM '/csv_import/user_extra_privileges.csv' DELIMITER ',' CSV HEADER;

CREATE TABLE wallet_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR(16) NOT NULL
);
COPY wallet_types FROM '/csv_import/wallet_types.csv' DELIMITER ',' CSV HEADER;
ALTER SEQUENCE wallet_types_id_seq RESTART WITH 1000;

CREATE TABLE wallets (
  id BIGSERIAL PRIMARY KEY,
  type_id INT NOT NULL REFERENCES wallet_types(id),
  amount BIGINT NOT NULL DEFAULT 0,
  frozen BOOL NOT NULL DEFAULT FALSE
);
COPY wallets FROM '/csv_import/wallets.csv' DELIMITER ',' CSV HEADER;
ALTER SEQUENCE wallets_id_seq RESTART WITH 1000;

CREATE TABLE wallet_acls(
  -- id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id),
  wallet_id BIGINT NOT NULL REFERENCES wallets(id),
  access_rights BIT(4) NOT NULL,
  PRIMARY KEY(user_id, wallet_id)
);
COPY wallet_acls FROM '/csv_import/wallet_acls.csv' DELIMITER ',' CSV HEADER;
-- ALTER SEQUENCE wallet_acls_id_seq RESTART WITH 1000;

CREATE TYPE fee_types AS ENUM ('fixed', 'fraction');

CREATE TABLE fees(
  id SERIAL PRIMARY KEY,
  type fee_types NOT NULL,
  wallet_type_id INT NOT NULL REFERENCES wallet_types(id) UNIQUE
);
COPY fees FROM '/csv_import/fees.csv' DELIMITER ',' CSV HEADER;
ALTER SEQUENCE fees_id_seq RESTART WITH 1000;

CREATE TABLE fixed_fees(
  id INT PRIMARY KEY REFERENCES fees(id),
  value INT NOT NULL DEFAULT 0
);
COPY fixed_fees FROM '/csv_import/fixed_fees.csv' DELIMITER ',' CSV HEADER;

CREATE TABLE fraction_fees(
  id INT PRIMARY KEY REFERENCES fees(id),
  value FLOAT NOT NULL DEFAULT 0
);
COPY fraction_fees FROM '/csv_import/fraction_fees.csv' DELIMITER ',' CSV HEADER;

CREATE TYPE transaction_statuses
AS ENUM ('created', 'reserved', 'completed', 'failed', 'reversed');

CREATE TABLE transactions(
  id BIGSERIAL PRIMARY KEY,
  wallet_src_id BIGINT ,
  wallet_dst_id BIGINT,
  status transaction_statuses NOT NULL,
  amount BIGINT NOT NULL,
  CONSTRAINT at_least_one_not_null CHECK(
    (wallet_src_id IS NOT NULL) OR (wallet_dst_id IS NOT NULL)
  )
);
