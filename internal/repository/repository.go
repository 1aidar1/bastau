package repository

import (
	"1aidar1/bastau/go-api/internal/entity"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UsersRepoI interface {
	GetById(ctx context.Context, id int) (*entity.User, error)
	Register(ctx context.Context, user entity.User) error
	GetByCredentials(ctx context.Context, email, password string) (entity.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error)
	Verify(ctx context.Context, userID int, code string) error
	//SetSession(ctx context.Context, userID int, session entity.Session) error
}

type Repositories struct {
	Users UsersRepoI
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Users: NewUsersRepo(db),
	}
}
