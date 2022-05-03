-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart(
                          id SERIAL,
                          user_id bigint not null,
                          final_product_id bigint not null,
                          quantity bigint not null
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart;
-- +goose StatementEnd
