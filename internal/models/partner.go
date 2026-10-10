package models

type Partner struct {
	Id     string
	Name   string
	Offers map[string]Offer
}
