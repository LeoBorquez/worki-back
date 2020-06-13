package handler

type GigRequest struct {
	Tittle   string  `sql:"tittle"`
	Desc     string  `sql:"desc"`
	Image    string  `sql:"image"`
	Hourly   float64 `sql:"hourly,type:real"`
	Location string  `sql:"location"`
	IsActive bool    `sql:"is_activate"`
}