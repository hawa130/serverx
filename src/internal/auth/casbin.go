package auth

import (
	"context"
	"fmt"

	"github.com/hawa130/computility-cloud/internal/perm"
)

func EnforceCtx(ctx context.Context, obj string, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforce(user.ID.String(), obj, act)
}

func EnforceXCtx(ctx context.Context, obj fmt.Stringer, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.EnforceX(user.ID, obj, act)
}

func IsAdmin(ctx context.Context) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
}