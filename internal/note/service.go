package note

import "context"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) GetNotes(
	ctx context.Context,
) ([]Note, error) {
	notes, err := serv.repo.FindNotes(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (serv *Service) GetNoteByID(
	ctx context.Context,
	id int64,
) (*Note, error) {
	note, err := serv.repo.FindNoteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, err
}

func checkNoteLegality(req CreateNoteRequest) error {
	return nil // not implemented here
}

func (serv *Service) CreateNote(
	ctx context.Context,
	req CreateNoteRequest,
) (int64, error) {
	if err := checkNoteLegality(req); err != nil {
		return -1, err
	}

	return serv.repo.CreateNote(ctx, req)
}
