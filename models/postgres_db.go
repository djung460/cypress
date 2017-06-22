package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type Connection struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type SqlDB struct {
	db *sql.DB
}

// Init connects to an external sqlite DB and returns a DB implementation
// that fronts that connection. call Close() on the returned value when done.
func Init() (SqlDB, error) {
	raw, err := ioutil.ReadFile("config/db_config.json")

	if err != nil {
		log.Fatal(err)
	}

	var conn Connection

	err = json.Unmarshal(raw, &conn)

	if err != nil {
		log.Fatal(err)
	}

	dbinfo := fmt.Sprintf("dbname=%s host=%s port=%d user=%s password=%s sslmode=disable", conn.Name, conn.Hostname, conn.Port, conn.User, conn.Password)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	return SqlDB{db: db}, err
}

// Close releases the underlying connections. always call this when
// completely done with operations, not before.
func (m *SqlDB) Close() {
	m.db.Close()
}

func (m SqlDB) CreateNugget(ngt Nugget) (bool, error) {
	query := `
  INSERT nugget
	(author_id, category_id, title, body, created_at, edited_at)
	values (?,?,?,?,?,?)`
	stmt, err := m.db.Prepare(query)
	if err != nil {
		log.Println("STATEMENT ERROR")
		log.Print(err)
		return false, err
	}
	res, err := stmt.Exec(
		ngt.AuthorID,
		ngt.CategoryID,
		ngt.Title,
		ngt.Body,
		ngt.CreatedAt,
		ngt.EditedAt,
	)
	_ = res

	if err != nil {
		log.Println("EXEC ERROR")
		log.Print(err)
		return false, err
	}
	return true, nil
}

func (m SqlDB) GetNuggetsByUser(id int) ([]Nugget, error) {
	query := `
  SELECT
  id, author_id, category_id, title, body, created_at, edited_at
  FROM nugget
  WHERE author_id=$1`

	log.Printf("Fetching nuggets for user id %d", id)

	rows, err := m.db.Query(query, id)

	var nugget Nugget
	var nuggets []Nugget

	if err != nil {
		log.Println("STATEMENT ERROR")
		return []Nugget{}, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&nugget.ID,
			&nugget.AuthorID,
			&nugget.CategoryID,
			&nugget.Title,
			&nugget.Body,
			&nugget.CreatedAt,
			&nugget.EditedAt,
		)
		if err != nil {
			return []Nugget{}, err
		}
		nuggets = append(nuggets, nugget)
	}
	return nuggets, err
}

func (m SqlDB) CreateCategory(cat Category) (bool, error) {
	query := `
  INSERT category
	(name)
	values (?)`
	stmt, err := m.db.Prepare(query)
	if err != nil {
		log.Println("STATEMENT ERROR")
		log.Print(err)
		return false, err
	}
	res, err := stmt.Exec(
		cat.Name,
	)
	_ = res

	if err != nil {
		log.Println("EXEC ERROR")
		log.Print(err)
		return false, err
	}
	return true, nil
}

func (m SqlDB) GetCategories() ([]Category, error) {
	query := `
  SELECT
  id, name
  FROM category`

	rows, err := m.db.Query(query)

	var category Category
	var categories []Category

	if err != nil {
		return []Category{}, err
	}

	for rows.Next() {
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)
		if err != nil {
			return []Category{}, err
		}
		categories = append(categories, category)
	}
	rows.Close()
	return categories, err
}
