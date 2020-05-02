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
	rows, err := repo.db.Query("SELECT user.id, user.name, user.email, user.avatar_url, list.id, list.name, list.description FROM user LEFT JOIN list ON user.id = list.user_id where user.id = ?", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var user DataBaseDTO
	for rows.Next() {
		var list DatabaseListDTO
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL, &list.ID, &list.Name, &list.Description)
		if err != nil {
			panic(err.Error())
		}
		if list.ID.Valid {
			user.Lists = append(user.Lists, list)
		}
	}
	return user
}

// GetByEmail a user from repository by its email
func (repo *Repository) GetByEmail(email string) DataBaseDTO {
	// Prepare statement for reading data
	rows, err := repo.db.Query("SELECT user.id, user.name, user.email, user.avatar_url, list.id, list.name, list.description FROM user LEFT JOIN list ON user.id = list.user_id where user.email = ?", email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var user DataBaseDTO
	for rows.Next() {
		var list DatabaseListDTO
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.AvatarURL, &list.ID, &list.Name, &list.Description)
		if err != nil {
			panic(err.Error())
		}
		if list.ID.Valid {
			user.Lists = append(user.Lists, list)
		}
	}
	return user
}

// CreateUser Create an user
func (repo *Repository) CreateUser(email string) int64 {
	// Prepare statement for reading data
	stmt, err := repo.db.Prepare("INSERT INTO user(email) VALUES (?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	result, err := stmt.Exec(email)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	res, _ := result.LastInsertId()
	return res
}

// UpdateUser Create an user
func (repo *Repository) UpdateUser(id string, name string) int64 {
	// Prepare statement for reading data
	fmt.Printf("name : %s, id: %s", name, id)
	stmt, err := repo.db.Prepare("UPDATE user SET name=? WHERE id =?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	result, err := stmt.Exec(name, id)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	res, _ := result.LastInsertId()
	return res
}
