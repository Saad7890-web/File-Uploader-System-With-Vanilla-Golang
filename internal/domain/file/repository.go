package file

import "context"

type Repository interface {
	Save(ctx context.Context, file *File) error
	FindByID(ctx context.Context, id string) (*File, error)
	Delete(ctx context.Context, id string) error
}
