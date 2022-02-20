package service

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository"
	"1aidar1/bastau/go-api/pkg/auth"
	"1aidar1/bastau/go-api/pkg/hash"
	"context"
	"errors"
	"fmt"
	"time"
)

type UsersService struct {
	repo         repository.UsersRepoI
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	//otpGenerator otp.Generator
	//dnsService   dns.DomainManager

	//emailService  Emails
	//schoolService Schools

	//accessTokenTTL         time.Duration
	//refreshTokenTTL        time.Duration
	//verificationCodeLength int

	//domain string
}

var _ UsersI = &UsersService{}

var EmptyUser entity.User

func NewUsersService(repo repository.UsersRepoI, hasher hash.PasswordHasher, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		repo:         repo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (s *UsersService) GetUserById(ctx context.Context, id int) (entity.User, error) {
	fmt.Println(s.tokenManager)
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return EmptyUser, err
	}
	return *user, nil
}

func (s *UsersService) Register(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:         input.Name,
		Password:     passwordHash,
		Phone:        input.Phone,
		Email:        input.Email,
		IsMale:       input.IsMale,
		CountryId:    input.CountryId,
		CityId:       input.CityId,
		DateOfBirth:  input.DateOfBirth,
		RegisteredAt: time.Now(),
		UpdatedAt:    time.Now(),
		LastVisitAt:  time.Now(),
	}

	if err := s.repo.Register(ctx, user); err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return err
		}

		return err
	}
	return nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (Tokens, error) {
	//passwordHash, err := s.hasher.Hash(input.Password)
	//if err != nil {
	//	return Tokens{}, err
	//}
	//
	//user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
	//if err != nil {
	//	if errors.Is(err, domain.ErrUserNotFound) {
	//		return Tokens{}, err
	//	}
	//
	//	return Tokens{}, err
	//}
	//
	//return s.createSession(ctx, user.ID)
	return Tokens{}, nil
}

func (s *UsersService) RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error) {
	//student, err := s.repo.GetByRefreshToken(ctx, refreshToken)
	//if err != nil {
	//	return Tokens{}, err
	//}
	//
	//return s.createSession(ctx, student.ID)
	return Tokens{}, nil
}

func (s *UsersService) Verify(ctx context.Context, userID int, hash string) error {
	//err := s.repo.Verify(ctx, userID, hash)
	//if err != nil {
	//	if errors.Is(err, domain.ErrVerificationCodeInvalid) {
	//		return err
	//	}
	//
	//	return err
	//}

	return nil
}

func (s *UsersService) createSession(ctx context.Context, userId int) (Tokens, error) {
	//var (
	//	res Tokens
	//	err error
	//)
	//
	//res.AccessToken, err = s.tokenManager.NewJWT(userId.Hex(), s.accessTokenTTL)
	//if err != nil {
	//	return res, err
	//}
	//
	//res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	//if err != nil {
	//	return res, err
	//}
	//
	//session := domain.Session{
	//	RefreshToken: res.RefreshToken,
	//	ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	//}
	//
	//err = s.repo.SetSession(ctx, userId, session)

	return Tokens{}, nil
}
