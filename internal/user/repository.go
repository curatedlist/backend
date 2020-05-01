package user

import (
	"backend/internal/config"
	"database/sql"
	"fmt"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Repository the User repository
type Repository struct {
	db *sql.DB
}

// NewRepository returns a Repository
func NewRepository(conf config.DBConfig) Repository {

	connectionString := fmt.Sprintf("%s:%s@%s/curatedlist", conf.Username, conf.Password, conf.URL)
	d, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return Repository{db: d}
}

// Get a user from repository by its id
func (repo *Repository) Get(id string) DataBaseDTO {
	// Prepare statement for reading data
	rows, err := repo.db.Query("SELECT user.id, user.name, user.email, user.avatar_url FROM user WHERE user.id = ?", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var user DataBaseDTO
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL)
		if err != nil {
			panic(err.Error())
		}
	}
	return user
}
