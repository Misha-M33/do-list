package database

import (
	"context"
	"do-list/internal/entities"
)

func (q *DB) CreateGroup(g *entities.Group) error {
	sql_statement := `INSERT INTO groups (name,owner_id) VALUES ($1,$2)`
	_, err := q.db.Exec(context.Background(), sql_statement, g.Name, g.Owner_id)
	if err != nil {
		return err
	}

	sql_stat := `SELECT id,name,owner_id FROM groups WHERE name=$1`
	row := q.db.QueryRow(context.Background(), sql_stat, g.Name)
	if err := row.Scan(&g.Id, &g.Name, &g.Owner_id); err != nil {
		return err
	}

	return nil
}

func (q *DB) UpdateGroup(g *entities.Group) error {
	sql_statement := `UPDATE groups SET name=$1, owner_id=$2 WHERE id=$3`
	_, err := q.db.Exec(context.Background(), sql_statement, g.Name, g.Owner_id, g.Id)
	if err != nil {
		return err
	}

	return nil
}

func (q *DB) DeleteGroup(g *entities.Group) (int64, error) {
	sql_statement := `DELETE FROM groups WHERE owner_id=$1 AND name=$2`
	res, err := q.db.Exec(context.Background(), sql_statement, g.Owner_id, g.Name)
	if err != nil {
		return 0, err
	}

	result := res.RowsAffected()

	return result, nil
}

func (q *DB) GetGroups(groups *entities.Groups) (int64, error) {
	sql_statement := `SELECT id, name, owner_id FROM groups`
	rows, err := q.db.Query(context.Background(), sql_statement)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var g entities.Group
	for rows.Next() {
		err := rows.Scan(&g.Id, &g.Name, &g.Owner_id)
		if err != nil {
			return 0, err
		}
		*groups = append(*groups, g)
	}
	result := rows.CommandTag().RowsAffected()

	return result, nil
}
