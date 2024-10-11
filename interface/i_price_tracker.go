package interfaces

type IPriceTracker interface {
	GetUrl() string
	GetPrice() float64
}
