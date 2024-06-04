-- +goose Up
CREATE TABLE devices (
    id  INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id  UUID NOT NULL,
    description VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    FOREIGN KEY(user_id) 
        REFERENCES users(id)
);
-- +goose Down
DROP TABLE sessions;