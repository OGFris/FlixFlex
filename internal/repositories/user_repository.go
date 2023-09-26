package repositories

import (
	"github.com/OGFris/FlixFlex/internal/models"
	"github.com/OGFris/FlixFlex/pkg/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
}

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {

	return &mysqlUserRepository{db: db}
}

func (r *mysqlUserRepository) FindByID(id uint) (*models.User, error) {
	user := &models.User{}
	result := r.db.First(user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {

			return nil, errors.ErrUserNotFound
		}

		return nil, errors.ErrDatabase
	}

	return user, nil
}

func (r *mysqlUserRepository) Create(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {

		return errors.ErrDatabase
	}

	return nil
}

func (r *mysqlUserRepository) Update(user *models.User) error {
	result := r.db.Save(user)
	if result.Error != nil {

		return errors.ErrDatabase
	}

	return nil
}

func (r *mysqlUserRepository) Delete(user *models.User) error {
	result := r.db.Delete(user)
	if result.Error != nil {

		return errors.ErrDatabase
	}

	return nil
}