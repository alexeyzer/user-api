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
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users;
-- +goose StatementEnd
