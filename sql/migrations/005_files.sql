-- +goose Up
CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    file_name VARCHAR(255) NOT NULL,
    file_url VARCHAR(511) NOT NULL,
    access_type int NOT NULL
		CONSTRAINT check_access_type check(access_type in(1,2,3))
);
-- +goose Down
DROP TABLE files;