package service

import (
	PGGormRepository "github.com/cocoasterr/gowarungapp/infra/db/postgres/repository/gorm"
)

type ServiceRepository struct {
	Repo PGGormRepository.PGGormRepository
}

func CreateNewService(repo PGGormRepository.PGGormRepository) *ServiceRepository {
	return &ServiceRepository{
		Repo: repo,
	}
}

func (s *ServiceRepository) CreateService(model interface{}) error {
	if err := s.Repo.Create(model); err != nil {
		return err
	}
	return nil
}
func (s *ServiceRepository) IndexService(tbName string, page, limit int) ([]map[interface{}]interface{}, error) {
	res, err := s.Repo.Index(tbName, page, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *ServiceRepository) FinByService(tbName, key, value string) ([]map[interface{}]interface{}, error) {
	res, err := s.Repo.FindBy(tbName, key, value)
	if err != nil {
		return nil, err
	}
	var result []map[interface{}]interface{}
	result = append(result, res)
	return result, nil
}

func (s *ServiceRepository) UpdateService(model interface{}, id, tbName string) error {
	if err := s.Repo.Update(model, id, tbName); err != nil {
		return err
	}
	return nil
}

func (s *ServiceRepository) DeleteService(model interface{}, id string) error {
	if err := s.Repo.Delete(model, id); err != nil {
		return err
	}
	return nil
}
