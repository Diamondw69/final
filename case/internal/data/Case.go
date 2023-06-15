package data

import (
	"case/pkg/proto"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type CaseModel struct {
	DB *sql.DB
}

func (m CaseModel) InsertItem(item *proto.Case) error {

	itemsJSON, err := json.Marshal(item.CaseItems)
	if err != nil {
		return err
	}

	query := `
INSERT INTO cases(name, price, items)
VALUES ($1, $2, $3::jsonb)
RETURNING id`
	args := []any{item.Name, item.Price, itemsJSON}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = m.DB.QueryRowContext(ctx, query, args...).Scan(&item.Id)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m CaseModel) GetCaseID(id int64) (*proto.Case, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
SELECT id, name, price, items 
FROM cases 
WHERE id = $1`
	var itemsJSON string
	var item proto.Case
	err := m.DB.QueryRow(query, id).Scan(
		&item.Id,
		&item.Name,
		&item.Price,
		&itemsJSON,
	)
	err = json.Unmarshal([]byte(itemsJSON), &item.CaseItems)
	if err != nil {
		return nil, err
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &item, nil
}

func (m CaseModel) GetAllCase() (*proto.Cases, error) {
	query := `
SELECT *
FROM cases`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var itemsJSON string
	items := proto.Cases{Cases: []*proto.Case{}}
	for rows.Next() {
		var item proto.Case
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Price,
			&itemsJSON,
		)
		err = json.Unmarshal([]byte(itemsJSON), &item.CaseItems)

		if err != nil {
			return nil, err // Update this to return an empty Metadata struct.
		}
		items.Cases = append(items.Cases, &item)
	}
	if err = rows.Err(); err != nil {
		return nil, err // Update this to return an empty Metadata struct.
	}

	return &items, nil
}

func (m CaseModel) DeleteCase(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
DELETE FROM cases
WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
