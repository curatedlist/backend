package user

import (

	// Mysql driver
	"backend/internal/database"
	"backend/internal/user/commands"
	"strconv"

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

func (repo *Repository) getBy(sb *sqlbuilder.SelectBuilder, column string, value string) Aggregate {
	sb.Select("user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url", "fav.list_id")
	sb.From("user")
	sb.JoinWithOption("LEFT", "fav", "fav.user_id = user.id")
	sb.Where(sb.Equal(column, value))
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	var user Aggregate
	for rows.Next() {
		var fav FavAggregate
		its := repo.userStruct.Addr(&user)
		its = append(its, repo.favStruct.Addr(&fav)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		if fav.ListID.Valid {
			user.Favs = append(user.Favs, fav)
		}
	}

	sb = sqlbuilder.NewSelectBuilder()
	sb.Select(sb.As("COUNT(*)", "c"))
	sb.From("list")
	sb.Join("user", "user.id = list.user_id")
	sb.Where(sb.Equal(column, value))
	sql, args = sb.Build()
	rows, err = repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&user.Lists)
		if err != nil {
			panic(err.Error())
		}
	}

	return user
}

// GetByID a user from repository by its id
func (repo *Repository) GetByID(id string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.id", id)
}

// GetByEmail a user from repository by its email
func (repo *Repository) GetByEmail(email string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.email", email)
}

// GetByUsername a user from repository by its id
func (repo *Repository) GetByUsername(username string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.username", username)
}

// GetLists a user from repository by its id
func (repo *Repository) GetLists(userID uint) []ListAggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description")
	sb.From("list")
	sb.Where(sb.Equal("list.user_id", userID))
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	lists := make([]ListAggregate, 0)
	for rows.Next() {
		var list ListAggregate
		err := rows.Scan(repo.listStruct.Addr(&list)...)
		if err != nil {
			panic(err.Error())
		}
		lists = append(lists, list)
	}
	return lists
}

// GetFavs a user from repository by its id
func (repo *Repository) GetFavs(userID uint) []ListAggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url")
	sb.From("list")
	sb.Join("fav", "list.id = fav.list_id")
	sb.Join("user", "user.id = list.user_id")
	sb.Where(sb.Equal("fav.user_id", userID))
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}
	lists := make([]ListAggregate, 0)
	for rows.Next() {
		var list ListAggregate
		var user Aggregate
		its := repo.listStruct.Addr(&list)
		its = append(its, repo.userStruct.Addr(&user)...)
		err := rows.Scan(its...)
		if err != nil {
			panic(err.Error())
		}
		list.Owner = user
		lists = append(lists, list)
	}
	return lists
}

// CreateUser Create an user
func (repo *Repository) CreateUser(email string) Aggregate {
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
	id, _ := result.LastInsertId()
	user := repo.GetByID(strconv.FormatInt(id, 10))
	return user
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
