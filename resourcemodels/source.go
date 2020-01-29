package resourcemodels

type Source struct {
	Url        string `json:"url"         validate:"required"`
	DisablePut bool   `json:"disable_put"`
	Debug      bool   `json:"debug"`
}
