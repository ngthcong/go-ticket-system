package service

import (
	"go-ticket-system/internal/repository"
	"go-ticket-system/internal/service/user"
	"go.uber.org/zap"
)

type (
	Service struct {
		UserService user.UserService
	}
)

func New(repo repository.Repository, logger *zap.SugaredLogger) Service {
	userRepo := user.New(repo.UserRepo, logger)

	return Service{
		UserService: userRepo,
	}
}
