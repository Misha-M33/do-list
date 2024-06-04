package database

import (
	"context"
	"do-list/internal/entities"
	"fmt"
)

func (q *DB) TableUsers(u *entities.User) error {
	sql_statement := `CREATE TABLE users (
		id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NULL,
		full_name VARCHAR(255) NULL,
		nickname VARCHAR(255) NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		is_deleted bool NULL,
		is_block bool NULL,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_DATE,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_DATE
	);`
	_, err := q.db.Exec(context.Background(), sql_statement)
	if err != nil {
		return err
	}

	return nil
}

func (q *DB) CreateUser(password string, u *entities.User) error {
	sql_statement := `INSERT INTO users (id,first_name,last_name,full_name,
		nickname,email,password,is_deleted,is_block,created_at,updated_at) 
		VALUES (gen_random_uuid(),$1,$2,$3,$4,$5,$6,false,false,now(),now())`
	_, err := q.db.Exec(context.Background(), sql_statement, u.First_name, u.Last_name, u.Full_name, u.Nickname, u.Email, password)
	if err != nil {
		return err
	}

	return nil
}

func (q *DB) GetUserByEmail(user *entities.User) (string, string, bool, bool, error) {
	sql_statement := `SELECT id, password,is_deleted, is_block FROM users WHERE email=$1`
	row := q.db.QueryRow(context.Background(), sql_statement, user.Email)

	var u entities.User
	err := row.Scan(&u.Id, &u.Password, &u.Is_deleted, &u.Is_block)
	if err != nil {
		fmt.Println("User not found:", err)
	}

	return u.Id.String(), string(u.Password), u.Is_deleted, u.Is_block, nil
}

func (q *DB) UsersGroup(ug *entities.UserGroup, users *entities.Users) error {
	sql_statement := `SELECT user_id FROM users_groups WHERE group_id=
					(SELECT group_id FROM users_groups WHERE user_id=$1) `

	rows, err := q.db.Query(context.Background(), sql_statement, ug.User_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var u entities.User
	for rows.Next() {
		err := rows.Scan(&u.Id)
		if err != nil {
			return err
		}

		*users = append(*users, u)
	}

	return nil
}

func (q *DB) DeleteUser(u *entities.User) (int64, error) {
	sql_statement := `UPDATE users SET is_deleted = $1, updated_at = now() WHERE id = $2`
	row, err := q.db.Query(context.Background(), sql_statement, u.Is_deleted, u.Id)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	err = row.Scan(&u.Is_deleted, &u.Id)

	result := row.CommandTag().RowsAffected()

	return result, nil
}

func (q *DB) BlockUser(u *entities.User) (int64, error) {
	sql_statement := `UPDATE users SET is_block = $1, updated_at = now() WHERE id = $2`
	row, err := q.db.Exec(context.Background(), sql_statement, u.Is_block, u.Id)
	if err != nil {
		return 0, err
	}

	result := row.RowsAffected()

	return result, nil
}

func (q *DB) EditPassword(passHash string, user *entities.User) error {
	sql_statement := `UPDATE users SET password = $1, updated_at = now() WHERE id = $2`
	row, err := q.db.Query(context.Background(), sql_statement, passHash, user.Id)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}

func (q *DB) GetUsers(users *entities.Users) (int64, error) {
	sql_statement := `SELECT id, first_name, last_name,full_name,nickname, 
		email, is_deleted, is_block, created_at, updated_at FROM users`
	rows, err := q.db.Query(context.Background(), sql_statement)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var u entities.User
	for rows.Next() {
		err := rows.Scan(&u.Id, &u.First_name, &u.Last_name, &u.Full_name, &u.Nickname, &u.Email, &u.Is_deleted, &u.Is_block, &u.Created_at, &u.Updated_at)
		if err != nil {
			return 0, err
		}

		*users = append(*users, u)
	}
	result := rows.CommandTag().RowsAffected()

	return result, nil
}

func (q *DB) UpdateUser(u *entities.User) error {
	sql_statement := `UPDATE users 
		SET first_name=$1,last_name=$2,full_name=$3,nickname=$4,updated_at=now() 
		WHERE id = $5`
	row, err := q.db.Query(context.Background(), sql_statement, u.First_name, u.Last_name, u.Full_name, u.Nickname, u.Id)
	if err != nil {
		return err
	}
	defer row.Close()

	return nil
}
