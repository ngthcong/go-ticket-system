package user

import (
	"go-ticket-system/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type (
	Repository struct {
		db     *gorm.DB
		logger *zap.SugaredLogger
	}

	UserRepository interface {
		GetByID(id int) (*model.Users, error)
		GetByEmail(email string) (*model.Users, error)
		GetAll() ([]model.Users, error)
		Update(user model.Users) error
		Delete(id int) error
		GetAsset(id int) ([]model.Assets, error)
	}
)

func New(db *gorm.DB, logger *zap.SugaredLogger) UserRepository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) GetByID(id int) (*model.Users, error) {
	var user model.Users
	records := r.db.Find(&user, id)
	return &user, records.Error
}
func (r *Repository) GetAsset(id int) ([]model.Assets, error) {
	var asset []model.Assets
	result := r.db.Raw("SELECT a.id,a.name,a.serial_number,a.type_id,a.user_id,a.description FROM assets a WHERE a.user_id = ?", id).Scan(&asset)
	return asset,result.Error
}
func (r *Repository) GetByEmail(email string) (*model.Users, error) {
	var user model.Users
	records := r.db.Where("email = ?", email).Find(&user)

	return &user, records.Error
}

func (r *Repository) GetAll() ([]model.Users, error) {
	var users []model.Users
	records := r.db.Find(&users)
	return users, records.Error
}

func (r *Repository) Update(user model.Users) error {
	return nil
}

func (r *Repository) Delete(id int) error {
	return nil
}
