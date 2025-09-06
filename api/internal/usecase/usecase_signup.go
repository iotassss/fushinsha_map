package usecase

type SignupInputData struct {
	GoogleAccountID string
}

type SignupOutputData struct {
	Success bool
}

type SignupPresenter interface {
	Present(outputData SignupOutputData) error
	PresentError(err error) error
}

type SignupUsecase interface {
	Execute(
		input SignupInputData,
		presenter SignupPresenter,
	) error
}
