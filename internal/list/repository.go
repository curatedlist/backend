package list

import (

	// Mysql driver
	"backend/internal/database"
	"backend/internal/list/commands"

	"github.com/huandu/go-sqlbuilder"
)

// Repository is the List repository
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

// FindAll list
func (repo *Repository) FindAll(filter string) []Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "list.deleted", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url")
	sb.From("list")
	sb.Join("user", "user.id = list.user_id")
	if filter == "trending" {
		sb.Join("fav", "fav.list_id = list.id")
		sb.OrderBy("COUNT(*)").Desc()
		sb.GroupBy("list.id")
	}
	sb.Limit(5)
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

// Get a list
func (repo *Repository) Get(id int64) Aggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("list.id", "list.name", "list.description", "list.deleted", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url", "item.id", "item.name", "item.description", "item.url", "item.pic_url", "item.deleted", "item.list_id")
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

	sb = sqlbuilder.NewSelectBuilder()
	sb.Select(sb.As("COUNT(*)", "c"))
	sb.From("list")
	sb.Join("fav", "fav.list_id = list.id")
	sb.Where(sb.Equal("list.id", id))
	sql, args = sb.Build()

	rows, err = repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&list.Favs)
		if err != nil {
			panic(err.Error())
		}
	}
	return list
}

// GetItem of a list
func (repo *Repository) GetItem(id int64) ItemAggregate {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("item.id", "item.name", "item.description", "item.url", "item.pic_url", "item.deleted", "item.list_id")
	sb.From("item")
	sb.Where(sb.Equal("item.id", id))
	sql, args := sb.Build()
	rows, err := repo.db.DB.Query(sql, args...)
	if err != nil {
		panic(err.Error())
	}

	var item ItemAggregate
	if rows.Next() {
		err := rows.Scan(repo.itemStruct.Addr(&item)...)
		if err != nil {
			panic(err.Error())
		}
	}
	return item
}

// Create a list
func (repo *Repository) Create(userID int64, command commands.CreateList) Aggregate {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("list")
	ib.Cols("name", "description", "user_id")
	ib.Values(command.Name, command.Description, userID)
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
	list := repo.Get(id)
	return list
}

// Delete a list
func (repo *Repository) Delete(id int64) Aggregate {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("list")
	ub.Set(ub.Assign("deleted", true))
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

	list := repo.Get(id)
	return list
}

// Fav a list
func (repo *Repository) Fav(id int64, userID int64) Aggregate {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("fav")
	ib.Cols("list_id", "user_id")
	ib.Values(id, userID)
	sql, args := ib.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	list := repo.Get(id)
	return list
}

// Unfav a list
func (repo *Repository) Unfav(id int64, userID int64) Aggregate {
	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("fav")
	db.Where(db.And(db.Equal("user_id", userID), db.Equal("list_id", id)))
	sql, args := db.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	list := repo.Get(id)
	return list
}

// CreateItem for a list
func (repo *Repository) CreateItem(id int64, command commands.CreateItem) ItemAggregate {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("item")
	ib.Cols("name", "description", "url", "pic_url", "list_id")
	ib.Values(command.Name, command.Description, command.URL, command.PicURL, id)
	sql, args := ib.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		panic(err.Error())
	}

	itemID, _ := result.LastInsertId()
	item := repo.GetItem(itemID)
	return item
}

// DeleteItem from a list
func (repo *Repository) DeleteItem(id int64) ItemAggregate {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("item")
	ub.Set(ub.Assign("deleted", true))
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

	item := repo.GetItem(id)
	return item
}
