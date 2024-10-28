package service

type userStore interface {
}

type tokenStore interface {
}

type Service struct {
	us userStore
	ts tokenStore
}

func New(us userStore, ts tokenStore) *Service {
	return &Service{
		us: us,
		ts: ts,
	}
}
