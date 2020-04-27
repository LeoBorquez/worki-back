package model

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

/*
	table name : gigs
*/
type Gig struct {
	ID        int       `sql:"id, pk"`
	Tittle    string    `sql:"tittle"`
	Desc      string    `sql:"desc"`
	Image     string    `sql:"image"`
	Hourly    float64   `sql:"hourly,type:real"`
	Location  string    `sql:"location"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_activate"`
}

//CreateGigTable * pointer to, & reference to
func CreateGigTable(db *pg.DB) error {
	err := db.CreateTable(&Gig{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	log.Panic(err)

	log.Printf("Table gigs created. \n")
	return nil

}
