-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_status AS ENUM ('CREATED', 'DECLINED', 'DONE');

CREATE TABLE orders(
                       id SERIAL,
                       user_id bigint not null,
                       order_status order_status not null,
                       order_date timestamp not null,
                       total_cost double precision not null,
                       items jsonb not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;

DROP TYPE order_status;
-- +goose StatementEnd
