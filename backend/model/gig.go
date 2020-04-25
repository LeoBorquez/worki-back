package model

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/v9/orm"
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

// * pointer to, & reference to
func CreateGigTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&Gig{}, opts)
	if err != nil {
		log.Panic("Error creating table gigs, Reason: %v\n", err)
		return err
	}
	log.Printf("Table gigs created. \n")
	return nil

}
