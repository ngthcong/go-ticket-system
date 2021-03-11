package user

import (
	"errors"
	"github.com/form3tech-oss/jwt-go"
	"go-ticket-system/internal/model"
	"go-ticket-system/internal/repository/user"
	"go.uber.org/zap"
	"os"
	"time"
)

type (
	Service struct {
		userRepo user.UserRepository
		logger   *zap.SugaredLogger
	}
	UserService interface {
		Get(id int) (*model.Users, error)
		Login(req model.LoginRequest) (string, error)
		GetAsset(id int) ([]model.Assets, error)
	}
)

func New(userRepo user.UserRepository, logger *zap.SugaredLogger) UserService {
	return &Service{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *Service) Get(id int) (*model.Users, error) {
	s.logger.Infof("Get user by id: %v", id)
	users, err := s.userRepo.GetByID(id)
	if err != nil {
		s.logger.Errorf("Get user error, %v", err)
	} else if users == nil {
		s.logger.Info("No user found")
		return nil, errors.New("no user found")
	}
	s.logger.Infof("User: %v", *users)
	return users, err
}
func (s *Service) GetAsset(id int) ([]model.Assets, error) {
	s.logger.Infof("Get asset by user id: %v", id)
	assets, err := s.userRepo.GetAsset(id)
	if err != nil {
		s.logger.Errorf("Get assets error, %v", err)
	} else if assets == nil {
		s.logger.Info("No assets found")
		return nil, errors.New("no asset found")
	}
	s.logger.Infof("Assets: %v", assets)
	return assets, err
}
func (s *Service) Login(req model.LoginRequest) (string, error) {
	s.logger.Infof("User login req: %v", req)
	users, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		s.logger.Errorf("Get user error, %v", err)
	} else if users == nil {
		s.logger.Info("No user found")
		return "", errors.New("email or password invalid")
	}
	s.logger.Infof("User: , %v", *users)
	if req.Email == users.Password {
		s.logger.Info("User login valid")

	}
	token, err := CreateToken(users.ID,users.Role)
	if err != nil {
		return "", err
	}

	return token, err
}

func CreateToken(userid int,role int) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["role"] = role
	atClaims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
