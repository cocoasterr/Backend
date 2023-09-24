package service

import (
	"encoding/json"

	PGRepository "github.com/cocoasterr/Backend/go_app/arch/db/pg/gorm/repository"
)

type ServiceInterface interface {
	CreateService(data interface{}) error
	IndexService(data interface{}, limit, page int) ([]map[string]interface{}, error)
}

type ServiceRepo struct {
	Repo PGRepository.Repository
}

func NewService(repo PGRepository.Repository) *ServiceRepo {
	return &ServiceRepo{
		Repo: repo,
	}
}

func (s *ServiceRepo) CreateService(data interface{}) error {
	if err := s.Repo.Create(data); err != nil {
		return err
	}
	return nil
}
func (s *ServiceRepo) IndexService(data interface{}, limit, page int) ([]map[string]interface{}, error) {
	offset := (page - 1) * limit
	res := s.Repo.Index(data, limit, offset)
	resJson, _ := json.Marshal(res)

	var result []map[string]interface{}
	if err := json.Unmarshal(resJson, &result); err != nil {
		return nil, err
	}
	return result, nil
}
