package user

import (

	// Mysql driver
	"backend/internal/database"
	"backend/internal/user/commands"

	"github.com/huandu/go-sqlbuilder"
)

// Repository the List repository
type Repository struct {
	db         database.DB
	userStruct *sqlbuilder.Struct
	listStruct *sqlbuilder.Struct
	favStruct  *sqlbuilder.Struct
}

// NewRepository returns a Repository
func NewRepository(db database.DB) Repository {
	return Repository{
		db:         db,
		userStruct: sqlbuilder.NewStruct(new(Aggregate)),
		listStruct: sqlbuilder.NewStruct(new(ListAggregate)),
		favStruct:  sqlbuilder.NewStruct(new(FavAggregate)),
	}
}

func (repo *Repository) getBy(sb *sqlbuilder.SelectBuilder, equal string) Aggregate {
	sb.Select("user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url", "list.id", "list.name", "list.description", "fav.list_id")
	sb.From("user")
	sb.JoinWithOption("LEFT", "list", "list.user_id = user.id")
	sb.JoinWithOption("LEFT", "fav", "fav.user_id = user.id")
	sb.Where(equal)
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	var user Aggregate
	for rows.Next() {
		var list ListAggregate
		var fav FavAggregate
		its := repo.userStruct.Addr(&user)
		its = append(its, repo.listStruct.Addr(&list)...)
		its = append(its, repo.favStruct.Addr(&fav)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		if list.ID.Valid {
			user.Lists = append(user.Lists, list)
		}
		if fav.ListID.Valid {
			user.Favs = append(user.Favs, fav)
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

// GetByUsername a user from repository by its id
func (repo *Repository) GetByUsername(username string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, sb.Equal("user.username", username))
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
	result, err := stmt.Exec(args...)

	if err != nil {
		panic(err.Error())
	}
	res, _ := result.LastInsertId()
	return res
}

// UpdateUser Create an user
func (repo *Repository) UpdateUser(id string, updateCommand commands.Update) Aggregate {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("user")
	ub.Set(
		ub.Assign("name", updateCommand.Name),
		ub.Assign("username", updateCommand.Username),
		ub.Assign("bio", updateCommand.Bio),
	)
	ub.Where(ub.Equal("id", id))
	sql, args := ub.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	userAggregate := repo.GetByID(id)
	return userAggregate
}
