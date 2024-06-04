package database

import (
	"context"
	"do-list/internal/entities"
)

func (pg *DB) CreateTask(task *entities.Task) error {
	sql_statement := `INSERT INTO tasks (title,description,responsible,priority,is_done,creator,group_id,deadline_date,created_at)
			VALUES ($1,$2,$3,$4,false,$5,$6,$7,now())`
	_, err := pg.db.Exec(context.Background(), sql_statement, task.Title, task.Description, task.Responsible, task.Priority, task.Creator, task.Group_id, task.Deadline_date)
	if err != nil {
		return err
	}

	return nil
}

func (pg *DB) GetTasks(task *entities.Tasks) error {
	sql_statement := `SELECT id,title,description,responsible,priority,is_done,creator,group_id,deadline_date,created_at FROM tasks`
	rows, err := pg.db.Query(context.Background(), sql_statement)
	if err != nil {
		return err
	}
	defer rows.Close()

	var t entities.Task
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Responsible, &t.Priority, &t.Is_done, &t.Creator, &t.Group_id, &t.Deadline_date, &t.Created_at)
		if err != nil {
			return err
		}

		*task = append(*task, t)
	}

	return nil
}

func (pg *DB) GetTaskByGroup(t *entities.Task, tasks *entities.Tasks) error {
	sql_statement := `SELECT id,title,description,responsible,priority,is_done,creator,deadline_date FROM tasks WHERE group_id=$1`
	rows, err := pg.db.Query(context.Background(), sql_statement, t.Group_id)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Responsible, &t.Priority, &t.Is_done, &t.Creator, &t.Deadline_date)
		if err != nil {
			return err
		}

		*tasks = append(*tasks, *t)
	}

	return nil
}

func (pg *DB) TasksByUserId(u *entities.User, tasks *entities.Tasks) error {
	sql_statement := `SELECT tasks.id,title,description,responsible,priority,is_done,creator,group_id,deadline_date FROM tasks JOIN groups 
		on groups.id = tasks.group_id  AND groups.id=(select users_groups.group_id 
		from users_groups where users_groups.user_id=$1)`
	rows, err := pg.db.Query(context.Background(), sql_statement, u.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	var t entities.Task
	for rows.Next() {
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Responsible, &t.Priority, &t.Is_done, &t.Creator, &t.Group_id, &t.Deadline_date)
		if err != nil {
			return err
		}

		*tasks = append(*tasks, t)
	}

	return nil
}

func (pg *DB) DeleteTask(t *entities.Task) error {
	sql_statement := `DELETE FROM tasks WHERE id = $1`
	_, err := pg.db.Exec(context.Background(), sql_statement, t.Id)
	if err != nil {
		return err
	}

	return nil
}

func (pg *DB) UpdateTask(t *entities.Task) error {
	sql_statement := `UPDATE tasks SET title=$1, description=$2, responsible=$3, 
					priority=$4, is_done=$5, group_id=$6, deadline_date=$7 WHERE id = $8`
	_, err := pg.db.Exec(context.Background(), sql_statement, t.Title, t.Description,
		t.Responsible, t.Priority, t.Is_done, t.Group_id, t.Deadline_date, t.Id)
	if err != nil {
		return err
	}

	return nil
}
