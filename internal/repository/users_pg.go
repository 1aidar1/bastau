package repository

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/pkg/logger"
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type UserRepo struct {
	db     *pgxpool.Pool
	logger logger.LoggerI
}

const users_table = "users"

var _ UsersRepoI = UserRepo{}

func NewUsersRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u UserRepo) GetById(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, name FROM ` + users_table + ` WHERE id = $1`
	var user entity.User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := u.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (u UserRepo) Register(ctx context.Context, user entity.User) error {
	q := `INSERT INTO ` + users_table +
		`(name,email,phone,password,is_male, country_id,city_id,date_of_birth,created_at,last_visit_at) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	args := []interface{}{
		user.Name,
		user.Email,
		user.Phone,
		user.Password,
		user.IsMale,
		user.CountryId,
		user.CityId,
		user.DateOfBirth,
		time.Now(),
		time.Now(),
	}
	_, err := u.db.Exec(ctx, q, args...)
	if err != nil {
		u.logger.Warn(err)
		return err
	}
	return nil
}

func (u UserRepo) GetByCredentials(ctx context.Context, email, password string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Verify(ctx context.Context, userID int, code string) error {
	//TODO implement me
	panic("implement me")
}
