package user

import "database/sql"

type UserService struct {
	repository *UserRepository
}

var Service UserService

func (s *UserService) InitService(db *sql.DB) {

	s.repository = &Repository
	s.repository.AssignDB(db)
}

func (s *UserService) createUser(payload UserDto) (sql.Result, error) {
	raw, err := s.repository.createUser(payload)

	return raw, err
}
