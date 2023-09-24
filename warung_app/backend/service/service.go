package service

import (
	PGRepository "github.com/cocoasterr/warung_app/infra/db/postgre/repository"
)

type Service struct {
	Repository PGRepository.Repository
}
type ServiceInterface interface {
	Create(domain interface{}) error
}

func NewService(repo PGRepository.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) Create(domain interface{}) error {
	if err := s.Repository.Create(domain); err != nil {
		return err
	}
	return nil
}

func (s *Service) FindBy(domain interface{}, key string, value interface{}) interface{} {
	return s.Repository.FindBy(domain, key, value)
}
