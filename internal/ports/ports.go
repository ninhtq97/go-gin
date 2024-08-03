package ports

type Server interface {
	Start() error
	Wait()
}

type Service interface {
	User() UserService
}

type Repository interface {
	User() UserRepository
}
