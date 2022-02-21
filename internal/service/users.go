package service

import (
	"1aidar1/bastau/go-api/internal/entity"
	"1aidar1/bastau/go-api/internal/repository"
	"1aidar1/bastau/go-api/internal/repository/postgres"
	"1aidar1/bastau/go-api/pkg/auth"
	"1aidar1/bastau/go-api/pkg/hash"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"
)

type UsersService struct {
	repo         repository.UsersRepoI
	tokenRepo    repository.TokenRepoI
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

func NewUsersService(repo repository.UsersRepoI, tokenRepo repository.TokenRepoI, hasher hash.PasswordHasher, tokenManager auth.TokenManager) *UsersService {
	return &UsersService{
		repo:         repo,
		tokenRepo:    tokenRepo,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}

func (s *UsersService) GetUserById(ctx context.Context, id int) (entity.User, error) {
	//fmt.Println(s.tokenManager)
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return entity.EmptyUser, err
	}
	return user, nil
}

func (s *UsersService) Register(ctx context.Context, input UserSignUpInput) error {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return err
	}

	user := entity.User{
		Name:     input.Name,
		Password: []byte(passwordHash),
		Phone:    input.Phone,
		Email:    input.Email,
		Role:     input.Role,
		//IsMale:       input.IsMale,
		//CountryId:    input.CountryId,
		//CityId:       input.CityId,
		//DateOfBirth:  input.DateOfBirth,
		//RegisteredAt: time.Now(),
		//UpdatedAt:    time.Now(),
		//LastVisitAt:  time.Now(),
	}

	if err := s.repo.Register(ctx, user); err != nil {
		if errors.Is(err, entity.ErrUserAlreadyExists) {
			return err
		}

		return err
	}
	return nil
}

func (s *UsersService) SignIn(ctx context.Context, input UserSignInInput) (entity.Token, error) {
	passwordHash, err := s.hasher.Hash(input.Password)
	if err != nil {
		return entity.EmptyToken, err
	}
	//
	user, err := s.repo.GetByCredentials(ctx, input.Email, passwordHash)
	if err != nil {
		return entity.EmptyToken, err
	}
	//
	//return s.createSession(ctx, user.ID)
	token, err := s.CreateToken(ctx, user.ID, string(postgres.TokensScopeAuth))
	if err != nil {
		return entity.EmptyToken, err
	}
	return token, nil
}

func (s *UsersService) GetUserByToken(ctx context.Context, hash []byte) (entity.User, error) {
	tokenHash := sha256.Sum256([]byte(hash))
	//fmt.Println("BOBA>>>>>>", hash, " AAAAAAAAAAAAA>>>>>", tokenHash)
	user, err := s.repo.GetByToken(ctx, tokenHash[:])

	if err != nil {
		return entity.EmptyUser, err
	}
	return user, nil
}

func (s *UsersService) CreateToken(ctx context.Context, userId int, scope string) (entity.Token, error) {
	token, err := generateToken(userId, time.Hour*72, scope)
	if err != nil {
		return entity.EmptyToken, err
	}
	err = s.tokenRepo.Create(ctx, token)
	//
	//res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return entity.EmptyToken, err
	}
	//
	//session := domain.Session{
	//	RefreshToken: res.RefreshToken,
	//	ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	//}
	//
	//err = s.repo.SetSession(ctx, userId, session)
	return token, nil
}

func generateToken(userID int, ttl time.Duration, scope string) (entity.Token, error) {
	// Create a Token instance containing the user ID, expiry, and scope information.
	// Notice that we add the provided ttl (time-to-live) duration parameter to the
	// current time to get the expiry time?
	token := entity.Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}
	// Initialize a zero-valued byte slice with a length of 16 bytes.
	randomBytes := make([]byte, 16)
	// Use the Read() function from the crypto/rand package to fill the byte slice with
	// random bytes from your operating system's CSPRNG. This will return an error if
	// the CSPRNG fails to function correctly.
	_, err := rand.Read(randomBytes)
	if err != nil {
		return entity.EmptyToken, err
	}
	// Encode the byte slice to a base-32-encoded string and assign it to the token
	// Plaintext field. This will be the token string that we send to the user in their
	// welcome email. They will look similar to this:
	//
	// Y3QMGX3PJ3WLRL2YRTQGQ6KRHU
	//
	// Note that by default base-32 strings may be padded at the end with the =
	// character. We don't need this padding character for the purpose of our tokens, so
	// we use the WithPadding(base32.NoPadding) method in the line below to omit them.
	plaintext := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	// Generate a SHA-256 hash of the plaintext token string. This will be the value
	// that we store in the `hash` field of our database table. Note that the
	// sha256.Sum256() function returns an *array* of length 32, so to make it easier to
	// work with we convert it to a slice using the [:] operator before storing it.

	hash := sha256.Sum256([]byte(plaintext))
	token.PlainPassword = plaintext
	token.Hash = hash[:]
	return token, nil
}
