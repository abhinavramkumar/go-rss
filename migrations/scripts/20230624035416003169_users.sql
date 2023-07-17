CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE OR REPLACE FUNCTION generate_ulid() RETURNS uuid
    AS $$
        SELECT (lpad(to_hex(floor(extract(epoch FROM clock_timestamp()) * 1000)::bigint), 12, '0') || encode(gen_random_bytes(10), 'hex'))::uuid;
    $$ LANGUAGE SQL;

SELECT generate_ulid();

create table users (
  user_id UUID UNIQUE NOT NULL DEFAULT generate_ulid(),
  name varchar(255) NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  created_at timestamp DEFAULT current_timestamp,
  updated_at timestamp DEFAULT current_timestamp
)