-- +goose Up
-- +goose StatementBegin
CREATE TABLE people (
                        id SERIAL PRIMARY KEY,
                        first_name VARCHAR(100) NOT NULL,
                        last_name VARCHAR(100) NOT NULL,
                        gender VARCHAR(10),
                        nationality VARCHAR(10),
                        age INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS people;
-- +goose StatementEnd
