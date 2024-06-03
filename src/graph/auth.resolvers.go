package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/user"
	"github.com/hawa130/computility-cloud/graph/model"
	"github.com/hawa130/computility-cloud/graph/reqerr"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/database"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.LoginPayload, error) {
	c := ent.FromContext(ctx)
	record, err := c.User.Query().Where(user.PhoneEQ(input.Phone)).Only(database.WrapAllowContext(ctx))
	if ent.IsNotFound(err) {
		if _, err := auth.HashPassword(input.Password); err != nil {
			return nil, err
		}
		return nil, reqerr.ErrInvalidLoginInput
	}
	if err != nil {
		return nil, err
	}

	ok, err := auth.ComparePasswordAndHash(input.Password, record.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, reqerr.ErrInvalidLoginInput
	}

	token, err := auth.GenerateToken(record.ID)
	if err != nil {
		return nil, err
	}

	return &model.LoginPayload{
		Token: &token,
		User:  record,
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*ent.User, error) {
	c := ent.FromContext(ctx)
	exist, err := c.User.Query().Where(user.PhoneEQ(input.Phone)).Exist(database.WrapAllowContext(ctx))
	if exist {
		return nil, reqerr.ErrPhoneAlreadyExist
	}
	if err != nil {
		return nil, err
	}

	return c.User.Create().
		SetPhone(input.Phone).
		SetPassword(input.Password).
		Save(database.WrapAllowContext(ctx))
}
