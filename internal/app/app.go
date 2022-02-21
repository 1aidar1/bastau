package app

import (
	"1aidar1/bastau/go-api/config"
	delivery "1aidar1/bastau/go-api/internal/controller/http"
	"1aidar1/bastau/go-api/internal/repository"
	"1aidar1/bastau/go-api/internal/service"
	"1aidar1/bastau/go-api/pkg/auth"
	"1aidar1/bastau/go-api/pkg/hash"
	"1aidar1/bastau/go-api/pkg/httpserver"
	"1aidar1/bastau/go-api/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config, l logger.LoggerI) {

	//Repository
	//pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	//if err != nil {
	//	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	//}
	//defer pg.Close()
	db, err := openDB(cfg, l)
	defer db.Close()
	if err != nil {
		l.Fatal("Can't connect to DB ", err)
	}
	repos := repository.NewRepositories(db)

	// Use case
	//translationUseCase := usecase.New(
	//	repo.New(pg),
	//	webapi.New(),
	//)

	tokenManager, err := auth.NewManager(cfg.Auth.JWTSecret)
	if err != nil {
		l.Error(err)

		return
	}
	services := service.NewServices(service.Deps{
		Repos:        repos,
		Hasher:       hash.NewSHA1Hasher(cfg.Auth.PasswordSalt),
		TokenManager: tokenManager,
	})

	// HTTP Server
	hndlrs := delivery.NewHandler(services, tokenManager)
	handler := hndlrs.Init(cfg)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}

	}
}

func openDB(cfg *config.Config, l logger.LoggerI) (*pgxpool.Pool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conf, err := pgxpool.ParseConfig(cfg.PG.URL)
	if err != nil {
		return nil, err
	}
	conf.MaxConns = int32(cfg.PG.PoolMax)
	conf.MaxConnIdleTime = time.Minute * time.Duration(cfg.PG.MaxIdleTime)
	conf.ConnConfig.Logger = l
	pool, err := pgxpool.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
