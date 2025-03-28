-- +goose Up
-- +goose StatementBegin
CREATE TABLE friendships (
                             id SERIAL PRIMARY KEY,
                             user_id INT REFERENCES people(id),
                             friend_id INT REFERENCES people(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS friendships;
-- +goose StatementEnd
