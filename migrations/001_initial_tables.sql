create table if not exists ip_request (
    id SERIAL PRIMARY KEY,
    ip varchar(15) NOT NULL,
    os varchar(20) NOT NULL,
    country varchar(32) NOT NULL,
    city varchar(32) NOT NULL,
    geolocation varchar(128) NOT NULL
);
