CREATE TABLE users
(
    id            serial       not null unique,
    firstname     varchar(255) not null,
    lastname      varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          int          not null
);