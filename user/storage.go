package user

type Storage interface {
	Get(int) User
	Add(User) error
}
