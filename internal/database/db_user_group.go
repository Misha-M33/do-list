package database

import (
	"context"
	"do-list/internal/entities"
)

func (q *DB) GetUsersGroups(usersGroups *entities.UsersGroups) (int64, error) {
	sql_statement := `SELECT id, user_id, group_id FROM users_groups`
	rows, err := q.db.Query(context.Background(), sql_statement)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var ug entities.UserGroup
	for rows.Next() {
		err := rows.Scan(&ug.Id, &ug.User_id, &ug.Group_id)
		if err != nil {
			return 0, err
		}
		*usersGroups = append(*usersGroups, ug)
	}
	result := rows.CommandTag().RowsAffected()

	return result, nil
}

func (q *DB) CreateUserGroup(g *entities.UserGroup) error {
	sql_statement := `INSERT INTO users_groups (user_id,group_id) VALUES ($1,$2)`
	_, err := q.db.Exec(context.Background(), sql_statement, g.User_id, g.Group_id)
	if err != nil {
		return err
	}

	return nil
}

func (q *DB) DeleteUserGroup(ug *entities.UserGroup) (int64, error) {
	sql_statement := `DELETE FROM users_groups WHERE user_id=$1 AND id=$2`
	res, err := q.db.Exec(context.Background(), sql_statement, ug.User_id, ug.Id)
	if err != nil {
		return 0, err
	}

	result := res.RowsAffected()

	return result, nil
}
