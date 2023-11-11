package bases

import "context"

type BaseRepository interface {
	Create(ctx context.Context, data interface{}) error
	Find(ctx context.Context, id int, entity interface{}) error
	FindAll(ctx context.Context, entities interface{}) error
	Update(ctx context.Context, id int, data interface{}) error
	Delete(ctx context.Context, id int) error
}
