-- +goose Up
CREATE TABLE groups (
    id  INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(255) NOT NULL,
    owner_id UUID NOT NULL,
        FOREIGN KEY(owner_id) 
            REFERENCES users(id)
);
-- +goose Down
DROP TABLE groups;

