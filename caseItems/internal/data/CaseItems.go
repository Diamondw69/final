package data

import (
	pb "ahah/pkg/proto"
	"context"
	"database/sql"
	"errors"
	"time"
)

type CaseItemModel struct {
	DB *sql.DB
}

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

func (m CaseItemModel) InsertItem(item *pb.CaseItem) error {
	query := `
INSERT INTO caseitems(itemname, itemdesc, type, stars,image)
VALUES ($1, $2, $3, $4,$5)
RETURNING id`
	args := []any{item.ItemName, item.ItemDescription, item.Type, item.Stars, item.Image}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&item.Id)
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

func (m CaseItemModel) GetCaseItem(id int64) (*pb.CaseItem, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
SELECT id, itemname, itemdesc, type,stars,image
FROM caseitems
WHERE id = $1`

	var item pb.CaseItem
	err := m.DB.QueryRow(query, id).Scan(
		&item.Id,
		&item.ItemName,
		&item.ItemDescription,
		&item.Type,
		&item.Stars,
		&item.Image,
	)

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

func (m CaseItemModel) DeleteCaseItems(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
DELETE FROM caseitems
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

func (m CaseItemModel) UpdateItem(item *pb.CaseItem) error {
	query := `
UPDATE caseitems
SET itemname = $1, itemdesc = $2, type = $3, stars = $4, image=$6
WHERE id = $5
RETURNING id`

	args := []interface{}{
		item.ItemName,
		item.ItemDescription,
		item.Type,
		item.Stars,
		item.Id,
		item.Image,
	}

	return m.DB.QueryRow(query, args...).Scan(&item.Id)
}

func (m CaseItemModel) GetCaseItemByName(name string) (*pb.CaseItem, error) {

	if name == "" {
		return nil, ErrRecordNotFound
	}

	query := `
SELECT id, itemname, itemdesc, type,stars,image
FROM caseitems
WHERE itemname = $1`

	var item pb.CaseItem
	err := m.DB.QueryRow(query, name).Scan(
		&item.Id,
		&item.ItemName,
		&item.ItemDescription,
		&item.Type,
		&item.Stars,
		&item.Image,
	)

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

func (m CaseItemModel) GetAllCaseItems() (*pb.CaseItems, error) {
	query := `
SELECT *
FROM caseitems`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err // Update this to return an empty Metadata struct.
	}
	defer rows.Close()
	items := pb.CaseItems{}
	for rows.Next() {
		var item pb.CaseItem
		err := rows.Scan(
			&item.Id,
			&item.ItemName,
			&item.ItemDescription,
			&item.Type,
			&item.Stars,
			&item.Image,
		)
		if err != nil {
			return nil, err // Update this to return an empty Metadata struct.
		}
		items.CaseItems = append(items.CaseItems, &item)
	}
	if err = rows.Err(); err != nil {
		return nil, err // Update this to return an empty Metadata struct.
	}
	return &items, nil
}
