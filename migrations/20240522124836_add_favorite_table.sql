-- +goose Up
-- +goose StatementBegin
CREATE TABLE favorite(
                       id SERIAL,
                       user_id bigint not null,
                       product_id bigint not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE favorite;
-- +goose StatementEnd
