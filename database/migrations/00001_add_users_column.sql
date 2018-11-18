-- +goose Up
CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(100),
    phone_number bigint,
    country_code varchar(5),
    updated_at TIMESTAMP NULL,
    created_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;