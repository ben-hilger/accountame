CREATE TABLE users (
    id serial,
    name varchar(200),
    email varchar(200),
    password_hash varchar(500),
    PRIMARY KEY(id)
);