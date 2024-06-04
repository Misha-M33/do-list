package database

import "context"

func (q *DB) SaveRefreshToken(id string, token string) error {
	sql_delete := `DELETE FROM tokens WHERE user_id=$1`
	_, err := q.db.Exec(context.Background(), sql_delete, id)
	if err != nil {
		return err
	}

	sql_statement := `INSERT INTO tokens(user_id,token) VALUES($1,$2)`
	_, err = q.db.Exec(context.Background(), sql_statement, id, token)
	if err != nil {
		return err
	}

	return nil
}
