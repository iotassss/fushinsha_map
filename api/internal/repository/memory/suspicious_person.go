// internal/repository/memory/person_repository.go
package memory

// import (
// 	"sync"

// 	"github.com/iotassss/fushinsha-map-api/internal/domain"
// )

// type PersonRepository struct {
// 	mu    sync.RWMutex
// 	store map[domain.UUID]*domain.Person
// }

// func NewPersonRepository() domain.PersonRepository {
// 	return &personRepository{
// 		store: make(map[domain.UUID]*domain.Person),
// 	}
// }

// // 指定範囲内の不審者一覧取得
// func (r *personRepository) FindInArea(area domain.Area) ([]domain.Person, error) {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	var result []domain.Person
// 	for _, p := range r.store {
// 		coord := p.Coordinates()
// 		if coord.Longitude() >= area.LX() &&
// 			coord.Longitude() <= area.RX() &&
// 			coord.Latitude() >= area.BY() &&
// 			coord.Latitude() <= area.TY() {
// 			result = append(result, *p)
// 		}
// 	}
// 	return result, nil
// }

// // UUIDで不審者詳細取得
// func (r *personRepository) FindByUUID(uuid domain.UUID) (*domain.Person, error) {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	p, ok := r.store[uuid]
// 	if !ok {
// 		return nil, domain.ErrNotFound
// 	}
// 	return p, nil
// }

// // 不審者新規登録
// func (r *personRepository) Create(person *domain.Person) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	if _, exists := r.store[person.UUID()]; exists {
// 		return domain.ErrAlreadyExists
// 	}
// 	r.store[person.UUID()] = person
// 	return nil
// }

// // 不審者情報編集
// func (r *personRepository) Update(person *domain.Person) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	if _, exists := r.store[person.UUID()]; !exists {
// 		return domain.ErrNotFound
// 	}
// 	r.store[person.UUID()] = person
// 	return nil
// }

// // 目撃カウント追加
// func (r *personRepository) AddSightingCount(uuid domain.UUID) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	p, ok := r.store[uuid]
// 	if !ok {
// 		return domain.ErrNotFound
// 	}
// 	p.SetSightingCount(p.SightingCount() + 1)
// 	return nil
// }

// // 目撃時刻追加
// func (r *personRepository) AddSightingTime(uuid domain.UUID, time domain.SightingTime) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	p, ok := r.store[uuid]
// 	if !ok {
// 		return domain.ErrNotFound
// 	}
// 	p.SetSightingTime(time)
// 	return nil
// }
