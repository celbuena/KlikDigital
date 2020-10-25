CREATE SCHEMA users;

CREATE TABLE users.user(
    id SERIAL NOT NULL,
    nama character varying (100) NOT NULL ,
    phone_number character varying (100) NOT NULL ,
    created_date timestamp without time zone ,
    password character varying (100) NOT NULL ,
    uuid character varying (100) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY(id)
);


CREATE SCHEMA transactions;

CREATE TABLE transactions.transactions(
    id SERIAL NOT NULL,
    user_id integer ,
    nomor_transaction character varying (20) NOT NULL ,
    created_date timestamp without time zone,

    CONSTRAINT transactions_pkey PRIMARY KEY(id)
)