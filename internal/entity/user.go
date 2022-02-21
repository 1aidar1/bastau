package entity

import (
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password []byte `json:"-"`
	//IsMale      bool         `json:"is_male"`
	//CountryId   int          `json:"country_id"`
	//CityId      int          `json:"city_id"`
	//Occupations []Occupation `json:"occupations"`
	//DateOfBirth time.Time    `json:"date_of_birth"`
	//CreatedAt   time.Time    `json:"created_at"`
	//UpdatedAt   time.Time    `json:"updated_at"`
	//LastVisitAt time.Time    `json:"last_visit_at"`
}

var EmptyUser User

var ErrUserAlreadyExists error = errors.New("user with this credentials already exists")
