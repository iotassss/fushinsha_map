package domain

// SuspiciousPersonRepository defines repository interface for suspicious persons.
type SuspiciousPersonRepository interface {
	// 指定範囲内の不審者一覧取得
	FindInArea(lx, rx, ty, by float64) ([]SuspiciousPerson, error)
	// UUIDで不審者詳細取得
	FindByUUID(uuid UUID) (*SuspiciousPerson, error)
	// 不審者新規登録
	Create(person *SuspiciousPerson) error
	// 不審者情報編集
	Update(person *SuspiciousPerson) error
	// 目撃カウント追加
	AddSightingCount(uuid UUID) error
	// 目撃時刻追加
	AddSightingTime(uuid UUID, time SightingTime) error
}
