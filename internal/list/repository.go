package list

import (

	// Mysql driver
	"backend/internal/database"
	"backend/internal/list/commands"
	"strconv"

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
	sb.Select("list.id", "list.name", "list.description", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url")
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
	sb.Select("list.id", "list.name", "list.description", "user.id", "user.name", "user.email", "user.username", "user.bio", "user.avatar_url", "item.id", "item.name", "item.description", "item.url", "item.pic_url", "item.deleted", "item.list_id")
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

// GetItem returns an item by ID
func (repo *Repository) GetItem(id string) ItemAggregate {
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

// CreateList creates a list
func (repo *Repository) CreateList(userID string, createListCommand commands.CreateList) int64 {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("list")
	ib.Cols("name", "description", "user_id")
	ib.Values(createListCommand.Name, createListCommand.Description, userID)
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

// CreateItem creates a item for a list
func (repo *Repository) CreateItem(listID string, createItemCommand commands.CreateItem) ItemAggregate {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("item")
	ib.Cols("name", "description", "url", "pic_url", "list_id")
	ib.Values(createItemCommand.Name, createItemCommand.Description, createItemCommand.URL, createItemCommand.PicURL, listID)
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
	itemAggregate := repo.GetItem(strconv.FormatInt(itemID, 10))
	return itemAggregate
}

// DeleteItem Create an user
func (repo *Repository) DeleteItem(id string) ItemAggregate {
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
	itemAggregate := repo.GetItem(id)
	return itemAggregate
}

// FavList favs a list
func (repo *Repository) FavList(listID string, userID string) Aggregate {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("fav")
	ib.Cols("list_id", "user_id")
	ib.Values(listID, userID)
	sql, args := ib.Build()
	stmt, err := repo.db.DB.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(args...)

	if err != nil {
		panic(err.Error())
	}
	listAggregate := repo.Get(listID)
	return listAggregate
}
