package usecase

import (
	"context"
	"fmt"

	"github.com/iotassss/fushinsha-map-api/internal/domain"
	"github.com/iotassss/fushinsha-map-api/internal/port"
)

type FindOrCreateUserByIDTokenInteractor struct {
	userRepo domain.UserRepository
	// authVerifier port.AuthVerifier
	// jwtParser    port.JWTParser

	googleAuthVerifier port.GoogleAuthVerifier
}

func NewFindOrCreateUserByIDTokenInteractor(
	userRepo domain.UserRepository,
	// authVerifier port.AuthVerifier,
	// jwtParser port.JWTParser,

	googleAuthVerifier port.GoogleAuthVerifier,
) *FindOrCreateUserByIDTokenInteractor {
	return &FindOrCreateUserByIDTokenInteractor{
		userRepo: userRepo,
		// authVerifier: authVerifier,
		// jwtParser:    jwtParser,

		googleAuthVerifier: googleAuthVerifier,
	}
}

func (uc *FindOrCreateUserByIDTokenInteractor) Execute(
	ctx context.Context,
	input FindOrCreateUserByIDTokenInput,
	presenter FindOrCreateUserByIDTokenPresenter,
) error {
	// // 1. IDトークンの検証
	// if err := uc.authVerifier.Verify(ctx, input.IDToken); err != nil {
	// 	return presenter.PresentError(fmt.Errorf("%w: %v", ErrUnauthorized, err))
	// }

	// // 2. JWTのパース
	// token, err := uc.jwtParser.Parse(input.IDToken)
	// if err != nil {
	// 	return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	// }

	// 上記の処理をまとめて実行できるGoogle公式ライブラリがあるため1,2を統合する

	// 1. 2. IDトークンを検証してパース済みトークンを取得
	token, err := uc.googleAuthVerifier.VerifyAndParse(ctx, input.IDToken)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrUnauthorized, err))
	}

	// 3. GoogleAccountID生成
	googleAccountID, err := domain.NewGoogleAccountID(token.Claims.Subject)
	if err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrValidation, err))
	}

	// 4. ユーザー検索
	user, err := uc.userRepo.FindByGoogleAccountID(ctx, googleAccountID)
	if err == nil && user != nil {
		// 既存ユーザー
		return presenter.Present(User{
			UUID: user.UUID().String(),
		})
	}
	// DBエラー以外のnot found以外はエラー返却
	if err != nil && err != domain.ErrNotFound {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrInternal, err))
	}

	// 5. 新規ユーザー作成
	uuid := domain.GenerateUUID()
	newUser := domain.NewUser(uuid, googleAccountID)
	if err := uc.userRepo.Create(ctx, &newUser); err != nil {
		return presenter.PresentError(fmt.Errorf("%w: %v", ErrInternal, err))
	}

	return presenter.Present(User{
		UUID: newUser.UUID().String(),
	})
}
