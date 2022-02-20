package service

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository"
	"1aidar1/bastau/go-api/pkg/auth"
	"1aidar1/bastau/go-api/pkg/hash"
	"context"
	"time"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

// TODO handle "not found" errors

type UserSignUpInput struct {
	Name        string
	Email       string
	Phone       string
	Password    string
	IsMale      bool
	CountryId   int
	CityId      int
	DateOfBirth time.Time
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// 1. Create School in DB
// 2. Generate Sub Domain

type UsersI interface {
	GetUserById(ctx context.Context, id int) (entity.User, error)
	Register(ctx context.Context, input UserSignUpInput) error
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	Verify(ctx context.Context, userID int, hash string) error
}

type Services struct {
	Users UsersI
}

type Deps struct {
	Repos *repository.Repositories
	//Cache                  cache.Cache
	Hasher       hash.PasswordHasher
	TokenManager auth.TokenManager
	//EmailSender            email.Sender
	//EmailConfig            config.EmailConfig
	//StorageProvider        storage.Provider
	//AccessTokenTTL         time.Duration
	//RefreshTokenTTL        time.Duration
	//FondyCallbackURL       string
	//CacheTTL               int64
	//OtpGenerator           otp.Generator
	//VerificationCodeLength int
	//Environment            string
	//Domain                 string
	//DNS                    dns.DomainManager
}

func NewServices(deps Deps) *Services {
	usersService := NewUsersService(deps.Repos.Users, deps.Hasher, deps.TokenManager)

	return &Services{
		Users: usersService,
	}
}
