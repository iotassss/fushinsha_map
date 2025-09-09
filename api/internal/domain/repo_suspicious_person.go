package domain

import "context"

type SuspiciousPersonRepository interface {
	// 指定範囲内の不審者一覧取得
	FindInArea(ctx context.Context, area Area) ([]SuspiciousPerson, error)
	// UUIDで不審者詳細取得
	FindByUUID(ctx context.Context, uuid UUID) (*SuspiciousPerson, error)
	// 不審者新規登録
	Create(ctx context.Context, person *SuspiciousPerson) error
	// 不審者情報編集
	Update(ctx context.Context, person *SuspiciousPerson) error
}
