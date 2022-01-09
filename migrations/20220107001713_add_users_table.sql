-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id SERIAL,
    name varchar(100) not null,
    surname varchar(100) not null,
    patronymic varchar(100),
    password bytea not null,
    phone varchar(100),
    email varchar(200) not null,
    username varchar(200) not null
);

create unique INDEX users_username_uqidx on users(username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index users_username_uqidx;

DROP TABLE users;
-- +goose StatementEnd
