package seeders

import (
	"github.com/mustafaakilll/go-site-exam/db/entity"
)

type UserSeeder struct{}

func (u *UserSeeder) Run() []entity.User {
	seed := []entity.User{
		{
			Username:   "admin",
			UserTypeID: 1,
			Email:      "admin@admin.com",
			FirstName:  "admin",
			LastName:   "admin",
			Password:   "$2a$14$Dkyl.rEzq9yyAA8f3WBscuT68KjNqnGItw2ML06JhVt2KL5489Kbq",
		},
	}

	return seed
}
