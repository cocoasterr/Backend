package service

import PGGormRepository "github.com/cocoasterr/gowarungapp/infra/db/postgres/repository/gorm"

type PersonService struct {
	ServiceRepository
}

func CreateNewPersonService(repo PGGormRepository.PGGormRepository) *PersonService {
	return &PersonService{
		ServiceRepository: ServiceRepository{
			Repo: repo,
		},
	}
}
