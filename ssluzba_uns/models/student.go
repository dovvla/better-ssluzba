package models

type Student struct {
	Ime     string `json:"ime"`
	Prezime string `json:"prezime"`
	JMBG    string `json:"jmbg" gorm:"primary_key"`
}
