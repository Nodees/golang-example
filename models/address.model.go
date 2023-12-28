package models

type Address struct {
	BaseModel
	Cep          *string `json:"cep"`
	Street       *string `json:"street"`
	Neighborhood *string `json:"neighborhood"`
	City         *string `json:"city"`
	State        *string `json:"state"`
}
