package seeders

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"
)

type UserTypeSeeder struct{}

func (u *UserTypeSeeder) Run() []entity.UserType {
	seed := []entity.UserType{
		{
			ID:   1,
			Name: "Admin",
		},

		{
			ID:   2,
			Name: "Teacher",
		},

		{
			ID:   3,
			Name: "Student",
		},
	}

	return seed
}
