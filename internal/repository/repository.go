package repository

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository/postgres"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersRepoI interface {
	GetById(ctx context.Context, id int) (entity.User, error)
	GetByToken(ctx context.Context, hash []byte) (entity.User, error)
	Register(ctx context.Context, user entity.User) error
	GetByCredentials(ctx context.Context, email, password string) (entity.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error)
	Verify(ctx context.Context, userID int, code string) error
	//SetSession(ctx context.Context, userID int, session entity.Session) error
}

type TokenRepoI interface {
	Create(ctx context.Context, token entity.Token) error
	DeleteAllForUser(ctx context.Context, userId int) error
	DeleteScopedForUser(ctx context.Context, userId int, scope string) error
}

type Repositories struct {
	Users  UsersRepoI
	Tokens TokenRepoI
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	sqlc := postgres.New(db)
	return &Repositories{
		Users:  NewUsersRepo(db, sqlc),
		Tokens: NewTokensRepo(db, sqlc),
	}
}
