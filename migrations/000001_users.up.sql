-- CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
DROP TYPE IF EXISTS role;
CREATE TYPE role AS ENUM ('worker', 'client', 'all');

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    email text NOT NULL UNIQUE,
    role role NOT NULL,
    phone text NOT NULL UNIQUE,
    password bytea NOT NULL
-- --     activated bool NOT NULL,
--
--     last_visit_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
--     updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
--     created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
);

-- CREATE INDEX IF NOT EXISTS user_email_idx ON users USING GIN (email);