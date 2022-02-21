// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"context"
)

type Querier interface {
	CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error)
	DeleteAllTokensForUser(ctx context.Context, userID int64) (Token, error)
	DeleteScopeTokenForUser(ctx context.Context, arg DeleteScopeTokenForUserParams) (Token, error)
	DeleteUser(ctx context.Context, id int64) (User, error)
	GetUserByCredentials(ctx context.Context, arg GetUserByCredentialsParams) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	GetUserByToken(ctx context.Context, arg GetUserByTokenParams) (GetUserByTokenRow, error)
	RegisterUser(ctx context.Context, arg RegisterUserParams) (int64, error)
}

var _ Querier = (*Queries)(nil)
