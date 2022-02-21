package repository

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository/postgres"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
	"time"
)

type UserRepo struct {
	db *pgxpool.Pool
	//logger *logger.LoggerI
	dbc postgres.Querier
}

const users_table = "users"

var _ UsersRepoI = &UserRepo{}

func NewUsersRepo(db *pgxpool.Pool, sqlc postgres.Querier) *UserRepo {
	return &UserRepo{
		//logger: l,
		db:  db,
		dbc: sqlc,
	}
}

func (u *UserRepo) GetById(ctx context.Context, id int) (entity.User, error) {
	row, err := u.dbc.GetUserById(ctx, int64(id))
	user := entity.User{
		ID:    int(row.ID),
		Name:  row.Name,
		Email: row.Email,
		Role:  string(row.Role),
		Phone: row.Phone,
	}

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return entity.EmptyUser, ErrRecordNotFound
		default:
			return entity.EmptyUser, err
		}
	}
	return user, nil
}

func (u *UserRepo) Register(ctx context.Context, user entity.User) error {
	args := postgres.RegisterUserParams{
		Name:     user.Name,
		Email:    strings.ToLower(user.Email),
		Phone:    user.Phone,
		Role:     postgres.Role(user.Role),
		Password: user.Password,
	}
	_, err := u.dbc.RegisterUser(ctx, args)
	fmt.Println(err)
	if err != nil {
		switch {
		case err.Error() == pgerrcode.DuplicateColumn:
			return ErrRecordNotFound
		default:
			return err
		}
	}
	return nil
}

func (u *UserRepo) GetByToken(ctx context.Context, hash []byte) (entity.User, error) {
	args := postgres.GetUserByTokenParams{
		Hash:   hash,
		Scope:  postgres.TokensScopeAuth,
		Expiry: time.Now(),
	}
	userRow, err := u.dbc.GetUserByToken(ctx, args)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return entity.EmptyUser, ErrRecordNotFound
		default:
			return entity.EmptyUser, err
		}
	}
	return entity.User{
		ID:    int(userRow.ID),
		Name:  userRow.Name,
		Email: userRow.Email,
		Phone: userRow.Phone,
		Role:  string(userRow.Role),
	}, nil
}

func (u *UserRepo) GetByCredentials(ctx context.Context, email, password string) (entity.User, error) {
	args := postgres.GetUserByCredentialsParams{
		Email:    email,
		Password: []byte(password),
	}
	user, err := u.dbc.GetUserByCredentials(ctx, args)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return entity.EmptyUser, ErrRecordNotFound
		default:
			return entity.EmptyUser, err
		}
	}

	return entity.User{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Verify(ctx context.Context, userID int, code string) error {
	//TODO implement me
	panic("implement me")
}
