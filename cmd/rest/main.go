// auth-service: OAuth2 Authorization Server tối giản để học (thay cho Keycloak).
//
//	POST /oauth/token            grant_type=password | client_credentials
//	GET  /.well-known/jwks.json  public key để các service verify JWT
//
// Production: dùng Keycloak/Auth0/Ory Hydra, backoffice dùng Authorization Code + PKCE.
package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/phillong97254/auth-service/cmd/rest/routes"
	"github.com/phillong97254/auth-service/config"
	"github.com/phillong97254/auth-service/internal/gateway"
	v1 "github.com/phillong97254/auth-service/internal/handler/rest/v1"
	"github.com/phillong97254/auth-service/internal/repository"
	"github.com/phillong97254/auth-service/internal/service/token"
	"github.com/phillong97254/auth-service/internal/signer"
)

func main() {
	// Load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)).With("server", cfg.RestServer.Name, "env", cfg.API.Env)
	slog.SetDefault(logger)

	// Signing key (demo: sinh mới mỗi lần khởi động, các service tự lấy lại qua JWKS)
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	sign := signer.New(cfg, key)

	// Gateway: client gọi authz-service (role/scope của user)
	gwy, err := gateway.New(cfg, sign)
	if err != nil {
		log.Fatal(err)
	}
	defer gwy.Close()

	// Repository -> Service -> Handler
	repo := repository.New(cfg)
	tokenSvc := token.New(cfg, logger, repo, gwy, sign)
	handler := v1.New(logger, tokenSvc)

	// Register Routes
	mux := http.NewServeMux()
	routes.V1Routes(mux, handler)

	// Start Server
	srv := &http.Server{Addr: cfg.RestServer.Addr, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		logger.Info("auth-service listening", "addr", cfg.RestServer.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.RestServer.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "error", err)
	}
	logger.Info("auth-service stopped")
}
