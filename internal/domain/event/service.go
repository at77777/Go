package event

type Service interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
	Create(id int64, name string) error
	Update(id int64, name string) error
	Delete(id int64, name string) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Event, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindOne(id int64) (*Event, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) Create(id int64, name string) error {
	return (*s.repo).Create(id, name)
}

func (s *service) Update(id int64, name string) error {
	return (*s.repo).Update(id, name)
}
func (s *service) Delete(id int64, name string) error {
	return (*s.repo).Delete(id, name)
}
