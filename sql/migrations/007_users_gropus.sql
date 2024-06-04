-- +goose Up
CREATE TABLE users_groups (
    id  INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id UUID NOT NULL,
    group_id  int NOT NULL,
    FOREIGN KEY(user_id) 
        REFERENCES users(id),
    FOREIGN KEY(group_id) 
        REFERENCES groups(id)
);
-- +goose Down
DROP TABLE users_groups;