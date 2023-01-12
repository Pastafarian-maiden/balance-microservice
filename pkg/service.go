package pkg

type Service interface {
	AddToBalance(userID, fund int) error
	ReserveFund(userID, serviceID, orderID, price int) error
	RevenueRecognition(userID, serviceID, orderID, price int) error
	CheckBalance(userID int) error
}

type dateService struct{}

func NewService() Service {
	return dateService{}
}

func (dateService) AddToBalance(userID, fund int) error {
	return nil
}

func (dateService) ReserveFund(userID, serviceID, orderID, price int) error {
	return nil
}

func (dateService) RevenueRecognition(userID, serviceID, orderID, price int) error {
	return nil
}

func (dateService) CheckBalance(userID int) error {
	return nil
}
