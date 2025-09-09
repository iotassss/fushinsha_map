package gormrepo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
	"gorm.io/gorm"
)

type SuspiciousPersonModel struct {
	UUID         string  `gorm:"column:uuid;primaryKey;not null"`
	Emoji        string  `gorm:"column:emoji;not null"`
	Sign         string  `gorm:"column:sign;not null"`
	ResisterUUID string  `gorm:"column:register_uuid;not null"`
	SightedCount int     `gorm:"column:sighted_count"`
	SightingTime string  `gorm:"column:sighting_time"`
	X            float64 `gorm:"column:x;not null"`
	Y            float64 `gorm:"column:y;not null"`
	Gender       string  `gorm:"column:gender"`
	Clothing     string  `gorm:"column:clothing"`
	Accessories  string  `gorm:"column:accessories"`
	Vehicle      string  `gorm:"column:vehicle"`
	Behavior     string  `gorm:"column:behavior"`
	Hairstyle    string  `gorm:"column:hairstyle"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (SuspiciousPersonModel) TableName() string {
	return "suspicious_persons"
}

type SuspiciousPersonRepository struct {
	db *gorm.DB
}

func NewSuspiciousPersonRepository(db *gorm.DB) *SuspiciousPersonRepository {
	return &SuspiciousPersonRepository{
		db: db,
	}
}

func toDomain(model SuspiciousPersonModel) (domain.SuspiciousPerson, error) {
	uuid, err := domain.NewUUID(model.UUID)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	emoji, err := domain.NewEmoji(model.Emoji)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	sign, err := domain.NewSign(model.Sign)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	resisterUUID, err := domain.NewUUID(model.ResisterUUID)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	sightingCount, err := domain.NewSightingCount(model.SightedCount)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	st, err := time.Parse("15:04", model.SightingTime)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	sightingTime, err := domain.NewSightingTime(st)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	coordinates, err := domain.NewCoordinates(model.Y, model.X)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	gender, err := domain.NewGender(model.Gender)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	clothing, err := domain.NewClothing(model.Clothing)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	accessories, err := domain.NewAccessories(model.Accessories)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	vehicle, err := domain.NewVehicle(model.Vehicle)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	behavior, err := domain.NewBehavior(model.Behavior)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	hairstyle, err := domain.NewHairstyle(model.Hairstyle)
	if err != nil {
		return domain.SuspiciousPerson{}, fmt.Errorf("%w: %v", domain.ErrValidation, err)
	}
	return domain.NewSuspiciousPerson(
		uuid,
		emoji,
		sign,
		resisterUUID,
		sightingCount,
		sightingTime,
		coordinates,
		&gender,
		&clothing,
		&accessories,
		&vehicle,
		&behavior,
		&hairstyle,
	), nil
}

func toModel(person *domain.SuspiciousPerson) SuspiciousPersonModel {
	personModel := SuspiciousPersonModel{
		UUID:         person.UUID().String(),
		Emoji:        person.Emoji().String(),
		Sign:         person.Sign().String(),
		ResisterUUID: person.RegistrarUUID().String(),
		SightedCount: person.SightingCount().Int(),
		SightingTime: person.SightingTime().Time().Format("15:04"),
		X:            person.Coordinates().Longitude(),
		Y:            person.Coordinates().Latitude(),
		Gender:       person.Gender().String(),
	}
	if person.Clothing() != nil {
		personModel.Clothing = person.Clothing().String()
	}
	if person.Accessories() != nil {
		personModel.Accessories = person.Accessories().String()
	}
	if person.Vehicle() != nil {
		personModel.Vehicle = person.Vehicle().String()
	}
	if person.Behavior() != nil {
		personModel.Behavior = person.Behavior().String()
	}
	if person.Hairstyle() != nil {
		personModel.Hairstyle = person.Hairstyle().String()
	}
	return personModel
}

func (r *SuspiciousPersonRepository) FindInArea(ctx context.Context, area domain.Area) ([]domain.SuspiciousPerson, error) {
	var models []SuspiciousPersonModel
	err := r.db.WithContext(ctx).
		Where("x BETWEEN ? AND ? AND y BETWEEN ? AND ?", area.LX(), area.RX(), area.BY(), area.TY()).
		Find(&models).Error
	if err != nil {
		return nil, fmt.Errorf("%w: %v", domain.ErrRepository, err)
	}

	var result []domain.SuspiciousPerson
	for _, m := range models {
		p, err := toDomain(m)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *SuspiciousPersonRepository) FindByUUID(ctx context.Context, uuid domain.UUID) (*domain.SuspiciousPerson, error) {
	var model SuspiciousPersonModel
	err := r.db.WithContext(ctx).Where("uuid = ?", uuid.String()).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: %v", domain.ErrNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", domain.ErrRepository, err)
	}

	p, err := toDomain(model)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *SuspiciousPersonRepository) Create(ctx context.Context, person *domain.SuspiciousPerson) error {
	model := toModel(person)
	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return fmt.Errorf("%w: %v", domain.ErrRepository, err)
	}
	return nil
}

func (r *SuspiciousPersonRepository) Update(ctx context.Context, person *domain.SuspiciousPerson) error {
	model := toModel(person)
	tx := r.db.WithContext(ctx).Model(&model).Where("uuid = ?", model.UUID).Updates(&model)
	if tx.Error != nil {
		return fmt.Errorf("%w: %v", domain.ErrRepository, tx.Error)
	}
	if tx.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
