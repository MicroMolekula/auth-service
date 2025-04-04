package repository

import (
	"github.com/MicroMolekula/auth-service/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) FindByIds(ids []uint) (map[uint]*models.User, error) {
	var users []*models.User
	if err := ur.db.Where("id in (?)", ids).Find(&users).Error; err != nil {
		return nil, err
	}
	result := make(map[uint]*models.User)
	for _, user := range users {
		result[user.ID] = user
	}
	return result, nil
}

func (ur *UserRepository) FindOneByEmail(email string) (*models.User, error) {
	var user *models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindOneByYandexId(id string) (*models.User, error) {
	var user *models.User
	if err := ur.db.Where("yandex_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
