package note

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) FindNotes(
	ctx context.Context,
) ([]Note, error) {
	notes := make([]Note, 0)

	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT id, title, content FROM notes",
	)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}

func (repo *Repository) FindNotesByID(
	ctx context.Context,
	id int,
) (*Note, error) {
	var note Note
	row := repo.db.QueryRowContext(
		ctx,
		"SELECT id, title, content FROM notes WHERE id = ?",
		id,
	)

	err := row.Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		return nil, err
	}

	return &note, nil
}
