package list

import (
	"backend/internal/config"
	"database/sql"
	"fmt"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Repository the List repository
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

// FindAll find alls models from repository
func (repo *Repository) FindAll() []DatabaseDTO {
	// Prepare statement for reading data
	rows, err := repo.db.Query("SELECT list.id, list.name, list.description FROM list")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	lists := make([]DatabaseDTO, 0)
	for rows.Next() {
		var listDTO DatabaseDTO
		err := rows.Scan(&listDTO.ID, &listDTO.Name, &listDTO.Description)
		if err != nil {
			panic(err.Error())
		}
		lists = append(lists, listDTO)
	}
	return lists
}

// Get a list by ID
func (repo *Repository) Get(id string) DatabaseDTO {
	// Prepare statement for reading data
	rows, err := repo.db.Query("SELECT list.id, list.name, list.description, item.id, item.name, item.url, item.pic_url FROM list LEFT JOIN item ON item.list_id = list.id WHERE list.id = ?", id)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var list DatabaseDTO
	for rows.Next() {
		var item DatabaseItemDTO
		err := rows.Scan(&list.ID, &list.Name, &list.Description, &item.ID, &item.Name, &item.URL, &item.PicURL)
		if err != nil {
			panic(err.Error())
		}
		if item.ID.Valid {
			list.Items = append(list.Items, item)
		}
	}
	return list
}
