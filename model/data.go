package model

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
	"github.com/jinzhu/gorm"
)

func FakeData(db *gorm.DB) {
	FakeGig(db)
	FakeProposal(db)
	FakeComment(db)
}

// FakeGig create fake data for testing
func FakeGig(db *gorm.DB) {

	var gigs []Gig

	for i := 0; i < 100; i++ {
		g := Gig{
			UserID:      uint(gofakeit.Number(1, 3)),
			CategoryID:  uint(gofakeit.Number(1, 10)),
			StatusID:    uint(gofakeit.Number(1, 4)),
			Tittle:      gofakeit.Sentence(3),
			Description: gofakeit.Paragraph(1, 2, 10, " "),
			Image:       gofakeit.ImageURL(500, 500),
			Rate:        gofakeit.Price(10, 100),
			Location:    gofakeit.City(),
		}
		gigs = append(gigs, g)
	}

	for _, gig := range gigs {
		if err := db.Create(&gig).Error; err != nil {
			fmt.Print(err)
		}
	}

}

// FakeProposal data
func FakeProposal(db *gorm.DB) {

	var proposals []Proposal

	for i := 0; i < 100; i++ {
		p := Proposal{
			UserID:      uint(gofakeit.Number(1, 3)),
			GigID:       uint(gofakeit.Number(1, 3)),
			StatusID:    uint(gofakeit.Number(1, 4)),
			Description: gofakeit.Paragraph(1, 2, 10, " "),
			Rate:        gofakeit.Price(10, 1000),
			TimeFrame:   int(gofakeit.Number(1, 500)),
		}
		proposals = append(proposals, p)
	}

	for _, prop := range proposals {
		if err := db.Create(&prop).Error; err != nil {
			fmt.Print(err)
		}
	}
}

func FakeComment(db *gorm.DB) {
	var comments []Comment

	for i := 0; i < 100; i++ {
		c := Comment{
			GigID:       uint(gofakeit.Number(1, 100)),
			UserID:      uint(gofakeit.Number(1, 5)),
			Description: gofakeit.Paragraph(1, 1, 20, ""),
			Like:        uint(gofakeit.Number(1, 100)),
		}
		comments = append(comments, c)
	}

	for _, comm := range comments {
		if err := db.Create(&comm).Error; err != nil {
			fmt.Print(err)
		}
	}
}
