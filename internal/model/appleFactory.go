package model

import (
	"helloGo/database/connection"
)

func FetchAppleByID(id uint) (*Apple, error) {
	db, err := connection.GetDbConnection()
	if err != nil {
		return nil, err
	}
	var apple Apple
	result := db.First(&apple, id)
	return &apple, result.Error
}

func (u *Apple) CreateApple() (*Apple, error) {
	panic("unimplemented")
}
func (u *Apple) FetchAllApples() ([]*Apple, error) {
	panic("unimplemented")
}
func (u *Apple) UpdateApple() (*Apple, error) {
	panic("unimplemented")
}
func (u *Apple) DeleteApple() error {
	panic("unimplemented")
}
