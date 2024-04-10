package data

import (
	"context"
	"database/sql"
	"time"
)

/*
CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    found_time timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    found_by bigint NOT NULL REFERENCES users ON DELETE CASCADE,
    embedding text,
    location text,
	category text NOT NULL,
    claimed bool NOT NULL default false,
    claimed_by text,
	image text,
    version integer NOT NULL DEFAULT 1
);

*/

type Item struct {
	ID        int64  `json:"id"`
	FoundTime string `json:"found_time"`
	FoundBy   int64  `json:"found_by"`
	Location  string `json:"location"`
	Category  string `json:"category"`
	Claimed   bool   `json:"claimed"`
	ClaimedBy string `json:"claimed_by"`
	Image     string `json:"image"`
	Version   int    `json:"-"`
}

type ItemModel struct {
	DB *sql.DB
}

func (m ItemModel) Insert(item *Item) error {
	query := `
		INSERT INTO item (found_by, location, category, image)
		VALUES ($1, $2, $3, $4)
		RETURNING id, found_time, found_by, location, category, version`

	args := []interface{}{item.FoundBy, item.Location, item.Category, item.Image}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&item.ID, &item.FoundTime, &item.FoundBy, &item.Location, &item.Category, &item.Version)
	return err
}

func (m ItemModel) GetAll() ([]Item, error) {
	query := `
		SELECT id, found_time, found_by, location, category, claimed, claimed_by, image
		FROM item`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.FoundTime, &item.FoundBy, &item.Location, &item.Category, &item.Claimed, &item.ClaimedBy, &item.Image)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (m ItemModel) GetAllUnclaimed() ([]Item, error) {
	query := `
		SELECT id, found_time, found_by, location, category, image
		FROM item 
		WHERE claimed = false`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.FoundTime, &item.FoundBy, &item.Location, &item.Category, &item.Image)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (m ItemModel) GetAllItems() ([]Item, error) {
	query := `
		SELECT id, found_time, found_by, location, category, image, claimed, COALESCE(claimed_by, '') as claimed_by
		FROM item`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.ID, &item.FoundTime, &item.FoundBy, &item.Location, &item.Category, &item.Image, &item.Claimed, &item.ClaimedBy)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (m ItemModel) Claim(id int64, claimed_by string) error {
	query := `
		UPDATE item
		SET claimed = true, claimed_by = $1
		WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, claimed_by, id)
	return err
}
