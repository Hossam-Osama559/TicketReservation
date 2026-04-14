package payment

// the payment service will be responsible for
type PaymentService struct {
	repo PaymentRepositroy
}

func NewPaymentService(repo PaymentRepositroy) *PaymentService {

	return &PaymentService{repo: repo}
}

func (service *PaymentService) Create(instance Payment) error {

	return service.repo.Create(instance)
}

func (service *PaymentService) Get(sessionid string) (*Payment, error) {

	return service.repo.Get(sessionid)
}

func (service *PaymentService) Update(sessionid string, instance Payment) error {

	return service.repo.Update(sessionid, instance)
}

func (service *PaymentService) Delete(sessionid string) error {

	return service.repo.Delete(sessionid)
}

func (service *PaymentService) MarkCompleter(sessionid string) error {

	return service.repo.MarkCompleted(sessionid)
}
