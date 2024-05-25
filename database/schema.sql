CREATE TABLE users (
    id varchar(36),
    name varchar(200),
    email varchar(200),
    username varchar(200),
    password_hash varchar(500),
    PRIMARY KEY(id)
);