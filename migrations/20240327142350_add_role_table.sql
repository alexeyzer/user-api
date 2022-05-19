-- +goose Up
-- +goose StatementBegin
CREATE TABLE role(
                      id SERIAL,
                      name varchar(100) not null,
                      description varchar(300)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP Table role;
-- +goose StatementEnd
