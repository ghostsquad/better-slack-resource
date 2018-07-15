package models

type Source struct {
	Url         string `json:"url"`
	Channel     string `json:"channel"`
	DisablePut  bool `json:"disable_put"`
	Debug       bool `json:"debug"`
}
