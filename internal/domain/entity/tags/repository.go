package tags

import "context"

type TagsRepository interface {
	GetAll(ctx context.Context) ([]Tag, error)
}
