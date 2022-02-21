package repository

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository/postgres"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TokenRepo struct {
	db *pgxpool.Pool
	//logger *logger.LoggerI
	dbc postgres.Querier
}

const tokens_table = "tokens"

var _ TokenRepoI = &TokenRepo{}

func NewTokensRepo(db *pgxpool.Pool, sqlc postgres.Querier) *TokenRepo {
	return &TokenRepo{
		//logger: l,
		db:  db,
		dbc: sqlc,
	}
}

func (repo *TokenRepo) Create(ctx context.Context, token entity.Token) error {
	args := postgres.CreateTokenParams{
		UserID: int64(token.UserID),
		Hash:   token.Hash,
		Scope:  postgres.TokensScope(token.Scope),
		Expiry: token.Expiry,
	}
	_, err := repo.dbc.CreateToken(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (repo *TokenRepo) DeleteAllForUser(ctx context.Context, userId int) error {
	_, err := repo.dbc.DeleteAllTokensForUser(ctx, int64(userId))
	if err != nil {
		return err
	}
	return nil
}

func (repo *TokenRepo) DeleteScopedForUser(ctx context.Context, userId int, scope string) error {
	args := postgres.DeleteScopeTokenForUserParams{
		UserID: int64(userId),
		Scope:  postgres.TokensScope(scope),
	}
	_, err := repo.dbc.DeleteScopeTokenForUser(ctx, args)
	if err != nil {
		return err
	}
	return nil
}
