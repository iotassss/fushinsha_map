// internal/repository/memory/suspicious_person_repository.go
package memory

import (
	"sync"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
)

type suspiciousPersonRepository struct {
	mu    sync.RWMutex
	store map[domain.UUID]*domain.SuspiciousPerson
}

func NewSuspiciousPersonRepository() domain.SuspiciousPersonRepository {
	return &suspiciousPersonRepository{
		store: make(map[domain.UUID]*domain.SuspiciousPerson),
	}
}

// 指定範囲内の不審者一覧取得
func (r *suspiciousPersonRepository) FindInArea(lx, rx, ty, by float64) ([]domain.SuspiciousPerson, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.SuspiciousPerson
	for _, p := range r.store {
		coord := p.Coordinates()
		if coord.Longitude >= lx && coord.Longitude <= rx && coord.Latitude >= by && coord.Latitude <= ty {
			result = append(result, *p)
		}
	}
	return result, nil
}

// UUIDで不審者詳細取得
func (r *suspiciousPersonRepository) FindByUUID(uuid domain.UUID) (*domain.SuspiciousPerson, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, ok := r.store[uuid]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return p, nil
}

// 不審者新規登録
func (r *suspiciousPersonRepository) Create(person *domain.SuspiciousPerson) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[person.UUID()]; exists {
		return domain.ErrAlreadyExists
	}
	r.store[person.UUID()] = person
	return nil
}

// 不審者情報編集
func (r *suspiciousPersonRepository) Update(person *domain.SuspiciousPerson) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[person.UUID()]; !exists {
		return domain.ErrNotFound
	}
	r.store[person.UUID()] = person
	return nil
}

// 目撃カウント追加
func (r *suspiciousPersonRepository) AddSightingCount(uuid domain.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.store[uuid]
	if !ok {
		return domain.ErrNotFound
	}
	p.SetSightingCount(p.SightingCount() + 1)
	return nil
}

// 目撃時刻追加
func (r *suspiciousPersonRepository) AddSightingTime(uuid domain.UUID, time domain.SightingTime) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.store[uuid]
	if !ok {
		return domain.ErrNotFound
	}
	p.SetSightingTime(time)
	return nil
}
