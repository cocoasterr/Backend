package service

import (
	"fmt"
	"regexp"

	"github.com/cocoasterr/warung_app/domain"
	PGRepository "github.com/cocoasterr/warung_app/infra/db/postgre/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Service
}

func NewAuthService(repo PGRepository.Repository) *AuthService {
	return &AuthService{
		Service: Service{
			Repository: repo,
		},
	}
}

func (s *AuthService) CheckUser(data domain.User, tb_name string) interface{} {
	query := fmt.Sprintf("SELECT * FROM %s", tb_name)
	cond := " where email = ? or username = ?"
	res := s.Repository.CustomQuery(&data, query, cond, data.Email, data.Username)
	return res
}

func (s *AuthService) CheckEmail(email string) (bool, error) {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	return regexp.MatchString(emailRegex, email)
}

func (s *AuthService) HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}
