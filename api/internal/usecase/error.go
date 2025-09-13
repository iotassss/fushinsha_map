package usecase

import "errors"

var (
	// バリデーションエラー domainでvoなどのバリデーションに失敗した場合など 400
	ErrValidation = errors.New("validation error")
	// 認証エラー（未ログインやトークン無効） 401
	ErrUnauthorized = errors.New("unauthorized")
	// 認可エラー（ログイン済みだが権限不足） 403
	ErrForbidden = errors.New("forbidden")
	// リソース未発見 404
	ErrNotFound = errors.New("not found")
	// ビジネスルール違反 domain serviceでのビジネスルール違反など 400/422
	ErrBusinessRule = errors.New("business rule violation")
	// 外部サービス連携エラー 502
	ErrExternal = errors.New("external service error")
	// 内部エラー 500
	ErrInternal = errors.New("internal service error")
)
