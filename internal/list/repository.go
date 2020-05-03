package list

import (

	// Mysql driver
	"backend/internal/database"

	"github.com/huandu/go-sqlbuilder"
)

// Repository the List repository
type Repository struct {
	db          database.DB
	listStruct  *sqlbuilder.Struct
	ownerStruct *sqlbuilder.Struct
	itemStruct  *sqlbuilder.Struct
}

// NewRepository returns a Repository
func NewRepository(db database.DB) Repository {
	return Repository{
		db:          db,
		listStruct:  sqlbuilder.NewStruct(new(Aggregate)),
		ownerStruct: sqlbuilder.NewStruct(new(OwnerAggregate)),
		itemStruct:  sqlbuilder.NewStruct(new(ItemAggregate)),
	}
}

// FindAll find alls models from repository
func (repo *Repository) FindAll() []Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "user.id", "user.name", "user.email", "user.avatar_url")
	sb.From("list")
	sb.Join("user", "user.id = list.user_id")
	sql, _ := sb.Build()
	rows, err := repo.db.DB.Query(sql)
	if err != nil {
		panic(err.Error())
	}
	lists := make([]Aggregate, 0)
	for rows.Next() {
		var list Aggregate
		its := repo.listStruct.Addr(&list)
		its = append(its, repo.ownerStruct.Addr(&list.Owner)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		lists = append(lists, list)
	}
	return lists
}

// Get a list by ID
func (repo *Repository) Get(id string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "user.id", "user.name", "user.email", "user.avatar_url", "item.id", "item.name", "item.url", "item.pic_url")
	sb.From("list")
	sb.Join("user", "user.id = list.user_id")
	sb.JoinWithOption("LEFT", "item", "item.list_id = list.id")
	sb.Where(sb.Equal("list.id", id))
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}

	var list Aggregate
	for rows.Next() {
		var item ItemAggregate
		its := repo.listStruct.Addr(&list)
		its = append(its, repo.ownerStruct.Addr(&list.Owner)...)
		its = append(its, repo.itemStruct.Addr(&item)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		if item.ID.Valid {
			list.Items = append(list.Items, item)
		}

	}
	return list
}
