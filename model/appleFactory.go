package model

import (
    "context"
    "database/sql"
)


func (u *User) Create() (*model.Apple, error) {
    // dbUsername := os.Getenv("postgres_username")
	// dbPassword := os.Getenv("postgres_password")
	// dbHost := os.Getenv("postgres_host")
	// dbPort := os.Getenv("postgres_port")
	// dbName := os.Getenv("postgres_db_name")

    // dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort)
    // db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})


    // apple = model.Apple("")
    // db.Create()
}
func (u *User) FetchByID(id uint) (*model.Apple, error) {
    dbUsername := os.Getenv("postgres_username")
	dbPassword := os.Getenv("postgres_password")
	dbHost := os.Getenv("postgres_host")
	dbPort := os.Getenv("postgres_port")
	dbName := os.Getenv("postgres_db_name")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUsername, dbPassword, dbName, dbPort)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    var apple Apple
    result = db.First(&apple, id)

    return result.Error
}
func (u *User) FetchAll(...) ([]*model.Apple, error) {
    ...
}
func (u *User) Update(...) (*model.Apple, error) {
    ...
}
func (u *User) Delete(...) error {
   ...
}
