package authz

import (
	"context"
)

const (
	UserKey string = "user"
)

func GetLoggedInUser(ctx context.Context) interface{} {
	return ctx.Value(UserKey)
}
