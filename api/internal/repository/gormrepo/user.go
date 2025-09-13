package gormrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
	"gorm.io/gorm"
)

type UserModel struct {
	UUID            string `gorm:"column:uuid;primaryKey;not null"`
	GoogleAccountID string `gorm:"column:google_account_id;size:191;uniqueIndex;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByGoogleAccountID(ctx context.Context, googleAccountID domain.GoogleAccountID) (*domain.User, error) {
	var m UserModel
	if err := r.db.WithContext(ctx).Where("google_account_id = ?", googleAccountID.String()).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("%w: %v", domain.ErrRepository, err)
	}

	userUUID, err := domain.NewUUID(m.UUID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	googleAccountIDVO, err := domain.NewGoogleAccountID(m.GoogleAccountID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}

	user := domain.NewUser(userUUID, googleAccountIDVO)

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	UserModel := UserModel{
		UUID:            user.UUID().String(),
		GoogleAccountID: user.GoogleAccountID().String(),
	}
	return r.db.WithContext(ctx).Create(&UserModel).Error
}
