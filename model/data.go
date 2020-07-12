package model

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/jinzhu/gorm"
)

// FakeGig create fake data for testing
func FakeGig(db *gorm.DB) {

	var gigs []Gig

	for i := 0; i < 100; i++ {
		g := Gig{
			UserID:   uint(gofakeit.Number(1, 3)),
			Tittle:   gofakeit.Sentence(3),
			Desc:     gofakeit.Paragraph(1, 3, 55, " "),
			Image:    gofakeit.ImageURL(500, 500),
			Rate:     gofakeit.Price(10, 100),
			Location: gofakeit.City(),
			IsActive: true,
		}
		gigs = append(gigs, g)
	}

	for _, gig := range gigs {
		if err := db.Create(&gig).Error; err != nil {
			fmt.Print(err)
		}
	}

}
