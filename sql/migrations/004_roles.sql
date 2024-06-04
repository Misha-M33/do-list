-- +goose Up
CREATE TABLE roles (
    id  INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    restrictions VARCHAR[] NOT NULL,
    user_id  UUID NOT NULL,
     FOREIGN KEY(user_id) 
        REFERENCES users(id)
);
-- +goose Down
DROP TABLE roles;
