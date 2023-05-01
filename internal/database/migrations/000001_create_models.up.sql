CREATE TABLE IF NOT EXISTS entain.users
(
    id bigserial NOT NULL,
    balance decimal NOT NULL DEFAULT 0,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS entain.requests
(
    id bigserial NOT NULL,
    user_id    bigint  NOT NULL,
    transaction_id text COLLATE pg_catalog."default" NOT NULL,
    state text COLLATE pg_catalog."default" NOT NULL,
    amount decimal NOT NULL,
    processed boolean NOT NULL DEFAULT FALSE,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT requests_pkey PRIMARY KEY (id),
    CONSTRAINT unique_entaint_transaction_id_user_id UNIQUE (user_id, transaction_id),
    CONSTRAINT fk_entain_user_id FOREIGN KEY (user_id)
    REFERENCES entain.users (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- add default user here
INSERT INTO entain.users(id, balance) values(1, 0)
