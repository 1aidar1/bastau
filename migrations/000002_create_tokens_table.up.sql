DROP TYPE IF EXISTS tokens_scope;
CREATE TYPE tokens_scope AS ENUM ('auth', 'activation');

CREATE TABLE IF NOT EXISTS tokens(
     hash bytea PRIMARY KEY,
     user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
     expiry timestamp(0) with time zone NOT NULL,
     scope tokens_scope NOT NULL
);