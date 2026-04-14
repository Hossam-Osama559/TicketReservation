package payment

type PaymentRepositroy interface {
	Create(instance Payment) error

	Get(sessionid string) (*Payment, error)

	Update(sessionid string, instance Payment) error

	Delete(sessionid string) error

	MarkCompleted(sessionid string) error
}
