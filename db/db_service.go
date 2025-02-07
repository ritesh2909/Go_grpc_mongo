package db

import "context"

type DBService interface {
	Connect(ctx context.Context, uri string) (interface{}, error)
	Disconnect(ctx context.Context) error
}
