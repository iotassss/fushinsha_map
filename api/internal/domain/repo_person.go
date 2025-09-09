package domain

import "context"

type PersonRepository interface {
	// 指定範囲内の不審者一覧取得
	FindInArea(ctx context.Context, area Area) ([]Person, error)
	// UUIDで不審者詳細取得
	FindByUUID(ctx context.Context, uuid UUID) (*Person, error)
	// 不審者新規登録
	Create(ctx context.Context, person *Person) error
	// 不審者情報編集
	Update(ctx context.Context, person *Person) error
}
