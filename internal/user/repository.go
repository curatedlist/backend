package user

import (

	// Mysql driver
	"backend/internal/database"
	"backend/internal/user/commands"

	"github.com/huandu/go-sqlbuilder"
)

// Repository the User repository
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

func (repo *Repository) getBy(sb *sqlbuilder.SelectBuilder, column string, value interface{}) Aggregate {
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

// GetByID a user
func (repo *Repository) GetByID(id int64) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.id", id)
}

// GetByIss a user
func (repo *Repository) GetByIss(iss string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.iss", iss)
}

// GetByEmail a user
func (repo *Repository) GetByEmail(email string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.email", email)
}

// GetByUsername a user from repository by its id
func (repo *Repository) GetByUsername(username string) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	return repo.getBy(sb, "user.username", username)
}

// GetLists of an user
func (repo *Repository) GetLists(id int64) []ListAggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "list.deleted")
	sb.From("list")
	sb.Where(sb.Equal("list.user_id", id))
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

// GetFavs of an user
func (repo *Repository) GetFavs(id int64) []ListAggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "list.deleted", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url")
	sb.From("list")
	sb.Join("fav", "list.id = fav.list_id")
	sb.Join("user", "user.id = list.user_id")
	sb.Where(sb.Equal("fav.user_id", id))
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

// Create an user
func (repo *Repository) Create(email string) Aggregate {
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
	user := repo.GetByID(id)
	return user
}

// Update an user
func (repo *Repository) Update(id int64, command commands.Update) Aggregate {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("user")
	ub.Set(
		ub.Assign("name", command.Name),
		ub.Assign("username", command.Username),
		ub.Assign("bio", command.Bio),
	)
	ub.Where(ub.Equal("id", id))
	sql, args := ub.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	user := repo.GetByID(id)
	return user
}
