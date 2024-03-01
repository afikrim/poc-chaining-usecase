package port

//go:generate mockgen -source=usecase.go -destination=mock/usecase_mock.go -package=mock
type Usecase interface {
	GetCustomerUsecase() CustomerUsecase
}
