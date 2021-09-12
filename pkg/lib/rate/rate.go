package rate

type Limiter interface {
	IsAllow(ip string) (bool, error)
}
