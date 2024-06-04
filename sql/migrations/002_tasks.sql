-- +goose Up
CREATE TABLE tasks (
    id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title VARCHAR(255) NOT NULL,
    description text NOT NULL,
    responsible UUID NULL,
    priority INT NOT NULL,
    is_done bool DEFAULT false NOT NULL,
    creator uuid NOT NULL,
    group_id INT NULL,
    deadline_date VARCHAR(60) NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_DATE,
    CONSTRAINT check_priority check(priority in(1, 2, 3, 4)),
    FOREIGN KEY(creator) REFERENCES users(id),
    FOREIGN KEY(responsible) REFERENCES users(id),
    FOREIGN KEY(group_id) REFERENCES groups(id)
);
-- +goose Down
DROP TABLE tasks;