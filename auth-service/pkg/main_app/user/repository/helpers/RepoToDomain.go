package helper

import (
	db "auth-service/pkg/main_app/user/repository"
	"auth-service/pkg/main_app/user/domain"
)

func ToDomainUser(u *db.User) *domain.User {
	return &domain.User{

		Firstname:  u.Firstname,
		Middlename: u.Middlename,
		Lastname:   u.Lastname,
		Email:      u.Email,
		Password:   u.Password,
	}

}
