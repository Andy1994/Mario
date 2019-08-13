package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Mario struct {
	UID int
	ID string `json:"id"`
	Name string `json:"name"`
	Nature string `json:"nature"`
}

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./mario.db")
	if err != nil {
		log.Fatalln(err)
	}

	sql := `
    CREATE TABLE IF NOT EXISTS mario (
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
		id VARCHAR(64) NULL,
        name VARCHAR(64) NULL,
        nature VARCHAR(64) NULL
    );
    `
	db.Exec(sql)
}

func (m Mario) Insert(mario Mario) (id int64, err error) {
	stmt, err := db.Prepare("INSERT INTO mario(id, name, nature) values(?,?,?)")
	if err != nil {
		return
	}
	res, err := stmt.Exec(mario.ID, mario.Name, mario.Nature)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	fmt.Println(id)
	return
}

func (m Mario) GetAll() (marios []Mario, err error) {
	rows, err := db.Query("SELECT * FROM mario")
	if err != nil {
		return
	}
	for rows.Next() {
		var mario = new(Mario)
		rows.Scan(&mario.UID, &mario.ID, &mario.Name, &mario.Nature)
		marios = append(marios, *mario)
	}
	defer rows.Close()
	return
}

func (m Mario) Get(id string) (mario Mario, err error) {
	rows, err := db.Query("SELECT * FROM mario WHERE id = ? LIMIT 1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		rows.Scan(&mario.UID, &mario.ID, &mario.Name, &mario.Nature)
		return
	}
	defer rows.Close()
	return
}