package user

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(userRequest UserRequest) (User, error)
}

type service struct {
	repository RepositoryUser
}

func NewService(repository RepositoryUser) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(userRequest UserRequest) (User, error) {

	age, _ := userRequest.Age.Int64()

	user := User{
		Name:     userRequest.Name,
		Age:      int(age),
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}
