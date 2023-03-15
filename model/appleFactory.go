package dal

import (
    "context"
    "database/sql"
    ...
)
type User struct {
   DB *sql.DB
   ...
}
func (u *User) Create(...) (*model.UserEntity, error) {
    ...
}
func (u *User) FetchByID(...) (*model.UserEntity, error) {
    ...
}
func (u *User) FetchAll(...) ([]*model.UserEntity, error) {
    ...
}
func (u *User) Update(...) (*model.UserEntity, error) {
    ...
}
func (u *User) Delete(...) error {
   ...
}
