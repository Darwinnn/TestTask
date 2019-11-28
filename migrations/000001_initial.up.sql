CREATE TABLE IF NOT EXISTS balances (
    id SERIAL PRIMARY KEY,
    value double precision NOT NULL,
    CONSTRAINT balances_value_check CHECK ((value >= (0.0)::double precision))
);

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    uuid uuid UNIQUE NOT NULL,
    balance_id integer NOT NULL,
    amount double precision NOT NULL,
    canceled boolean DEFAULT false NOT NULL,
    state text
);

