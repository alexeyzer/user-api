-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_role(
    id SERIAL,
    user_id bigint not null,
    role_id bigint not null,
    CONSTRAINT role_user_unique UNIQUE (user_id, role_id)
)


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_role
-- +goose StatementEnd
