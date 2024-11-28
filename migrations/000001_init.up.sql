create extension if not exists "uuid-ossp";

create table "user" (
    id UUID PRIMARY KEY default uuid_generate_v4()
);

create table token (
    user_id UUID references "user" (id) not null unique,
    refresh_token TEXT unique not null,
    exp TIMESTAMP not null
);
