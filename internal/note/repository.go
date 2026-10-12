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
	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func (repo *Repository) FindNoteByID(
	ctx context.Context,
	id int64,
) (*Note, error) {
	var note Note
	row := repo.db.QueryRowContext(
		ctx,
		"SELECT id, title, content FROM notes WHERE id = ?",
		id,
	)

	err := row.Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &note, nil
}

func (repo *Repository) CreateNote(
	ctx context.Context,
	req CreateNoteRequest,
) (int64, error) {

	res, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO notes (title, content) VALUES (?, ?)",
		req.Title,
		req.Content,
	)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}
