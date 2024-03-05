CREATE TABLE users (
    id serial primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);