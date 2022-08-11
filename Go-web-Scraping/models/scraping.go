package models

type UrlString struct {
	UrlStr string `json:"urlStr"`
}

type Scarping struct {
	Version   string     `json:"version"`
	Title     string     `json:"title"`
	Headings  [][]string `json:"headings"`
	Links     [][]string `json:"link"`
	Loginform bool       `json:"loginform"`
}

var zeroScarping = &Scarping{}

func (a *Scarping) Reset() {
	*a = *zeroScarping
}

var zeroUrlString = &UrlString{}

func (a *UrlString) Reset() {
	*a = *zeroUrlString
}
