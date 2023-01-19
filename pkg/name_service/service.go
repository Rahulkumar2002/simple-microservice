package name_service

import (
	"context"
)

type Service interface {
	// Get the name of the user :
	GiveName(ctx context.Context, name string) (string, error)
}
