package common

import "context"

type Container interface {
	ConnectionString(ctx context.Context) (string, error)
}
