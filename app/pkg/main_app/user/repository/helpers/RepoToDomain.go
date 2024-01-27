package helper

import (
	"app/pkg/main_app/user/domain"
	db "app/pkg/main_app/user/repository"
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
