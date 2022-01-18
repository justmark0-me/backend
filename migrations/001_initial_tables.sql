create table if not exists request_info (
    i_d SERIAL PRIMARY KEY,
    i_p varchar(15) NOT NULL,
    os varchar(128) NOT NULL,
    country varchar(52) NOT NULL,
    region varchar(128) not null,
    city varchar(32) NOT NULL,
    latitude float NOT NULL,
    longitude float NOT NULL,
    user_agent varchar(128) NULL,
    endpoint varchar(128) NULL,
    additional_info text NOT NULL,
    timestamp timestamp NOT NULL
);
