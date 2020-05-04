package user

import (

	// Mysql driver
	"backend/internal/database"

	"github.com/huandu/go-sqlbuilder"
)

// Repository the List repository
type Repository struct {
	db         database.DB
	userStruct *sqlbuilder.Struct
	listStruct *sqlbuilder.Struct
}

// NewRepository returns a Repository
func NewRepository(db database.DB) Repository {
	return Repository{
		db:         db,
		userStruct: sqlbuilder.NewStruct(new(Aggregate)),
		listStruct: sqlbuilder.NewStruct(new(ListAggregate)),
	}
}

func (repo *Repository) getBy(sb *sqlbuilder.SelectBuilder, equal string) Aggregate {
	sb.Select("user.id", "user.name", "user.email", "user.avatar_url", "list.id", "list.name", "list.description")
	sb.From("user")
	sb.JoinWithOption("LEFT", "list", "list.user_id = user.id")
	sb.Where(equal)
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	var user Aggregate
	for rows.Next() {
		var list ListAggregate
		its := repo.userStruct.Addr(&user)
		its = append(its, repo.listStruct.Addr(&list)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		if list.ID.Valid {
			user.Lists = append(user.Lists, list)
		}
	}
	return user
}

// GetByID a user from repository by its id
func (repo *Repository) GetByID(id string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, sb.Equal("user.id", id))
}

// GetByEmail a user from repository by its email
func (repo *Repository) GetByEmail(email string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, sb.Equal("user.email", email))
}

// CreateUser Create an user
func (repo *Repository) CreateUser(email string) int64 {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("user")
	ib.Cols("email")
	ib.Values(email)
	sql, args := ib.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(args)

	if err != nil {
		panic(err.Error())
	}
	res, _ := result.LastInsertId()
	return res
}

// UpdateUser Create an user
func (repo *Repository) UpdateUser(id string, name string) int64 {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("user")
	ub.Set(ub.Assign("name", name))
	ub.Where(ub.Equal("id", id))
	sql, args := ub.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	result, err := stmt.Exec(args)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	res, _ := result.LastInsertId()
	return res
}
