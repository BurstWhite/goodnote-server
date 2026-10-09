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

func (serv *Service) GetNotesByID(
	ctx context.Context,
	id int,
) (*Note, error) {
	note, err := serv.repo.FindNotesByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return note, err
}
