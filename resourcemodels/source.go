package resourcemodels

type Source struct {
	Url         string 	`json:"url"           validate:"required"`
	Channel     string 	`json:"channel"`
	DisablePut  bool 		`json:"disable_put"`
	Debug       bool 		`json:"debug"`
}
