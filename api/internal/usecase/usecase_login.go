package usecase

type LoginInputData struct {
	GoogleAccountID string
}

type LoginOutputData struct {
	Token string
}

type LoginPresenter interface {
	Present(outputData LoginOutputData) error
	PresentError(err error) error
}

type LoginUsecase interface {
	Execute(
		input LoginInputData,
		presenter LoginPresenter,
	) error
}
