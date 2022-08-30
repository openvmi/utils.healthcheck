package health

type IHealthHandler interface {
	GetStatus() string
}
